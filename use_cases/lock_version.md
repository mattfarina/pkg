# As a developer, I can pin a dependency tree to specific versions.

When dependencies can have versions or ranges of supported versions (e.g.,
`>= 1.2.3, < 2.0.0`) and there is a dependency tree I want to be able to pin to
specific versions of packages for the entire tree. The specific version should
be a commit id or an alias to the specific commit.

## Notes

- Package managers for other languages (e.g., Composer/PHP, Cargo/Rust,
  NPM/Bower/JavaScript) provide this feature.

## Solution
Provide a `pkg.lock` file containing the specific versions for the dependency
tree. This file would be managed by tools rather than people. For example the following snippet is a subset of the `pkg.lock` file recording the specific version as a commit id,

```json
{
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
