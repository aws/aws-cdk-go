package awsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataProperty := &MetadataProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CfnEventIntegration_MetadataProperty struct {
	// `CfnEventIntegration.MetadataProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnEventIntegration.MetadataProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

