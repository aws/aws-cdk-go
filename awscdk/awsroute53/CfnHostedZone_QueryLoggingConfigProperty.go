package awsroute53


// A complex type that contains information about a configuration for DNS query logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryLoggingConfigProperty := &QueryLoggingConfigProperty{
//   	CloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html
//
type CfnHostedZone_QueryLoggingConfigProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html#cfn-route53-hostedzone-queryloggingconfig-cloudwatchlogsloggrouparn
	//
	CloudWatchLogsLogGroupArn *string `field:"required" json:"cloudWatchLogsLogGroupArn" yaml:"cloudWatchLogsLogGroupArn"`
}

