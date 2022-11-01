package cloudassemblyschema


// Packaging strategy for file assets.
// Experimental.
type FileAssetPackaging string

const (
	// Upload the given path as a file.
	// Experimental.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	// The given path is a directory, zip it and upload.
	// Experimental.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

