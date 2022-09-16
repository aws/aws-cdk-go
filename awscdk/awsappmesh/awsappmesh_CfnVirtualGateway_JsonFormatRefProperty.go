package awsappmesh


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFormatRefProperty := &jsonFormatRefProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnVirtualGateway_JsonFormatRefProperty struct {
	// `CfnVirtualGateway.JsonFormatRefProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnVirtualGateway.JsonFormatRefProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

