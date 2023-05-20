package awscdk


// Determines how symlinks are followed.
type SymlinkFollowMode string

const (
	// Never follow symlinks.
	SymlinkFollowMode_NEVER SymlinkFollowMode = "NEVER"
	// Materialize all symlinks, whether they are internal or external to the source directory.
	SymlinkFollowMode_ALWAYS SymlinkFollowMode = "ALWAYS"
	// Only follows symlinks that are external to the source directory.
	SymlinkFollowMode_EXTERNAL SymlinkFollowMode = "EXTERNAL"
	// Forbids source from having any symlinks pointing outside of the source tree.
	//
	// This is the safest mode of operation as it ensures that copy operations
	// won't materialize files from the user's file system. Internal symlinks are
	// not followed.
	//
	// If the copy operation runs into an external symlink, it will fail.
	SymlinkFollowMode_BLOCK_EXTERNAL SymlinkFollowMode = "BLOCK_EXTERNAL"
)

