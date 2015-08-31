# As a developer, I can retrieve private packages.

A dependent package for my application is a private package, not currently
publicly available. I need a way to import a package and let the VCS use
my SSH keys when accessing the repository.

## Notes
- An example could be a package `github.com/example/private` where this is a
  GitHub private repo used internally for a company.
- Other tools, such as [Composer](https://getcomposer.org/doc/04-schema.md#repositories)
  (for PHP) and [npm](https://docs.npmjs.com/files/package.json#repository) (for Node.js),
  provide a means to specify a repository. This does it a little different as it
  can be associated with each import.

## Solution
The json configuration allows one to specify the the repository. For example,

```json
{
    "name": "github.com/example/app",
    "imports": [
        {
            "name": "github.com/example/private",
            "repository": "git@github.com:example/private.git",
            "type": "git"
        }
    ]
}
```

Here the the dependent package will be checked out from `git@github.com:example/private.git`
to the location `github.com/example/private`. SSH keys will be properly used.
