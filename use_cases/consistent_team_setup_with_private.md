# As a developer on a team, we can have the same setup to retrieve packages.

All the packages on the project we, the team, are using need to be setup in a
consistent manner on each development machine and in the testing and deployment
system.

Note, some of the packages we're using are private and we need to consistently
install each of them.

## Solution
The packages defined in the `pkg.json` file will be consistently installed in
each environment. That includes those that define a repository, as outlined
in the [private repository use case](work_with_private_packages.md).
