package awsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRotation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRotationProps := &CfnRotationProps{
//   	ContactIds: []*string{
//   		jsii.String("contactIds"),
//   	},
//   	Name: jsii.String("name"),
//   	Recurrence: &RecurrenceSettingsProperty{
//   		NumberOfOnCalls: jsii.Number(123),
//   		RecurrenceMultiplier: jsii.Number(123),
//
//   		// the properties below are optional
//   		DailySettings: []*string{
//   			jsii.String("dailySettings"),
//   		},
//   		MonthlySettings: []interface{}{
//   			&MonthlySettingProperty{
//   				DayOfMonth: jsii.Number(123),
//   				HandOffTime: jsii.String("handOffTime"),
//   			},
//   		},
//   		ShiftCoverages: []interface{}{
//   			&ShiftCoverageProperty{
//   				CoverageTimes: []interface{}{
//   					&CoverageTimeProperty{
//   						EndTime: jsii.String("endTime"),
//   						StartTime: jsii.String("startTime"),
//   					},
//   				},
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   			},
//   		},
//   		WeeklySettings: []interface{}{
//   			&WeeklySettingProperty{
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   				HandOffTime: jsii.String("handOffTime"),
//   			},
//   		},
//   	},
//   	StartTime: jsii.String("startTime"),
//   	TimeZoneId: jsii.String("timeZoneId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html
//
type CfnRotationProps struct {
	// The Amazon Resource Names (ARNs) of the contacts to add to the rotation.
	//
	// > Only the `PERSONAL` contact type is supported. The contact types `ESCALATION` and `ONCALL_SCHEDULE` are not supported for this operation.
	//
	// The order in which you list the contacts is their shift order in the rotation schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-contactids
	//
	ContactIds *[]*string `field:"required" json:"contactIds" yaml:"contactIds"`
	// The name for the rotation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information about the rule that specifies when shift team members rotate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-recurrence
	//
	Recurrence interface{} `field:"required" json:"recurrence" yaml:"recurrence"`
	// The date and time the rotation goes into effect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-starttime
	//
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// The time zone to base the rotation’s activity on, in Internet Assigned Numbers Authority (IANA) format.
	//
	// For example: "America/Los_Angeles", "UTC", or "Asia/Seoul". For more information, see the [Time Zone Database](https://docs.aws.amazon.com/https://www.iana.org/time-zones) on the IANA website.
	//
	// > Designators for time zones that don’t support Daylight Savings Time rules, such as Pacific Standard Time (PST), are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-timezoneid
	//
	TimeZoneId *string `field:"required" json:"timeZoneId" yaml:"timeZoneId"`
	// Optional metadata to assign to the rotation.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For more information, see [Tagging Incident Manager resources](https://docs.aws.amazon.com/incident-manager/latest/userguide/tagging.html) in the *Incident Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html#cfn-ssmcontacts-rotation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

