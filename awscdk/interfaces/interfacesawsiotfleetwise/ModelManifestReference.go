package interfacesawsiotfleetwise


// A reference to a ModelManifest resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelManifestReference := &ModelManifestReference{
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	ModelManifestName: jsii.String("modelManifestName"),
//   }
//
type ModelManifestReference struct {
	// The ARN of the ModelManifest resource.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The Name of the ModelManifest resource.
	ModelManifestName *string `field:"required" json:"modelManifestName" yaml:"modelManifestName"`
}

