package awsiot


// The HTTP action header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpActionHeaderProperty := &httpActionHeaderProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_HttpActionHeaderProperty struct {
	// The HTTP header key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The HTTP header value.
	//
	// Substitution templates are supported.
	Value *string `field:"required" json:"value" yaml:"value"`
}

