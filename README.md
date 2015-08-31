# Go Package Specification

This outlines a proposed specification for documenting the metadata associated
with a Go package described in a `pkg.json` file.

For example,
```json
{
    "name": "github.com/mattfarina/pkg",
    "description": "A package definition for Go.",
    "keywords": ["vendor", "package"],
    "imports": [
        {
            "name": "github.com/Masterminds/vcs"
        },
        {
            "name": "github.com/Masterminds/cookoo",
            "version": "v1.0.0"
        }
    ]
}
```

## Where Inspiration Came From

The ideas outlined in this setup come from other package managers that solved
the problems Go is working on. The ideas are are from lessons learned elsewhere.
These other package managers include [npm](http://npmjs.com/) (for Node.js),
[Bower](http://bower.io) (browser side JavaScript), [Composer](https://getcomposer.org)
(for PHP), [Crates](https://crates.io) (for Rust), [Bundler](http://bundler.io)
and [Ruby Gems](https://rubygems.org), and others. This is a well worn path
we can learn from.

## Package Information

Some information describes the package itself. For example,
```json
{
    "name": "github.com/Masterminds/glide",
    "description": "A package manager for Go.",
    "keywords": ["vendor", "package", "VCS"],
    "homepage": "https://example.com",
    "license": "MIT",
    "authors": [
        {
            "name": "Matt Butcher",
            "email": "something@example.com",
            "homepage": "http://technosophos.com"
        },
        {
            "name": "Matt Farina",
            "email": "other@example.com",
            "homepage": "http://mattfarina.com"
        }
    ],
    "imports": []
}
```

Each of these properties is useful for tools and people who need to know about
a project. For example, the `homepage` property for GB would be `http://getgb.io`
which tells you something about the project not in the code but useful to others.
Another example is the `license` property. It notes an [SPDX](http://spdx.org/)
compatible license or a path to a file with the license. This is useful for tools
that need to look over the packages in an application to determine license
information.

All of these properties are useful for applications that need to display or
otherwise work with information about an application or package that's not
part of the source for making the application run.

## Importing Dependent Packages

Go packages will often import dependent packages. While working with those
dependencies there are cases where you're developing one at the same time
as your application, cases for working with forks, and other ways packages need
to be handled. To accommodate those details, the `import` property is an array
of objects containing details about a package. For example,

```json
{
    "name": "github.com/Masterminds/glide",
    "imports": [
        {
            "name": "github.com/Masterminds/cookoo",
            "version": "1.2.0",
            "repository": "git@github.com:Masterminds/cookoo.git",
            "type": "git",
            "updateTime": "2015-04-09T05:06:17-08:00",
            "arch": ["windows"],
            "os": ["386"]
        }
    ]
}
```

Aside from the `name` field, all the others are optional. If just the `name` is
specified, the package dependent package would be retrieved and the version used
would be the latest on the default branch. This is the same version `go get`
retrieves today. With other properties are used the `name` is the place where
the package will be located in the directory tree.

The `version` field allows you to set a branch, tag, commit id, or Semantic
Version constraint (e.g, `">= 1.0, < 1.4"`).

The `repository` and `type` information are useful for a couple reasons including:

1. The package is in a private repository and you need SSH keys to access it.
2. Forks. For example what if the `repository` were `git@github.com:mattfarina/cookoo.git`.
3. When you want to checkout a dependency in a manner you can commit and push
   changes while working on the main project.

The `updateTime` is useful to know when the package was last updated. This would
be typically updated and otherwise managed by the tooling.

The `arch` and `os` optional properties allow you to filter the operating systems
and architectures to install a package with.

## Vendoring?

Vendoring, where the code from the dependent package is stored in the parent
projects VCS, can be used with this setup but isn't required. A tool could update
vendored code or install the proper code and version in each environment. The way
dependencies are stored is left as an exercise for the developer.

## Work In Progress

This is currently a work in progress.

Todo:

- [ ] Write use cases.
- [ ] Create package to read `pkg.json` files.
- [ ] Create validator for `pkg.json` files.
- [ ] Add BSD license details.
