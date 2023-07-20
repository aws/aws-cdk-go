package awsmwaa


// Defines the type of logs to send for the Apache Airflow log type (e.g. `DagProcessingLogs` ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   moduleLoggingConfigurationProperty := &ModuleLoggingConfigurationProperty{
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	Enabled: jsii.Boolean(false),
//   	LogLevel: jsii.String("logLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html
//
type CfnEnvironment_ModuleLoggingConfigurationProperty struct {
	// The ARN of the CloudWatch Logs log group for each type of Apache Airflow log type that you have enabled.
	//
	// > `CloudWatchLogGroupArn` is available only as a return value, accessible when specified as an attribute in the [`Fn:GetAtt`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#aws-resource-mwaa-environment-return-values) intrinsic function. Any value you provide for `CloudWatchLogGroupArn` is discarded by Amazon MWAA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Indicates whether to enable the Apache Airflow log type (e.g. `DagProcessingLogs` ) in CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Defines the Apache Airflow logs to send for the log type (e.g. `DagProcessingLogs` ) to CloudWatch Logs. Valid values: `CRITICAL` , `ERROR` , `WARNING` , `INFO` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

