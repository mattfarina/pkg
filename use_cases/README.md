# Use Cases

A package is a bundle of source code. But as developers, we often need
to know more about the package than what is immediately apparent from
the code. The challenge of any suitable package metadata standard is to
capture the information that is important to package users, but is not
overly onerous for package maintainers to collect and maintain.

This directory shares a set of use cases for tooling (e.g. Go package managers)
to solve that need specification information. Part of this is to identify the
information not available today to solve the use cases. [If you have any feedback
please feel free to share](https://github.com/mattfarina/pkg/issues).

Current use cases:

* [As a developer, I want to use a version or version range.](use_specific_version.md)
* [As a developer, I can pin a dependency tree to specific versions.](lock_version.md)
* [As a developer, I can retrieve private packages.](work_with_private_packages.md)
* [As a developer, I want to fork a Go package and use my fork.](working_with_forks.md)
* [As a developer, I want to have a single version of an import for my application and all dependencies.](single_import.md)
* [As a developer, store dependent packages in my VCS and update them.](managed_vendored_dependencies.md)
* [As a developer on a team, we can have the same setup to retrieve packages.](consistent_team_setup_with_private.md)
* [As a developer, I can easily scan my package tree to know the licenses in use.](license_scan.md)
* [As an application, I can find general information about a package.](application_information.md)
* [As a developer, I can contact a package's owner](contact_owners.md)
