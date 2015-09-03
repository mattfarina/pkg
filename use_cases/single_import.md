# As a developer, I want to have a single version of an import for my application and all dependencies.

Two or more dependencies of an application can have the same dependency. The
application may even share that dependency. There are cases, for example a
database driver, where only one package should be used across everything.

## Notes
- Dependency versions come into play and any tooling will need to make this
  easy to work with.
- Looking for a dependency starts in a package’s immediate `vendor/` directory,
  starts going up the package tree looking in `vendor/` directories, then looks
  in the `$GOPATH`, and finally looks in the `$GOROOT`.
- A simple use case is the Deis project that has multiple packages and
  sub-packages importing Docker. They only need this to happen one time.
- Certain frameworks in Go's core, such as the `database/sql` package,
  are prone to issues when unflattened dependencies result in loading
  two or more versions of the same driver library.

## Solution
There are two `flatten` properties. One at the package level that notes that all
packages should only be installed at the top level (e.g., the application
`vendor/` directory).

The `flatten` property on individual imports can be used to override the global
`flatten` property. This notes that this particular import should be at the top
level.

```json
{
    "name": "github.com/example/app",
    "flatten": true,
    "imports": [
        {
            "name": "github.com/ACME/foo",
            "flatten": true
        }
    ]
}
```
