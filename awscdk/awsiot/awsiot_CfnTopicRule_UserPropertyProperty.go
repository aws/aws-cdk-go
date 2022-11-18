package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPropertyProperty := &userPropertyProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_UserPropertyProperty struct {
	// `CfnTopicRule.UserPropertyProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnTopicRule.UserPropertyProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

