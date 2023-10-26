package awssagemaker


// Configuration details about the monitoring schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigProperty := &ScheduleConfigProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	DataAnalysisEndTime: jsii.String("dataAnalysisEndTime"),
//   	DataAnalysisStartTime: jsii.String("dataAnalysisStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-scheduleconfig.html
//
type CfnMonitoringSchedule_ScheduleConfigProperty struct {
	// A cron expression that describes details about the monitoring schedule.
	//
	// The supported cron expressions are:
	//
	// - If you want to set the job to start every hour, use the following:
	//
	// `Hourly: cron(0 * ? * * *)`
	// - If you want to start the job daily:
	//
	// `cron(0 [00-23] ? * * *)`
	// - If you want to run the job one time, immediately, use the following keyword:
	//
	// `NOW`
	//
	// For example, the following are valid cron expressions:
	//
	// - Daily at noon UTC: `cron(0 12 ? * * *)`
	// - Daily at midnight UTC: `cron(0 0 ? * * *)`
	//
	// To support running every 6, 12 hours, the following are also supported:
	//
	// `cron(0 [00-23]/[01-24] ? * * *)`
	//
	// For example, the following are valid cron expressions:
	//
	// - Every 12 hours, starting at 5pm UTC: `cron(0 17/12 ? * * *)`
	// - Every two hours starting at midnight: `cron(0 0/2 ? * * *)`
	//
	// > - Even though the cron expression is set to start at 5PM UTC, note that there could be a delay of 0-20 minutes from the actual requested time to run the execution.
	// > - We recommend that if you would like a daily schedule, you do not provide this parameter. Amazon SageMaker will pick a time for running every day.
	//
	// You can also specify the keyword `NOW` to run the monitoring job immediately, one time, without recurring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-scheduleconfig.html#cfn-sagemaker-monitoringschedule-scheduleconfig-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Sets the end time for a monitoring job window.
	//
	// Express this time as an offset to the times that you schedule your monitoring jobs to run. You schedule monitoring jobs with the `ScheduleExpression` parameter. Specify this offset in ISO 8601 duration format. For example, if you want to end the window one hour before the start of each monitoring job, you would specify: `"-PT1H"` .
	//
	// The end time that you specify must not follow the start time that you specify by more than 24 hours. You specify the start time with the `DataAnalysisStartTime` parameter.
	//
	// If you set `ScheduleExpression` to `NOW` , this parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-scheduleconfig.html#cfn-sagemaker-monitoringschedule-scheduleconfig-dataanalysisendtime
	//
	DataAnalysisEndTime *string `field:"optional" json:"dataAnalysisEndTime" yaml:"dataAnalysisEndTime"`
	// Sets the start time for a monitoring job window.
	//
	// Express this time as an offset to the times that you schedule your monitoring jobs to run. You schedule monitoring jobs with the `ScheduleExpression` parameter. Specify this offset in ISO 8601 duration format. For example, if you want to monitor the five hours of data in your dataset that precede the start of each monitoring job, you would specify: `"-PT5H"` .
	//
	// The start time that you specify must not precede the end time that you specify by more than 24 hours. You specify the end time with the `DataAnalysisEndTime` parameter.
	//
	// If you set `ScheduleExpression` to `NOW` , this parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-scheduleconfig.html#cfn-sagemaker-monitoringschedule-scheduleconfig-dataanalysisstarttime
	//
	DataAnalysisStartTime *string `field:"optional" json:"dataAnalysisStartTime" yaml:"dataAnalysisStartTime"`
}

