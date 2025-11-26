package previewawsmwaamixins


// The type of Apache Airflow logs to send to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	DagProcessingLogs: &ModuleLoggingConfigurationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		Enabled: jsii.Boolean(false),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   	SchedulerLogs: &ModuleLoggingConfigurationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		Enabled: jsii.Boolean(false),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   	TaskLogs: &ModuleLoggingConfigurationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		Enabled: jsii.Boolean(false),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   	WebserverLogs: &ModuleLoggingConfigurationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		Enabled: jsii.Boolean(false),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   	WorkerLogs: &ModuleLoggingConfigurationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		Enabled: jsii.Boolean(false),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html
//
type CfnEnvironmentPropsMixin_LoggingConfigurationProperty struct {
	// Defines the processing logs sent to CloudWatch Logs and the logging level to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-dagprocessinglogs
	//
	DagProcessingLogs interface{} `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// Defines the scheduler logs sent to CloudWatch Logs and the logging level to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-schedulerlogs
	//
	SchedulerLogs interface{} `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// Defines the task logs sent to CloudWatch Logs and the logging level to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-tasklogs
	//
	TaskLogs interface{} `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// Defines the web server logs sent to CloudWatch Logs and the logging level to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-webserverlogs
	//
	WebserverLogs interface{} `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// Defines the worker logs sent to CloudWatch Logs and the logging level to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-workerlogs
	//
	WorkerLogs interface{} `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

