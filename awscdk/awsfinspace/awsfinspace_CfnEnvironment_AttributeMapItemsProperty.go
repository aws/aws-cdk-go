package awsfinspace


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeMapItemsProperty := &AttributeMapItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CfnEnvironment_AttributeMapItemsProperty struct {
	// `CfnEnvironment.AttributeMapItemsProperty.Key`.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// `CfnEnvironment.AttributeMapItemsProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

