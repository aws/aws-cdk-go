package previewawsinspectorv2mixins


// The schedule the CIS scan configuration runs on.
//
// Each CIS scan configuration has exactly one type of schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var oneTime interface{}
//
//   scheduleProperty := &ScheduleProperty{
//   	Daily: &DailyScheduleProperty{
//   		StartTime: &TimeProperty{
//   			TimeOfDay: jsii.String("timeOfDay"),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   	},
//   	Monthly: &MonthlyScheduleProperty{
//   		Day: jsii.String("day"),
//   		StartTime: &TimeProperty{
//   			TimeOfDay: jsii.String("timeOfDay"),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   	},
//   	OneTime: oneTime,
//   	Weekly: &WeeklyScheduleProperty{
//   		Days: []*string{
//   			jsii.String("days"),
//   		},
//   		StartTime: &TimeProperty{
//   			TimeOfDay: jsii.String("timeOfDay"),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-schedule.html
//
type CfnCisScanConfigurationPropsMixin_ScheduleProperty struct {
	// A daily schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-schedule.html#cfn-inspectorv2-cisscanconfiguration-schedule-daily
	//
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// A monthly schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-schedule.html#cfn-inspectorv2-cisscanconfiguration-schedule-monthly
	//
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// A one time schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-schedule.html#cfn-inspectorv2-cisscanconfiguration-schedule-onetime
	//
	OneTime interface{} `field:"optional" json:"oneTime" yaml:"oneTime"`
	// A weekly schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-schedule.html#cfn-inspectorv2-cisscanconfiguration-schedule-weekly
	//
	Weekly interface{} `field:"optional" json:"weekly" yaml:"weekly"`
}

