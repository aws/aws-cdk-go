package awsmwaa


// The type of Apache Airflow logs to send to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	dagProcessingLogs: &moduleLoggingConfigurationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		enabled: jsii.Boolean(false),
//   		logLevel: jsii.String("logLevel"),
//   	},
//   	schedulerLogs: &moduleLoggingConfigurationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		enabled: jsii.Boolean(false),
//   		logLevel: jsii.String("logLevel"),
//   	},
//   	taskLogs: &moduleLoggingConfigurationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		enabled: jsii.Boolean(false),
//   		logLevel: jsii.String("logLevel"),
//   	},
//   	webserverLogs: &moduleLoggingConfigurationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		enabled: jsii.Boolean(false),
//   		logLevel: jsii.String("logLevel"),
//   	},
//   	workerLogs: &moduleLoggingConfigurationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		enabled: jsii.Boolean(false),
//   		logLevel: jsii.String("logLevel"),
//   	},
//   }
//
type CfnEnvironment_LoggingConfigurationProperty struct {
	// Defines the processing logs sent to CloudWatch Logs and the logging level to send.
	DagProcessingLogs interface{} `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// Defines the scheduler logs sent to CloudWatch Logs and the logging level to send.
	SchedulerLogs interface{} `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// Defines the task logs sent to CloudWatch Logs and the logging level to send.
	TaskLogs interface{} `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// Defines the web server logs sent to CloudWatch Logs and the logging level to send.
	WebserverLogs interface{} `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// Defines the worker logs sent to CloudWatch Logs and the logging level to send.
	WorkerLogs interface{} `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

