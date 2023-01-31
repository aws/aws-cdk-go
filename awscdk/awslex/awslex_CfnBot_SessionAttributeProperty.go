package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionAttributeProperty := &sessionAttributeProperty{
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnBot_SessionAttributeProperty struct {
	// `CfnBot.SessionAttributeProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnBot.SessionAttributeProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

