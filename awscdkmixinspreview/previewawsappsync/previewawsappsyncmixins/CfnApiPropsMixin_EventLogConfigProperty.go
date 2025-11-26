package previewawsappsyncmixins


// Describes the CloudWatch Logs configuration for the Event API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventLogConfigProperty := &EventLogConfigProperty{
//   	CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	LogLevel: jsii.String("logLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html
//
type CfnApiPropsMixin_EventLogConfigProperty struct {
	// The IAM service role that AWS AppSync assumes to publish CloudWatch Logs in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html#cfn-appsync-api-eventlogconfig-cloudwatchlogsrolearn
	//
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// The type of information to log for the Event API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html#cfn-appsync-api-eventlogconfig-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

