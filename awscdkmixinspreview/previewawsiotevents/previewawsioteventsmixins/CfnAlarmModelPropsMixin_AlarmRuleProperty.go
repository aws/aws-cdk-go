package previewawsioteventsmixins


// Defines when your alarm is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   alarmRuleProperty := &AlarmRuleProperty{
//   	SimpleRule: &SimpleRuleProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		InputProperty: jsii.String("inputProperty"),
//   		Threshold: jsii.String("threshold"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmrule.html
//
type CfnAlarmModelPropsMixin_AlarmRuleProperty struct {
	// A rule that compares an input property value to a threshold value with a comparison operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmrule.html#cfn-iotevents-alarmmodel-alarmrule-simplerule
	//
	SimpleRule interface{} `field:"optional" json:"simpleRule" yaml:"simpleRule"`
}

