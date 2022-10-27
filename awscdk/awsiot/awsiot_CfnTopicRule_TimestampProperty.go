package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestampProperty := &timestampProperty{
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	unit: jsii.String("unit"),
//   }
//
type CfnTopicRule_TimestampProperty struct {
	// `CfnTopicRule.TimestampProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
	// `CfnTopicRule.TimestampProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

