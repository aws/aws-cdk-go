package awsroute53


// Properties for defining a `CfnQueryLoggingConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueryLoggingConfigProps := &CfnQueryLoggingConfigProps{
//   	CloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-queryloggingconfig.html
//
type CfnQueryLoggingConfigProps struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group in us-east-1 that Amazon Route 53 publishes query logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-queryloggingconfig.html#cfn-route53-queryloggingconfig-cloudwatchlogsloggrouparn
	//
	CloudWatchLogsLogGroupArn *string `field:"required" json:"cloudWatchLogsLogGroupArn" yaml:"cloudWatchLogsLogGroupArn"`
	// The ID of the public hosted zone that Amazon Route 53 logs queries for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-queryloggingconfig.html#cfn-route53-queryloggingconfig-hostedzoneid
	//
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

