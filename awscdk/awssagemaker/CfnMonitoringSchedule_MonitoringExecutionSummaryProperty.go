package awssagemaker


// Summary of information about the last monitoring job to run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringExecutionSummaryProperty := &MonitoringExecutionSummaryProperty{
//   	CreationTime: jsii.String("creationTime"),
//   	LastModifiedTime: jsii.String("lastModifiedTime"),
//   	MonitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   	MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   	ScheduledTime: jsii.String("scheduledTime"),
//
//   	// the properties below are optional
//   	EndpointName: jsii.String("endpointName"),
//   	FailureReason: jsii.String("failureReason"),
//   	ProcessingJobArn: jsii.String("processingJobArn"),
//   }
//
type CfnMonitoringSchedule_MonitoringExecutionSummaryProperty struct {
	// The time at which the monitoring job was created.
	CreationTime *string `field:"required" json:"creationTime" yaml:"creationTime"`
	// A timestamp that indicates the last time the monitoring job was modified.
	LastModifiedTime *string `field:"required" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// The status of the monitoring job.
	MonitoringExecutionStatus *string `field:"required" json:"monitoringExecutionStatus" yaml:"monitoringExecutionStatus"`
	// The name of the monitoring schedule.
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The time the monitoring job was scheduled.
	ScheduledTime *string `field:"required" json:"scheduledTime" yaml:"scheduledTime"`
	// The name of the endpoint used to run the monitoring job.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// The Amazon Resource Name (ARN) of the monitoring job.
	ProcessingJobArn *string `field:"optional" json:"processingJobArn" yaml:"processingJobArn"`
}

