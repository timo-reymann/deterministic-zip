File patterns
===

While zip doesn't make a difference between a single and multiple stars, this is a major difference for this tool.

## Examples

| Pattern   | Match with zip                        | Match with deterministic-zip
| :-------- | :------------------------------------ | :--------------------------------
| `.git*`   | Matches `.git` and all children       | Doesnt ignore `.git`
| `.git/*`  | Matches `.git` and all children       | Matches `.git` and all children
| `.git/**`  | Matches `.git` and all children      | Matches `.git` and all children

As you can see the deterministic-zip DOES NOT use the star as an wildcard rather than a glob pattern.

> Advanced matching with sets e. g. `*.[!o]` works out of the box and should behave simliar to zip.

**If you encounter everything that also doesn't work or is not similar, please create an issue.**
