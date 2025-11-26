package previewawsworkspacesthinclientmixins


// Describes the maintenance window for a thin client device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   maintenanceWindowProperty := &MaintenanceWindowProperty{
//   	ApplyTimeOf: jsii.String("applyTimeOf"),
//   	DaysOfTheWeek: []*string{
//   		jsii.String("daysOfTheWeek"),
//   	},
//   	EndTimeHour: jsii.Number(123),
//   	EndTimeMinute: jsii.Number(123),
//   	StartTimeHour: jsii.Number(123),
//   	StartTimeMinute: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html
//
type CfnEnvironmentPropsMixin_MaintenanceWindowProperty struct {
	// The option to set the maintenance window during the device local time or Universal Coordinated Time (UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-applytimeof
	//
	ApplyTimeOf *string `field:"optional" json:"applyTimeOf" yaml:"applyTimeOf"`
	// The days of the week during which the maintenance window is open.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-daysoftheweek
	//
	DaysOfTheWeek *[]*string `field:"optional" json:"daysOfTheWeek" yaml:"daysOfTheWeek"`
	// The hour for the maintenance window end ( `00` - `23` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-endtimehour
	//
	EndTimeHour *float64 `field:"optional" json:"endTimeHour" yaml:"endTimeHour"`
	// The minutes for the maintenance window end ( `00` - `59` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-endtimeminute
	//
	EndTimeMinute *float64 `field:"optional" json:"endTimeMinute" yaml:"endTimeMinute"`
	// The hour for the maintenance window start ( `00` - `23` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-starttimehour
	//
	StartTimeHour *float64 `field:"optional" json:"startTimeHour" yaml:"startTimeHour"`
	// The minutes past the hour for the maintenance window start ( `00` - `59` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-starttimeminute
	//
	StartTimeMinute *float64 `field:"optional" json:"startTimeMinute" yaml:"startTimeMinute"`
	// An option to select the default or custom maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

