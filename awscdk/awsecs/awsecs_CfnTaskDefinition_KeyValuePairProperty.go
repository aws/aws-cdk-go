package awsecs


// The `KeyValuePair` property specifies a key-value pair object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValuePairProperty := &keyValuePairProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_KeyValuePairProperty struct {
	// The name of the key-value pair.
	//
	// For environment variables, this is the name of the environment variable.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair.
	//
	// For environment variables, this is the value of the environment variable.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

