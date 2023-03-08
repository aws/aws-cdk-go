package awsapprunner


// Describes a key-value pair, which is a string-to-string mapping.
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
type CfnService_KeyValuePairProperty struct {
	// The key name string to map to a value.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value string to which the key name is mapped.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

