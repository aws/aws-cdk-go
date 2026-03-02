package previewawscloudwatchmixins


// The rule for the mute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleProperty := &RuleProperty{
//   	Schedule: &ScheduleProperty{
//   		Duration: jsii.String("duration"),
//   		Expression: jsii.String("expression"),
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-rule.html
//
type CfnAlarmMuteRulePropsMixin_RuleProperty struct {
	// Schedule for the mute to be active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-rule.html#cfn-cloudwatch-alarmmuterule-rule-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

