# Go Package Specification

This outlines a proposed manner for documenting the metadata associated
with a Go package. The current examples use json but that's not a requirement.
Instead, the goal is to solve the [use cases](use_cases). Proposed changes
are welcome. Ideally, there would be one format for documenting the information
that numerous tools (e.g. package managers) could all use and share.

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

Of course, much can also be learned from the venerable operating system
package managers like RPM and dpkg, which have long been used to manage
development libraries.

## Use Cases

At this point the focus is on use cases that need to be solved. The requirements are
needed to understand what needs to be built and solved. The use cases are currently being
documented in the [use cases](use_cases) sub-directory.

## Package Information

Some information describes the package itself. For example,
```json
{
    "name": "github.com/Masterminds/glide",
    "image": "logo.png",
    "description": "A package manager for Go.",
    "keywords": ["vendor", "package", "VCS"],
    "homepage": "https://example.com",
    "license": "MIT",
    "owners": [
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

## Version Locking

When you can support version ranges it's useful to lock to a specific version
of each package in the dependency tree. This is supported through the use of a
`pkg.lock` file containing the locking information.

```json
{
    "hash": "2cb908fb4479bec8ed4fb8b6e719207fcf11d97e",
    "updated": "2006-01-02T15:04:05Z07:00",
    "imports":[
        {
            "name": "github.com/Masterminds/semver",
            "version": "6333b7bd29aad1d79898ff568fd90a8aa533ae82"
        },
        {
            "name": "github.com/Masterminds/vcs",
            "version": "eaee272c8fa4514e1572e182faecff5be20e792a"
        }
    ]
}
```

The `hash` property is a sha256 hash of the `pkg.json` file. This is used to
make sure the `pkg.lock` file is in sync with the `pkg.json` file. When they are
not in sync the `pkg.lock` information should not be used until the `pkg.lock`
is regenerated with the based on the current `pkg.json` file. `updated` provides
a timestamp in [RFC 3339](http://tools.ietf.org/html/rfc3339) format for the
last time the `pkg.lock` file was generated.

The `imports` property is a list of all the packages in the entire dependency
tree along with the specific version, as opposed to version range, the project
is locked to.

## Vendoring?

Vendoring, where the code from the dependent package is stored in the parent
projects VCS, can be used with this setup but isn't required. A tool could update
vendored code or install the proper code and version in each environment. The way
dependencies are stored is left as an exercise for the developer.

## Work In Progress

This is currently a work in progress.

Todo:

- [x] Write initial use cases.
- [x] Get feedback on use cases and refine.
- [ ] Create package to read `pkg.json` files.
- [ ] Update package to create and update `pkg.json` files.
- [ ] Create validator for `pkg.json` files.
- [ ] Add BSD license details.
- [ ] Decide where this should live.
