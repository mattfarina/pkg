# As a developer, I want to fork a Go package and use my fork.

There are times developers want to fork a package, make a change, and use the
fork from the original import path. For example, I've forked `github.com/ACME/foo`
to `github.com/myaccount/foo` and made some changes. But, I still want to use
`github.com/ACME/foo` in my imports, to avoid import rewriting.

## Notes
- A simple example of this is forking an application, wanting to use a change,
  and submitting a pull request to the original at the same time. For
  some short period of time, a fork may be used while waiting for a
  stable upstream release.
- This fixes the import rewriting problem sometimes seen in Go.

## Solution
In this case the `repository` can be set to an alternative location, such as my
fork. The package will be checked out to the location listed by `name`.

```json
{
    "name": "github.com/example/app",
    "imports": [
        {
            "name": "github.com/ACME/foo",
            "repository": "https://github.com/myaccount/foo",
            "type": "git"
        }
    ]
}
```
