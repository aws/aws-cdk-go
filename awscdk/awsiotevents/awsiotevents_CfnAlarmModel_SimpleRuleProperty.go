package awsiotevents


// A rule that compares an input property value to a threshold value with a comparison operator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleRuleProperty := &simpleRuleProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	inputProperty: jsii.String("inputProperty"),
//   	threshold: jsii.String("threshold"),
//   }
//
type CfnAlarmModel_SimpleRuleProperty struct {
	// The comparison operator.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value on the left side of the comparison operator.
	//
	// You can specify an AWS IoT Events input attribute as an input property.
	InputProperty *string `field:"required" json:"inputProperty" yaml:"inputProperty"`
	// The value on the right side of the comparison operator.
	//
	// You can enter a number or specify an AWS IoT Events input attribute.
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

