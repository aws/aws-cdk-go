package assets


// Symlink follow mode.
// Deprecated: see `core.SymlinkFollowMode`
type FollowMode string

const (
	// Never follow symlinks.
	// Deprecated: see `core.SymlinkFollowMode`
	FollowMode_NEVER FollowMode = "NEVER"
	// Materialize all symlinks, whether they are internal or external to the source directory.
	// Deprecated: see `core.SymlinkFollowMode`
	FollowMode_ALWAYS FollowMode = "ALWAYS"
	// Only follows symlinks that are external to the source directory.
	// Deprecated: see `core.SymlinkFollowMode`
	FollowMode_EXTERNAL FollowMode = "EXTERNAL"
	// Forbids source from having any symlinks pointing outside of the source tree.
	//
	// This is the safest mode of operation as it ensures that copy operations
	// won't materialize files from the user's file system. Internal symlinks are
	// not followed.
	//
	// If the copy operation runs into an external symlink, it will fail.
	// Deprecated: see `core.SymlinkFollowMode`
	FollowMode_BLOCK_EXTERNAL FollowMode = "BLOCK_EXTERNAL"
)

