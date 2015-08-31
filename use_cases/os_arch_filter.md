# As a developer, I want to filter the installed packages by OS and Architecture.

There are cases where a package is specific to an operating system or architecture.
For example, `github.com/opencontainers/runc` is a Linux only package at the
moment. Attempting to install it on other systems can cause errors.

## Notes
- There are examples of windows only packages causing issues in Mac development
  environments as well.

## Solution
The specification allows one to specify supported operating systems and architectures
for each package.

```json
{
    "name": "github.com/example/app",
    "imports": [
        {
            "name": "github.com/example/foo",
            "os": ["windows"],
            "arch": ["386"]
        }
    ]
}
```
