deterministic-zip
===

> Work In Progress

Simple (almost drop-in) replacement for zip that produces deterministic files.

## Installation

> TBD

## Usage

> TBD

## FAQ

### Why?!

Why another zip-tool? What is this deterministic stuff?!

When we are talking about deterministic it means that the hash of the zip file won't change unless the contents of the
zip file changes.

This means only the content, no metadata. You can achieve this with zip, yes.

The problem that still remains is that the order is almost unpredictable and zip is very platform specific, so you will
end up with a bunch of crazy shell pipelines. And I am not event talking about windows at this point.

So this is where this tool comes in, it is intended to be a drop-in replacement for zip in your build process.

The use cases for this are primary:

- Zipping serverless code
- Backups or other files that get rsynced

### How reliable is it?

Of course, it is not as reliable as the battle-proven and billions of times executed zip.

Even though I am heavily relying on the go stdlib this software can of course have bugs. And you are welcome to report
them and help make this even more stable. Of course there will be tests to cover most use cases but at the end this is
still starting from scratch, so if you need advanced features or just dont feel comfortable about using this tool don't
do it!

### Differences between zip and deterministic-zip

Please see [docs/differences](./docs/differences)

### Why not just using another project out there?

As far as I know the following (GitHub) projects exist:

- [bboe/deterministic_zip (Python)](https://github.com/bboe/deterministic_zip)
    - You must list files explicitly
    - Changed order -> changed zip
    - You will need to install Python (no problem on Linux/Mac) and the package
- [bitgenics/deterministic-zip (NodeJS/JavaScript)](https://github.com/bitgenics/deterministic-zip#readme)
    - Support for globs and ignores order
    - You need to install node.js, the package, and it has no cli interface
- [orf/deterministic-zip (Rust)](https://github.com/orf/deterministic-zip)
    - has prebuilt binaries for all relevant platforms (and other can be built easily)
    - very basic, but you can customize compression (nice feature)

All in all they are just simply not what I needed. My favourite is Rust, because its just simply dropping
in a binary. Something that's very convenient especially when it comes to Docker builds.

The main problem that all these solutions share is that it in my opinion cool things like excluding patterns,
that I regularly use are simply not implemented, and i REALLY love glob patterns.
