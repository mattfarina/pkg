package pkg

// Lock represents the data in a pkg.lock file.
type Lock struct {

	// Hash contains a sha256 of the pkg.json file to verify the lock file is
	// for the same pkg.json file and version.
	Hash string `json:"hash"`

	// Updated contains an RFC 3339 timestamp of when the pkg.lock file
	// was last updated.
	Updated string `json:"updated"`

	// LockedPackages is a slice of all the packages used within a project. That
	// includes the packages imported packages. The entire dependency tree.
	LockedPackages []LockedPackage `json:"imports"`
}

// LockedPackage represents a single locked package.
type LockedPackage struct {

	// Name is the name of the locked package.
	Name string `json:"name"`

	// Version is the specific version the package is locked to. This should not
	// be a branch. It should be a specific commit id.
	Version string `json:"version"`
}
