package cloudassemblyschema


// Packaging strategy for file assets.
type FileAssetPackaging string

const (
	// Upload the given path as a file.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	// The given path is a directory, zip it and upload.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

