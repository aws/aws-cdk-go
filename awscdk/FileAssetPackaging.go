package awscdk


// Packaging modes for file assets.
// Experimental.
type FileAssetPackaging string

const (
	// The asset source path points to a directory, which should be archived using zip and and then uploaded to Amazon S3.
	// Experimental.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
	// The asset source path points to a single file, which should be uploaded to Amazon S3.
	// Experimental.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
)

