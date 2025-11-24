package mixinsawsinspectorv2


// A weekly schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   weeklyScheduleProperty := &WeeklyScheduleProperty{
//   	Days: []*string{
//   		jsii.String("days"),
//   	},
//   	StartTime: &TimeProperty{
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-weeklyschedule.html
//
type CfnCisScanConfigurationPropsMixin_WeeklyScheduleProperty struct {
	// The weekly schedule's days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-weeklyschedule.html#cfn-inspectorv2-cisscanconfiguration-weeklyschedule-days
	//
	Days *[]*string `field:"optional" json:"days" yaml:"days"`
	// The weekly schedule's start time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-weeklyschedule.html#cfn-inspectorv2-cisscanconfiguration-weeklyschedule-starttime
	//
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
}

