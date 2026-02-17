/*
 * Copyright (C) 2019 The ORT Project Copyright Holders <https://github.com/oss-review-toolkit/ort/blob/main/NOTICE>
 * Copyright (C) 2026 Timo Reymann
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 * License-Filename: LICENSE
 */


/**
 * Import the license classifications from license-classifications.yml.
 */
val permissiveLicenses = licenseClassifications.licensesByCategory["permissive"].orEmpty()
val copyleftLicenses = licenseClassifications.licensesByCategory["copyleft"].orEmpty()
val copyleftLimitedLicenses = licenseClassifications.licensesByCategory["copyleft-limited"].orEmpty()
val publicDomainLicenses = licenseClassifications.licensesByCategory["public-domain"].orEmpty()

val LicensePresets = mapOf(
    "Apache-2.0" to permissiveLicenses + copyleftLimitedLicenses + publicDomainLicenses + setOf("Unlicense"),
    "MIT" to permissiveLicenses + copyleftLicenses + copyleftLimitedLicenses + publicDomainLicenses + setOf("Unlicense"),
    "GPL-3.0" to setOf("GPL-3.0", "LGPL-3.0", "GPL-2.0", "AGPL-3.0") + permissiveLicenses + publicDomainLicenses + setOf("Unlicense"),
    "Unlicense" to permissiveLicenses + copyleftLimitedLicenses + publicDomainLicenses + setOf("Unlicense")
)

val defaultAllowedLicenses = permissiveLicenses + copyleftLimitedLicenses + publicDomainLicenses + setOf("Unlicense")

val detectedRootLicense = ortResult.getScanResults()
    .flatMap { (_, results) -> results }
    .flatMap { it.summary.licenseFindings }
    .filter { it.location.path == "LICENSE" }
    .map { it.license.toString() }
    .toSet()
    .firstOrNull() ?: "Apache-2.0"

// The complete set of licenses covered by policy rules.
val handledLicenses = listOf(
    permissiveLicenses,
    publicDomainLicenses,
    copyleftLicenses,
    copyleftLimitedLicenses
).flatten().let {
    it.getDuplicates().let { duplicates ->
        require(duplicates.isEmpty()) {
            "The classifications for the following licenses overlap: $duplicates"
        }
    }

    it.toSet()
}


/**
 * Return the Markdown-formatted text to aid users with resolving violations.
 */
fun PackageRule.howToFixDefault() = """
        A text written in MarkDown to help users resolve policy violations
        which may link to additional resources.
    """.trimIndent()

/**
 * Set of matchers to help keep policy rules easy to understand
 */

fun PackageRule.LicenseRule.isHandled() =
    object : RuleMatcher {
        override val description = "isHandled($license)"

        override fun matches() =
            license in handledLicenses && ("-exception" !in license.toString() || " WITH " in license.toString())
    }

/**
 * Policy rules
 */

fun RuleSet.unhandledLicenseRule() = packageRule("UNHANDLED_LICENSE") {
    // Do not trigger this rule on packages that have been excluded in the .ort.yml.
    require {
        -isExcluded()
    }

    // Define a rule that is executed for each license of the package.
    licenseRule("UNHANDLED_LICENSE", LicenseView.CONCLUDED_OR_DECLARED_AND_DETECTED) {
        require {
            -isExcluded()
            -isHandled()
        }

        // Throw an error message including guidance how to fix the issue.
        error(
            "The license $license is currently not covered by policy rules. " +
                    "The license was ${licenseSource.name.lowercase()} in package " +
                    "${pkg.metadata.id.toCoordinates()}.",
            howToFixDefault()
        )
    }
}

fun RuleSet.unmappedDeclaredLicenseRule() = packageRule("UNMAPPED_DECLARED_LICENSE") {
    require {
        -isExcluded()
    }

    resolvedLicenseInfo.licenseInfo.declaredLicenseInfo.processed.unmapped.forEach { unmappedLicense ->
        warning(
            "The declared license '$unmappedLicense' could not be mapped to a valid license or parsed as an SPDX " +
                    "expression. The license was found in package ${pkg.metadata.id.toCoordinates()}.",
            howToFixDefault()
        )
    }
}

fun RuleSet.dependencyInProjectSourceRule() = projectSourceRule("DEPENDENCY_IN_PROJECT_SOURCE_RULE") {
    val denyDirPatterns = listOf(
        "**/node_modules" to setOf("NPM", "Yarn", "PNPM"),
        "**/vendor" to setOf("GoMod")
    )

    denyDirPatterns.forEach { (pattern, packageManagers) ->
        val offendingDirs = projectSourceFindDirectories(pattern)

        if (offendingDirs.isNotEmpty()) {
            issue(
                Severity.ERROR,
                "The directories ${offendingDirs.joinToString()} belong to the package manager(s) " +
                        "${packageManagers.joinToString()} and must not be committed.",
                "Please delete the directories: ${offendingDirs.joinToString()}."
            )
        }
    }
}

fun RuleSet.vulnerabilityInPackageRule() = packageRule("VULNERABILITY_IN_PACKAGE") {
    require {
        -isExcluded()
        +hasVulnerability()
    }

    issue(
        Severity.WARNING,
        "The package ${pkg.metadata.id.toCoordinates()} has a vulnerability",
        howToFixDefault()
    )
}

fun RuleSet.highSeverityVulnerabilityInPackageRule() = packageRule("HIGH_SEVERITY_VULNERABILITY_IN_PACKAGE") {
    val scoreThreshold = 5.0f
    val scoringSystem = "CVSS:3.1"

    require {
        -isExcluded()
        +hasVulnerability(scoreThreshold, scoringSystem)
    }

    issue(
        Severity.ERROR,
        "The package ${pkg.metadata.id.toCoordinates()} has a vulnerability with $scoringSystem severity > " +
            "$scoreThreshold.",
        howToFixDefault()
    )
}

fun RuleSet.deprecatedScopeExcludeReasonInOrtYmlRule() = ortResultRule("DEPRECATED_SCOPE_EXCLUDE_REASON_IN_ORT_YML") {
    val reasons = ortResult.repository.config.excludes.scopes.mapTo(mutableSetOf()) { it.reason }

    @Suppress("DEPRECATION")
    val deprecatedReasons = setOf(ScopeExcludeReason.TEST_TOOL_OF)

    reasons.intersect(deprecatedReasons).forEach { offendingReason ->
        warning(
            "The repository configuration is using the deprecated scope exclude reason '$offendingReason'.",
            "Please use only non-deprecated scope exclude reasons, see " +
                    "https://github.com/oss-review-toolkit/ort/blob/main/model/src/main/" +
                    "kotlin/config/ScopeExcludeReason.kt."
        )
    }
}

fun RuleSet.missingCiConfigurationRule() = projectSourceRule("MISSING_CI_CONFIGURATION") {
    require {
        -AnyOf(
            projectSourceHasFile(
                ".appveyor.yml",
                ".bitbucket-pipelines.yml",
                ".gitlab-ci.yml",
                ".travis.yml"
            ),
            projectSourceHasDirectory(
                ".circleci",
                ".github/workflows"
            )
        )
    }

    error(
        message = "This project does not have any known CI configuration files.",
        howToFix = "Please setup a CI. If you already have setup a CI and the error persists, please contact support."
    )
}

fun RuleSet.missingContributingFileRule() = projectSourceRule("MISSING_CONTRIBUTING_FILE") {
    require {
        -projectSourceHasFile("CONTRIBUTING.md")
    }

    error("The project's code repository does not contain the file 'CONTRIBUTING.md'.")
}

fun RuleSet.missingReadmeFileRule() = projectSourceRule("MISSING_README_FILE") {
    require {
        -projectSourceHasFile("README.md")
    }

    error("The project's code repository does not contain the file 'README.md'.")
}

fun RuleSet.missingReadmeFileLicenseSectionRule() = projectSourceRule("MISSING_README_FILE_LICENSE_SECTION") {
    require {
        +projectSourceHasFile("README.md")
        -projectSourceHasFileWithContent(".*^#{1,2} License$.*", "README.md")
    }

    error(
        message = "The file 'README.md' is missing a \"License\" section.",
        howToFix = "Please add a \"License\" section to the file 'README.md'."
    )
}

fun RuleSet.wrongLicenseInLicenseFileRule() = projectSourceRule("WRONG_LICENSE_IN_LICENSE_FILE_RULE") {
    require {
        +projectSourceHasFile("LICENSE")
    }

    val allowedRootLicenses = setOf("Apache-2.0", "MIT", "GPL-3.0", "Unlicense")
    val detectedRootLicenses = projectSourceGetDetectedLicensesByFilePath("LICENSE").values.flatten().toSet()
    val wrongLicenses = detectedRootLicenses - allowedRootLicenses

    if (wrongLicenses.isNotEmpty()) {
        error(
            message = "The file 'LICENSE' contains the following disallowed licenses ${wrongLicenses.joinToString()}.",
            howToFix = "Please use only the following allowed licenses: ${allowedRootLicenses.joinToString()}."
        )
    } else if (detectedRootLicenses.isEmpty()) {
        error(
            message = "The file 'LICENSE' does not contain any license which is not allowed.",
            howToFix = "Please use one of the following allowed licenses: ${allowedRootLicenses.joinToString()}."
        )
    }
}

fun RuleSet.licenseCompatibilityRule() = packageRule("LICENSE_COMPATIBILITY") {
    require { -isExcluded() }

    val allowedLicenses = LicensePresets[detectedRootLicense] ?: defaultAllowedLicenses

    licenseRule("LICENSE_COMPATIBILITY", LicenseView.CONCLUDED_OR_DECLARED_AND_DETECTED) {
        require {
            -isExcluded()
            license !in allowedLicenses
        }

        error(
            "The license '$license' is incompatible with the project's root license '$detectedRootLicense'. " +
            "Allowed licenses: ${allowedLicenses.joinToString()}.",
            howToFixDefault()
        )
    }
}


/**
 * The set of policy rules.
 */
val ruleSet = ruleSet(ortResult, licenseInfoResolver, resolutionProvider) {
    unhandledLicenseRule()
    unmappedDeclaredLicenseRule()
    vulnerabilityInPackageRule()
    highSeverityVulnerabilityInPackageRule()
    deprecatedScopeExcludeReasonInOrtYmlRule()
    dependencyInProjectSourceRule()
    missingCiConfigurationRule()
    missingContributingFileRule()
    missingReadmeFileRule()
    wrongLicenseInLicenseFileRule()
    licenseCompatibilityRule()
}

// Populate the list of policy rule violations to return.
ruleViolations += ruleSet.violations