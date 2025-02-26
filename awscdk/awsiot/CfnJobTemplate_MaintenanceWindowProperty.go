package awsiot


// An optional configuration within the `SchedulingConfig` to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowProperty := &MaintenanceWindowProperty{
//   	DurationInMinutes: jsii.Number(123),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-maintenancewindow.html
//
type CfnJobTemplate_MaintenanceWindowProperty struct {
	// Displays the duration of the next maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-maintenancewindow.html#cfn-iot-jobtemplate-maintenancewindow-durationinminutes
	//
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// Displays the start time of the next maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-maintenancewindow.html#cfn-iot-jobtemplate-maintenancewindow-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

