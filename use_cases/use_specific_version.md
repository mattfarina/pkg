# As a developer, I want to use a specific version of a package.

For any given package upon which my application depends, I need to be able to
specify, with a high degree of reliability, which version or versions of
that package work with my application.

While it is a common idiom in Go to not make breaking API chanages
between major version changes (e.g. 1.X to 2.x), three things are clear:

1. Even small "non-breaking" changes have a tendency to break things.
2. Many large and established projects (such as Docker) simply ignore
   this practice.
3. Developers cannot and should not rely upon external parties to retain
   backwards compatibility.

I need to be able to choose the version or version constraint of a package. If a
package has dependent packages I need those properly fetched and their versions
set.

## Notes
- Common methods for versions are VCS branches, tags, and commit ids.
- Common methods for version constraints are [Semantic Version](http://semver.org/)
  ranges. For example `">= 1.0, < 1.4"`.
- There are Go packages today that support constraints checking such as
  `github.com/hashicorp/go-version`.
- Other language tools that support versions and version constraints this way
  include npm (for Node.js), Composer (for PHP), and Bundler (for Ruby).
  Cargo (for Rust) does not appear to support constraints.
- This is also a standard feature of package managers for operating
  systems (including their `-dev` packages), including Apt and RPM.

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
