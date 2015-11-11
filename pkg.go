// Package pkg provides the ability to parse and write pkg.json and pkg.lock files.
package pkg

// Pkg provides the data structure backing pkg.json files
type Pkg struct {
	// Name is the name of an application or package. For example, main or
	// github.com/Masterminds/pkg.
	Name string `json:"name,omitempty"`

	// Image is an optional image such as a logo for a package or application.
	Image string `json:"image,omitempty"`

	// Description is a short description for a package. This description is
	// similar but different to a Go package description as it is for
	// marketing and presentation purposes rather than technical ones.
	Description string `json:"description,omitempty"`

	// Keyworkds is a small list of categories for a package.
	Keywords []string `json:"keywords,omitempty"`

	// Home is a url to a website for the package.
	Home string `json:"homepage,omitempty"`

	// License provides either a SPDX license or a path to a file containing
	// the license. For more information on SPDX see http://spdx.org/licenses/.
	// When more than one license an SPDX expression can be used.
	License string `json:"license,omitempty"`

	// Owners is an array of owners for a project. See the Owner type for
	// more detail. These can be one or more people, companies, or other
	// organizations.
	Owners []Owner `json:"owners,omitempty"`

	// Imports are the packages imported by this project. For more detail see
	// the Import type.
	Imports []Import `json:"imports"`

	// DevImports are the imports needed for testing or other development
	// activities. See the Import type for more detail.
	DevImports []Import `json:"devImports,omitempty"`
}

// Owner describes an owner of a package. This can be a person, company, or
// other organization. This is useful if someone needs to contact the
// owner of a package to address things like a security issue.
type Owner struct {

	// Name describes the name of an organization.
	Name string `json:"name,omitempty"`

	// Email is an email address to reach the owner at.
	Email string `json:"email,omitempty"`

	// Home is a url to a website for the owner.
	Home string `json:"homepage,omitempty"`
}

// Import represents an import required by this package.
type Import struct {

	// Name is the name of a package such as github.com/Masterminds/pkg
	Name string `json:"name"`

	// Version is a string representing a version. This can be a version control
	// system identifier or a semantic version. If used in a version control
	// system it can be a commit id, branch, or tag. If it's a semantic verion
	// it can be either a version or range. For more information on ranges see
	// the syntax at https://github.com/Masterminds/semver.
	Version string `json:"version,omitempty"`

	// Repo is the location of a repository to obtain a package. This is useful
	// for cases such as a fork or private package.
	Repo string `json:"repository,omitempty"`

	// Type represents the type of repo defined in Repo. This is for cases the
	// repo type cannot be detected from the Repo string. The typical options
	// will be git, svn, hg, and bzr representing the four most popular types
	// of version control systems.
	Type string `json:"type,omitempty"`
}
