# As a developer, I want to use a specific version of a package.

I need to be able to choose the version or version constraint of a package.

## Notes
- Common methods for versions are VCS branches, tags, and commit ids.
- Common methods for version constraints are [Semantic Version](http://semver.org/)
  ranges. For example `">= 1.0, < 1.4"`.
- There are Go packages today that support constraints checking such as
  `github.com/hashicorp/go-version`.
- Other language tools that support versions and version constraints this way
  include npm (for Node.js), Composer (for PHP), and Bundler (for Ruby).
  Cargo (for Rust) does not appear to support constraints. 

## Solution
Each import specified can have a version associated with it. This can be a branch,
tag, commit id, or Semantic Version constraint (e.g, `">= 1.0, < 1.4"`). From
this information a version can be selected to use.

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
