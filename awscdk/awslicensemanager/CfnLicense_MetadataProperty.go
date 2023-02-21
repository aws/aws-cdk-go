package awslicensemanager


// Describes key/value pairs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataProperty := &MetadataProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
type CfnLicense_MetadataProperty struct {
	// The key name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

