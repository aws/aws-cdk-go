package awsworkspacesthinclient


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowProperty := &MaintenanceWindowProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ApplyTimeOf: jsii.String("applyTimeOf"),
//   	DaysOfTheWeek: []*string{
//   		jsii.String("daysOfTheWeek"),
//   	},
//   	EndTimeHour: jsii.Number(123),
//   	EndTimeMinute: jsii.Number(123),
//   	StartTimeHour: jsii.Number(123),
//   	StartTimeMinute: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html
//
type CfnEnvironment_MaintenanceWindowProperty struct {
	// The type of maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The desired time zone maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-applytimeof
	//
	ApplyTimeOf *string `field:"optional" json:"applyTimeOf" yaml:"applyTimeOf"`
	// The date of maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-daysoftheweek
	//
	DaysOfTheWeek *[]*string `field:"optional" json:"daysOfTheWeek" yaml:"daysOfTheWeek"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-endtimehour
	//
	EndTimeHour *float64 `field:"optional" json:"endTimeHour" yaml:"endTimeHour"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-endtimeminute
	//
	EndTimeMinute *float64 `field:"optional" json:"endTimeMinute" yaml:"endTimeMinute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-starttimehour
	//
	StartTimeHour *float64 `field:"optional" json:"startTimeHour" yaml:"startTimeHour"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesthinclient-environment-maintenancewindow.html#cfn-workspacesthinclient-environment-maintenancewindow-starttimeminute
	//
	StartTimeMinute *float64 `field:"optional" json:"startTimeMinute" yaml:"startTimeMinute"`
}

