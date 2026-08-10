package awssagemaker


// Properties for defining a `CfnMonitoringScheduleAlert`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringScheduleAlertProps := &CfnMonitoringScheduleAlertProps{
//   	DatapointsToAlert: jsii.Number(123),
//   	EvaluationPeriod: jsii.Number(123),
//   	MonitoringAlertName: jsii.String("monitoringAlertName"),
//   	MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedulealert.html
//
type CfnMonitoringScheduleAlertProps struct {
	// Within EvaluationPeriod, how many execution failures will raise an alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedulealert.html#cfn-sagemaker-monitoringschedulealert-datapointstoalert
	//
	DatapointsToAlert *float64 `field:"required" json:"datapointsToAlert" yaml:"datapointsToAlert"`
	// The number of most recent monitoring executions to consider when evaluating alert status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedulealert.html#cfn-sagemaker-monitoringschedulealert-evaluationperiod
	//
	EvaluationPeriod *float64 `field:"required" json:"evaluationPeriod" yaml:"evaluationPeriod"`
	// The name of the monitoring alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedulealert.html#cfn-sagemaker-monitoringschedulealert-monitoringalertname
	//
	MonitoringAlertName *string `field:"required" json:"monitoringAlertName" yaml:"monitoringAlertName"`
	// The name of the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedulealert.html#cfn-sagemaker-monitoringschedulealert-monitoringschedulename
	//
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
}

