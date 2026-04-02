package awsemr


// Holds CloudWatch log configuration settings and metadata that specify settings like log files to monitor and where to send them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudWatchLogConfigurationProperty := &CloudWatchLogConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	LogTypes: map[string][]*string{
//   		"logTypesKey": []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html
//
type CfnClusterPropsMixin_CloudWatchLogConfigurationProperty struct {
	// Specifies if CloudWatch logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the encryption key used to encrypt the logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The name of the CloudWatch log group where logs are published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The prefix of the log stream name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// A map of log types to file names for publishing logs to the standard output or standard error streams for CloudWatch.
	//
	// Valid log types include STEP_LOGS, SPARK_DRIVER, and SPARK_EXECUTOR. Valid file names for each type include STDOUT and STDERR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-logtypes
	//
	LogTypes interface{} `field:"optional" json:"logTypes" yaml:"logTypes"`
}

