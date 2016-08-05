# As a project maintainer, I can pick and choose which vendored dependencies are committed to VCS

There is no hard and fast rule on whether a given project should choose to commit its dependencies to VCS or not.

* Some dependencies (eg. kubernetes) may be extremely large
* Some dependencies may recursively include parts of the current project
* Legal reasons may preclude committing third-party code with certain licenses
