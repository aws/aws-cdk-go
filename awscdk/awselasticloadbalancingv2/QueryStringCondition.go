package awselasticloadbalancingv2


// Properties for the key/value pair of the query string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringCondition := &QueryStringCondition{
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Key: jsii.String("key"),
//   }
//
type QueryStringCondition struct {
	// The query string value for the condition.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The query string key for the condition.
	// Default: - Any key can be matched.
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

