# As a developer, I can easily scan my package tree to know the licenses in use.

When I have a package or application that uses other packages I need a way to
automatically look through all the licenses in use among all the packages in my
application.

This commonly occurs for:

* Sarbanes-Oxley reviews
* Merger/Acquisition due diligence
* Legal complaince

## Notes
- [SPDX](http://spdx.org/) provides a list of common open source licenses.

## Solution
The license key allows one to specify either the license file location or the
name of a license on the SPDX list. This allows for easy identification of
license without needing to parse text and take into account differences in files.

```json
{
    "name": "My App",
    "license": "MIT"
}
```
