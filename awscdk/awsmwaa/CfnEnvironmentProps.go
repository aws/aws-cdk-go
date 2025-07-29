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
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AirflowConfigurationOptions: airflowConfigurationOptions,
//   	AirflowVersion: jsii.String("airflowVersion"),
//   	DagS3Path: jsii.String("dagS3Path"),
//   	EndpointManagement: jsii.String("endpointManagement"),
//   	EnvironmentClass: jsii.String("environmentClass"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	KmsKey: jsii.String("kmsKey"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		DagProcessingLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		SchedulerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		TaskLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WebserverLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WorkerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	MaxWebservers: jsii.Number(123),
//   	MaxWorkers: jsii.Number(123),
//   	MinWebservers: jsii.Number(123),
//   	MinWorkers: jsii.Number(123),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	PluginsS3ObjectVersion: jsii.String("pluginsS3ObjectVersion"),
//   	PluginsS3Path: jsii.String("pluginsS3Path"),
//   	RequirementsS3ObjectVersion: jsii.String("requirementsS3ObjectVersion"),
//   	RequirementsS3Path: jsii.String("requirementsS3Path"),
//   	Schedulers: jsii.Number(123),
//   	SourceBucketArn: jsii.String("sourceBucketArn"),
//   	StartupScriptS3ObjectVersion: jsii.String("startupScriptS3ObjectVersion"),
//   	StartupScriptS3Path: jsii.String("startupScriptS3Path"),
//   	Tags: tags,
//   	WebserverAccessMode: jsii.String("webserverAccessMode"),
//   	WeeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   	WorkerReplacementStrategy: jsii.String("workerReplacementStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html
//
type CfnEnvironmentProps struct {
	// The name of your Amazon MWAA environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	//
	// For example, `core.default_timezone: utc` . To learn more, see [Apache Airflow configuration options](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowconfigurationoptions
	//
	AirflowConfigurationOptions interface{} `field:"optional" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// The version of Apache Airflow to use for the environment.
	//
	// If no value is specified, defaults to the latest version.
	//
	// If you specify a newer version number for an existing environment, the version update requires some service interruption before taking effect.
	//
	// *Allowed Values* : `1.10.12` | `2.0.2` | `2.2.2` | `2.4.3` | `2.5.1` | `2.6.3` | `2.7.2` | `2.8.1` | `2.9.2` | `2.10.1` (latest)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowversion
	//
	AirflowVersion *string `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	//
	// For example, `dags` . To learn more, see [Adding or updating DAGs](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-dags3path
	//
	DagS3Path *string `field:"optional" json:"dagS3Path" yaml:"dagS3Path"`
	// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
	//
	// If set to `SERVICE` , Amazon MWAA will create and manage the required VPC endpoints in your VPC. If set to `CUSTOMER` , you must create, and manage, the VPC endpoints in your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-endpointmanagement
	//
	EndpointManagement *string `field:"optional" json:"endpointManagement" yaml:"endpointManagement"`
	// The environment class type.
	//
	// Valid values: `mw1.micro` , `mw1.small` , `mw1.medium` , `mw1.large` , `mw1.1large` , and `mw1.2large` . To learn more, see [Amazon MWAA environment class](https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-environmentclass
	//
	EnvironmentClass *string `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access AWS resources in your environment.
	//
	// For example, `arn:aws:iam::123456789:role/my-execution-role` . To learn more, see [Amazon MWAA Execution role](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	//
	// You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Apache Airflow logs being sent to CloudWatch Logs: `DagProcessingLogs` , `SchedulerLogs` , `TaskLogs` , `WebserverLogs` , `WorkerLogs` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The maximum number of web servers that you want to run in your environment.
	//
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for `MaxWebservers` when you interact with your Apache Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. For example, in scenarios where your workload requires network calls to the Apache Airflow REST API with a high transaction-per-second (TPS) rate, Amazon MWAA will increase the number of web servers up to the number set in `MaxWebserers` . As TPS rates decrease Amazon MWAA disposes of the additional web servers, and scales down to the number set in `MinxWebserers` .
	//
	// Valid values: For environments larger than mw1.micro, accepts values from `2` to `5` . Defaults to `2` for all environment sizes except mw1.micro, which defaults to `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-maxwebservers
	//
	MaxWebservers *float64 `field:"optional" json:"maxWebservers" yaml:"maxWebservers"`
	// The maximum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. For example, `20` . When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or the number you specify in `MinWorkers` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-maxworkers
	//
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The minimum number of web servers that you want to run in your environment.
	//
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for `MaxWebservers` when you interact with your Apache Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. As the transaction-per-second rate, and the network load, decrease, Amazon MWAA disposes of the additional web servers, and scales down to the number set in `MinxWebserers` .
	//
	// Valid values: For environments larger than mw1.micro, accepts values from `2` to `5` . Defaults to `2` for all environment sizes except mw1.micro, which defaults to `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-minwebservers
	//
	MinWebservers *float64 `field:"optional" json:"minWebservers" yaml:"minWebservers"`
	// The minimum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the worker count you specify in the `MinWorkers` field. For example, `2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-minworkers
	//
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	//
	// To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3objectversion
	//
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// The relative path to the `plugins.zip` file on your Amazon S3 bucket. For example, `plugins.zip` . To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3path
	//
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// The version of the requirements.txt file on your Amazon S3 bucket. To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3objectversion
	//
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// The relative path to the `requirements.txt` file on your Amazon S3 bucket. For example, `requirements.txt` . To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3path
	//
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// The number of schedulers that you want to run in your environment. Valid values:.
	//
	// - *v2* - For environments larger than mw1.micro, accepts values from 2 to 5. Defaults to 2 for all environment sizes except mw1.micro, which defaults to 1.
	// - *v1* - Accepts 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-schedulers
	//
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.
	//
	// For example, `arn:aws:s3:::my-airflow-bucket-unique-name` . To learn more, see [Create an Amazon S3 bucket for Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-sourcebucketarn
	//
	SourceBucketArn *string `field:"optional" json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// The version of the startup shell script in your Amazon S3 bucket.
	//
	// You must specify the [version ID](https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html) that Amazon S3 assigns to the file every time you update the script.
	//
	// Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than 1,024 bytes long. The following is an example:
	//
	// `3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo`
	//
	// For more information, see [Using a startup script](https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-startupscripts3objectversion
	//
	StartupScriptS3ObjectVersion *string `field:"optional" json:"startupScriptS3ObjectVersion" yaml:"startupScriptS3ObjectVersion"`
	// The relative path to the startup shell script in your Amazon S3 bucket. For example, `s3://mwaa-environment/startup.sh` .
	//
	// Amazon MWAA runs the script as your environment starts, and before running the Apache Airflow process. You can use this script to install dependencies, modify Apache Airflow configuration options, and set environment variables. For more information, see [Using a startup script](https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-startupscripts3path
	//
	StartupScriptS3Path *string `field:"optional" json:"startupScriptS3Path" yaml:"startupScriptS3Path"`
	// The key-value tag pairs associated to your environment. For example, `"Environment": "Staging"` . To learn more, see [Tagging](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	//
	// If you specify new tags for an existing environment, the update requires service interruption before taking effect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Apache Airflow *Web server* access mode.
	//
	// To learn more, see [Apache Airflow access modes](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html) . Valid values: `PRIVATE_ONLY` or `PUBLIC_ONLY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-webserveraccessmode
	//
	WebserverAccessMode *string `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM` .
	//
	// For example: `TUE:03:30` . You can specify a start time in 30 minute increments only. Supported input includes the following:
	//
	// - MON|TUE|WED|THU|FRI|SAT|SUN:([01]\\d|2[0-3]):(00|30).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-weeklymaintenancewindowstart
	//
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
	// The worker replacement strategy to use when updating the environment.
	//
	// Valid values: `FORCED`, `GRACEFUL`. FORCED means Apache Airflow workers will be stopped and replaced without waiting for tasks to complete before an update. GRACEFUL means Apache Airflow workers will be able to complete running tasks for up to 12 hours during an update before being stopped and replaced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-workerreplacementstrategy
	//
	WorkerReplacementStrategy *string `field:"optional" json:"workerReplacementStrategy" yaml:"workerReplacementStrategy"`
}

