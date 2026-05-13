package awscdk


// Determines how symlinks are followed.
type SymlinkFollowMode string

const (
	// Copy symlinks themselves (do not follow, preserve the symlink contents).
	//
	// If any of the symlinks point outside the source directory and the Cloud
	// Assembly gets moved to a different computer, the asset may end up pointing
	// to files that don't exist on that other computer.
	SymlinkFollowMode_NEVER SymlinkFollowMode = "NEVER"
	// Copy all the files the symlinks point to (always follow).
	//
	// This reads target files even if the symlinks point outside the source
	// directory.  Uses more disk space but is guaranteed to end up with a
	// complete asset directory.
	SymlinkFollowMode_ALWAYS SymlinkFollowMode = "ALWAYS"
	// Materialize symlinks pointing outside, leave symlinks pointing inside.
	//
	// If the symlink points to a file outside the source, copy the target file.
	// Otherwise copy the symlink itself.
	//
	// This produces a complete asset bundle, while saving space.
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

