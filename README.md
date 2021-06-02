deterministic-zip
===

> Work In Progress

Simple (almost drop-in) replacement for zip.

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
