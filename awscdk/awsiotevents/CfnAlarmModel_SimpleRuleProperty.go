package awsiotevents


// A rule that compares an input property value to a threshold value with a comparison operator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleRuleProperty := &SimpleRuleProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	InputProperty: jsii.String("inputProperty"),
//   	Threshold: jsii.String("threshold"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-simplerule.html
//
type CfnAlarmModel_SimpleRuleProperty struct {
	// The comparison operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-simplerule.html#cfn-iotevents-alarmmodel-simplerule-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value on the left side of the comparison operator.
	//
	// You can specify an AWS IoT Events input attribute as an input property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-simplerule.html#cfn-iotevents-alarmmodel-simplerule-inputproperty
	//
	InputProperty *string `field:"required" json:"inputProperty" yaml:"inputProperty"`
	// The value on the right side of the comparison operator.
	//
	// You can enter a number or specify an AWS IoT Events input attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-simplerule.html#cfn-iotevents-alarmmodel-simplerule-threshold
	//
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

