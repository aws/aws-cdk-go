package awsappintegrations


// The metadata associated with the client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataProperty := &metadataProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnEventIntegration_MetadataProperty struct {
	// The key name.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

