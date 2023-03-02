package awsroute53


// A complex type that contains information about a configuration for DNS query logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryLoggingConfigProperty := &queryLoggingConfigProperty{
//   	cloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   }
//
type CfnHostedZone_QueryLoggingConfigProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
	CloudWatchLogsLogGroupArn *string `field:"required" json:"cloudWatchLogsLogGroupArn" yaml:"cloudWatchLogsLogGroupArn"`
}

