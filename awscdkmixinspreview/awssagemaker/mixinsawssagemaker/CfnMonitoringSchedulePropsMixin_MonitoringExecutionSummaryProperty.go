package mixinsawssagemaker


// Summary of information about the last monitoring job to run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringExecutionSummaryProperty := &MonitoringExecutionSummaryProperty{
//   	CreationTime: jsii.String("creationTime"),
//   	EndpointName: jsii.String("endpointName"),
//   	FailureReason: jsii.String("failureReason"),
//   	LastModifiedTime: jsii.String("lastModifiedTime"),
//   	MonitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   	MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   	ProcessingJobArn: jsii.String("processingJobArn"),
//   	ScheduledTime: jsii.String("scheduledTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html
//
type CfnMonitoringSchedulePropsMixin_MonitoringExecutionSummaryProperty struct {
	// The time at which the monitoring job was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-creationtime
	//
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The name of the endpoint used to run the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-failurereason
	//
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// A timestamp that indicates the last time the monitoring job was modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-lastmodifiedtime
	//
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// The status of the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-monitoringexecutionstatus
	//
	MonitoringExecutionStatus *string `field:"optional" json:"monitoringExecutionStatus" yaml:"monitoringExecutionStatus"`
	// The name of the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-monitoringschedulename
	//
	MonitoringScheduleName *string `field:"optional" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The Amazon Resource Name (ARN) of the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-processingjobarn
	//
	ProcessingJobArn *string `field:"optional" json:"processingJobArn" yaml:"processingJobArn"`
	// The time the monitoring job was scheduled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.html#cfn-sagemaker-monitoringschedule-monitoringexecutionsummary-scheduledtime
	//
	ScheduledTime *string `field:"optional" json:"scheduledTime" yaml:"scheduledTime"`
}

