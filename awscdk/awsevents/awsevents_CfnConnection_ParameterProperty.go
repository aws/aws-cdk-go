package awsevents


// Additional query string parameter for the connection.
//
// You can include up to 100 additional query string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &parameterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	isValueSecret: jsii.Boolean(false),
//   }
//
type CfnConnection_ParameterProperty struct {
	// The key for a query string parameter.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value associated with the key for the query string parameter.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies whether the value is secret.
	IsValueSecret interface{} `field:"optional" json:"isValueSecret" yaml:"isValueSecret"`
}

