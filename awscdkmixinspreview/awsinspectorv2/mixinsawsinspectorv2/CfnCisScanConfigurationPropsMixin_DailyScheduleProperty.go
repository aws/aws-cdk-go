package mixinsawsinspectorv2


// A daily schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dailyScheduleProperty := &DailyScheduleProperty{
//   	StartTime: &TimeProperty{
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-dailyschedule.html
//
type CfnCisScanConfigurationPropsMixin_DailyScheduleProperty struct {
	// The schedule start time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-dailyschedule.html#cfn-inspectorv2-cisscanconfiguration-dailyschedule-starttime
	//
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
}

