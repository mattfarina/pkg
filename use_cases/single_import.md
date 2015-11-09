# As a developer, I want to have a single version of an import for my application and all dependencies.

There are known issues with multiple versions of the same package within an
application or project. For example, there can be type mismatches (e.g., [an
example can be seen here](https://github.com/mattfarina/golang-broken-vendor)).
To avoid these situations it should be possible to only have a single version
of a dependency within the dependency tree.

## Solution
The recommendation is that a dependency tree be flattened to one instance of
an external package and it being set to a single version that is the most
appropriate for the codebase. That could either be in a single `vendor/` folder
for the `GO15VENDOREXPERIMENT` or the `GOPATH`.

While this is a recommendation the final solution will be left up to any package
managers or other tooling.
