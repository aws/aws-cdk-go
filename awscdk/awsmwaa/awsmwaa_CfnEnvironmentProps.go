package awsmwaa


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var airflowConfigurationOptions interface{}
//   var tags interface{}
//
//   cfnEnvironmentProps := &cfnEnvironmentProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	airflowConfigurationOptions: airflowConfigurationOptions,
//   	airflowVersion: jsii.String("airflowVersion"),
//   	dagS3Path: jsii.String("dagS3Path"),
//   	environmentClass: jsii.String("environmentClass"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	kmsKey: jsii.String("kmsKey"),
//   	loggingConfiguration: &loggingConfigurationProperty{
//   		dagProcessingLogs: &moduleLoggingConfigurationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			enabled: jsii.Boolean(false),
//   			logLevel: jsii.String("logLevel"),
//   		},
//   		schedulerLogs: &moduleLoggingConfigurationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			enabled: jsii.Boolean(false),
//   			logLevel: jsii.String("logLevel"),
//   		},
//   		taskLogs: &moduleLoggingConfigurationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			enabled: jsii.Boolean(false),
//   			logLevel: jsii.String("logLevel"),
//   		},
//   		webserverLogs: &moduleLoggingConfigurationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			enabled: jsii.Boolean(false),
//   			logLevel: jsii.String("logLevel"),
//   		},
//   		workerLogs: &moduleLoggingConfigurationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			enabled: jsii.Boolean(false),
//   			logLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	maxWorkers: jsii.Number(123),
//   	minWorkers: jsii.Number(123),
//   	networkConfiguration: &networkConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	pluginsS3ObjectVersion: jsii.String("pluginsS3ObjectVersion"),
//   	pluginsS3Path: jsii.String("pluginsS3Path"),
//   	requirementsS3ObjectVersion: jsii.String("requirementsS3ObjectVersion"),
//   	requirementsS3Path: jsii.String("requirementsS3Path"),
//   	schedulers: jsii.Number(123),
//   	sourceBucketArn: jsii.String("sourceBucketArn"),
//   	tags: tags,
//   	webserverAccessMode: jsii.String("webserverAccessMode"),
//   	weeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   }
//
type CfnEnvironmentProps struct {
	// The name of your Amazon MWAA environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	//
	// For example, `core.default_timezone: utc` . To learn more, see [Apache Airflow configuration options](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html) .
	AirflowConfigurationOptions interface{} `field:"optional" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// The version of Apache Airflow to use for the environment.
	//
	// If no value is specified, defaults to the latest version. Valid values: `2.0.2` , `1.10.12` .
	AirflowVersion *string `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	//
	// For example, `dags` . To learn more, see [Adding or updating DAGs](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html) .
	DagS3Path *string `field:"optional" json:"dagS3Path" yaml:"dagS3Path"`
	// The environment class type.
	//
	// Valid values: `mw1.small` , `mw1.medium` , `mw1.large` . To learn more, see [Amazon MWAA environment class](https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html) .
	EnvironmentClass *string `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access AWS resources in your environment.
	//
	// For example, `arn:aws:iam::123456789:role/my-execution-role` . To learn more, see [Amazon MWAA Execution role](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) .
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	//
	// You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Apache Airflow logs being sent to CloudWatch Logs: `DagProcessingLogs` , `SchedulerLogs` , `TaskLogs` , `WebserverLogs` , `WorkerLogs` .
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The maximum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. For example, `20` . When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or the number you specify in `MinWorkers` .
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The minimum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the worker count you specify in the `MinWorkers` field. For example, `2` .
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	//
	// To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// The relative path to the `plugins.zip` file on your Amazon S3 bucket. For example, `plugins.zip` . To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// The version of the requirements.txt file on your Amazon S3 bucket. To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// The relative path to the `requirements.txt` file on your Amazon S3 bucket. For example, `requirements.txt` . To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// The number of schedulers that you want to run in your environment. Valid values:.
	//
	// - *v2.0.2* - Accepts between 2 to 5. Defaults to 2.
	// - *v1.10.12* - Accepts 1.
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.
	//
	// For example, `arn:aws:s3:::my-airflow-bucket-unique-name` . To learn more, see [Create an Amazon S3 bucket for Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html) .
	SourceBucketArn *string `field:"optional" json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// The key-value tag pairs associated to your environment.
	//
	// For example, `"Environment": "Staging"` . To learn more, see [Tagging](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Apache Airflow *Web server* access mode.
	//
	// To learn more, see [Apache Airflow access modes](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html) . Valid values: `PRIVATE_ONLY` or `PUBLIC_ONLY` .
	WebserverAccessMode *string `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM` .
	//
	// For example: `TUE:03:30` . You can specify a start time in 30 minute increments only. Supported input includes the following:
	//
	// - MON|TUE|WED|THU|FRI|SAT|SUN:([01]\\d|2[0-3]):(00|30).
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

