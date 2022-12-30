package awsappmesh


// An object that represents the key value pairs for the JSON.
//
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
	// The specified key for the JSON.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The specified value for the JSON.
	Value *string `field:"required" json:"value" yaml:"value"`
}

