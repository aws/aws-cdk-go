package previewawsathenamixins


// Configuration settings for delivering logs to Amazon CloudWatch log groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchLoggingConfigurationProperty := &CloudWatchLoggingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	LogGroup: jsii.String("logGroup"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	LogTypes: map[string][]*string{
//   		"logTypesKey": []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-cloudwatchloggingconfiguration.html
//
type CfnWorkGroupPropsMixin_CloudWatchLoggingConfigurationProperty struct {
	// Enables CloudWatch logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-cloudwatchloggingconfiguration.html#cfn-athena-workgroup-cloudwatchloggingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the log group in Amazon CloudWatch Logs where you want to publish your logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-cloudwatchloggingconfiguration.html#cfn-athena-workgroup-cloudwatchloggingconfiguration-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Prefix for the CloudWatch log stream name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-cloudwatchloggingconfiguration.html#cfn-athena-workgroup-cloudwatchloggingconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// The types of logs that you want to publish to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-cloudwatchloggingconfiguration.html#cfn-athena-workgroup-cloudwatchloggingconfiguration-logtypes
	//
	LogTypes interface{} `field:"optional" json:"logTypes" yaml:"logTypes"`
}

