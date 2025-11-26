package previewawsemrserverlessmixins


// The Amazon CloudWatch configuration for monitoring logs.
//
// You can configure your jobs to send log information to CloudWatch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchLoggingConfigurationProperty := &CloudWatchLoggingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	LogTypeMap: []interface{}{
//   		&LogTypeMapKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html
//
type CfnApplicationPropsMixin_CloudWatchLoggingConfigurationProperty struct {
	// Enables CloudWatch logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html#cfn-emrserverless-application-cloudwatchloggingconfiguration-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The AWS Key Management Service (KMS) key ARN to encrypt the logs that you store in CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html#cfn-emrserverless-application-cloudwatchloggingconfiguration-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The name of the log group in Amazon CloudWatch Logs where you want to publish your logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html#cfn-emrserverless-application-cloudwatchloggingconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Prefix for the CloudWatch log stream name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html#cfn-emrserverless-application-cloudwatchloggingconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// The specific log-streams which need to be uploaded to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-cloudwatchloggingconfiguration.html#cfn-emrserverless-application-cloudwatchloggingconfiguration-logtypemap
	//
	LogTypeMap interface{} `field:"optional" json:"logTypeMap" yaml:"logTypeMap"`
}

