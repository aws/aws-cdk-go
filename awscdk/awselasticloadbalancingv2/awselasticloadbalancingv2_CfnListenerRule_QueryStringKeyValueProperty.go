package awselasticloadbalancingv2


// Information about a key/value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringKeyValueProperty := &queryStringKeyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnListenerRule_QueryStringKeyValueProperty struct {
	// The key.
	//
	// You can omit the key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

