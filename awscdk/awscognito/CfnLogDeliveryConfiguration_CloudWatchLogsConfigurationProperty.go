package awscognito


// Configuration for the CloudWatch log group destination of user pool detailed activity logging, or of user activity log export with advanced security features.
//
// This data type is a request parameter of [SetLogDeliveryConfiguration](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetLogDeliveryConfiguration.html) and a response parameter of [GetLogDeliveryConfiguration](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_GetLogDeliveryConfiguration.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsConfigurationProperty := &CloudWatchLogsConfigurationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-cloudwatchlogsconfiguration.html
//
type CfnLogDeliveryConfiguration_CloudWatchLogsConfigurationProperty struct {
	// The Amazon Resource Name (arn) of a CloudWatch Logs log group where your user pool sends logs.
	//
	// The log group must not be encrypted with AWS Key Management Service and must be in the same AWS account as your user pool.
	//
	// To send logs to log groups with a resource policy of a size greater than 5120 characters, configure a log group with a path that starts with `/aws/vendedlogs` . For more information, see [Enabling logging from certain AWS services](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-cloudwatchlogsconfiguration.html#cfn-cognito-logdeliveryconfiguration-cloudwatchlogsconfiguration-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

