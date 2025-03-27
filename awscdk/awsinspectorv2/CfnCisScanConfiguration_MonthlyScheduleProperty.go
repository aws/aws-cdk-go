package awsinspectorv2


// A monthly schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monthlyScheduleProperty := &MonthlyScheduleProperty{
//   	Day: jsii.String("day"),
//   	StartTime: &TimeProperty{
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-monthlyschedule.html
//
type CfnCisScanConfiguration_MonthlyScheduleProperty struct {
	// The monthly schedule's day.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-monthlyschedule.html#cfn-inspectorv2-cisscanconfiguration-monthlyschedule-day
	//
	Day *string `field:"required" json:"day" yaml:"day"`
	// The monthly schedule's start time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-monthlyschedule.html#cfn-inspectorv2-cisscanconfiguration-monthlyschedule-starttime
	//
	StartTime interface{} `field:"required" json:"startTime" yaml:"startTime"`
}

