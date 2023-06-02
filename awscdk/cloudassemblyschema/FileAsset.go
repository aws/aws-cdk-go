package cloudassemblyschema


// A file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAsset := &FileAsset{
//   	Destinations: map[string]fileDestination{
//   		"destinationsKey": &fileDestination{
//   			"bucketName": jsii.String("bucketName"),
//   			"objectKey": jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	Source: &FileSource{
//   		Executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		Packaging: awscdk.Cloud_assembly_schema.FileAssetPackaging_FILE,
//   		Path: jsii.String("path"),
//   	},
//   }
//
type FileAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*FileDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *FileSource `field:"required" json:"source" yaml:"source"`
}

