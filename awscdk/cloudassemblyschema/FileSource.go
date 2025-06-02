package cloudassemblyschema


// Describe the source of a file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSource := &FileSource{
//   	Executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	Packaging: awscdk.Cloud_assembly_schema.FileAssetPackaging_FILE,
//   	Path: jsii.String("path"),
//   }
//
type FileSource struct {
	// External command which will produce the file asset to upload.
	// Default: - Exactly one of `executable` and `path` is required.
	//
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// Packaging method.
	//
	// Only allowed when `path` is specified.
	// Default: FILE.
	//
	Packaging FileAssetPackaging `field:"optional" json:"packaging" yaml:"packaging"`
	// The filesystem object to upload.
	//
	// This path is relative to the asset manifest location.
	// Default: - Exactly one of `executable` and `path` is required.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

