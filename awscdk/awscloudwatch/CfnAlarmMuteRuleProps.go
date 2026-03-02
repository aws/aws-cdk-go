package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAlarmMuteRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmMuteRuleProps := &CfnAlarmMuteRuleProps{
//   	Rule: &RuleProperty{
//   		Schedule: &ScheduleProperty{
//   			Duration: jsii.String("duration"),
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExpireDate: jsii.String("expireDate"),
//   	MuteTargets: &MuteTargetsProperty{
//   		AlarmNames: []*string{
//   			jsii.String("alarmNames"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	StartDate: jsii.String("startDate"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html
//
type CfnAlarmMuteRuleProps struct {
	// The rule for the mute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The description of the AlarmMuteRule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date, with the same timezone offset as "ScheduleTimezone" after which the alarm mute rule will be expired.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-expiredate
	//
	ExpireDate *string `field:"optional" json:"expireDate" yaml:"expireDate"`
	// Targets to be muted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-mutetargets
	//
	MuteTargets interface{} `field:"optional" json:"muteTargets" yaml:"muteTargets"`
	// The name of the AlarmMuteRule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The date, with the same timezone offset as "ScheduleTimezone", after which the alarm mute rule will become active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-startdate
	//
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarmmuterule.html#cfn-cloudwatch-alarmmuterule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

