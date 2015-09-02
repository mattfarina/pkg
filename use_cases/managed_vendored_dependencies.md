# As a developer, store dependent packages in my VCS and update them.

The 3rd party packages in my VCS need to be tied to versions of the package and
I need a process to update to newer versions of the dependency.

## Notes
- Vendoring, as described on StackOverflow (with a Ruby specific case):
  > Vendoring is the moving of all 3rd party items such as plugins, gems and even rails into the /vendor directory.
- Vendor packages are those from a 3rd party. The 3rd party is the vendor.

## Solution
The `version` property associated with any import allows a version to be set. A
tool can retrieve a package, set its version, put it in the proper place
(e.g. the `vendor/` directory), and remove any VCS data (e.g. The `.git` directory).
This would allow for packages to be checkout out and updated while being
vendored.

The example above is simply one process to handle vendoring. Having package
information, including the `version`, allows for numerous process to locally
store and update packages.

```json
{
    "name": "My App",
    "imports": [
        {
            "name": "github.com/ACME/foo",
            "version": "1.0.0"
        }
    ]
}
```
