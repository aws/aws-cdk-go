package previewawsschedulermixins


// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flexibleTimeWindowProperty := &FlexibleTimeWindowProperty{
//   	MaximumWindowInMinutes: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-flexibletimewindow.html
//
type CfnSchedulePropsMixin_FlexibleTimeWindowProperty struct {
	// The maximum time window during which a schedule can be invoked.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1440`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-flexibletimewindow.html#cfn-scheduler-schedule-flexibletimewindow-maximumwindowinminutes
	//
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
	// Determines whether the schedule is invoked within a flexible time window.
	//
	// You must use quotation marks when you specify this value in your JSON or YAML template.
	//
	// *Allowed Values* : `"OFF"` | `"FLEXIBLE"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-flexibletimewindow.html#cfn-scheduler-schedule-flexibletimewindow-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

