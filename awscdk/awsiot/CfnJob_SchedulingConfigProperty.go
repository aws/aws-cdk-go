package awsiot


// Specifies the date and time that a job will begin the rollout of the job document to all devices in the target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schedulingConfigProperty := &SchedulingConfigProperty{
//   	EndBehavior: jsii.String("endBehavior"),
//   	EndTime: jsii.String("endTime"),
//   	MaintenanceWindows: []interface{}{
//   		&MaintenanceWindowProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-schedulingconfig.html
//
type CfnJob_SchedulingConfigProperty struct {
	// Specifies the end behavior for all job executions after a job reaches the selected endTime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-schedulingconfig.html#cfn-iot-job-schedulingconfig-endbehavior
	//
	EndBehavior *string `field:"optional" json:"endBehavior" yaml:"endBehavior"`
	// The time a job will stop rollout of the job document to all devices in the target group for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-schedulingconfig.html#cfn-iot-job-schedulingconfig-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-schedulingconfig.html#cfn-iot-job-schedulingconfig-maintenancewindows
	//
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
	// The time a job will begin rollout of the job document to all devices in the target group for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-schedulingconfig.html#cfn-iot-job-schedulingconfig-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

