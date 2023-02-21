package awscloudfront


// Contains information about the Amazon Kinesis data stream where you are sending real-time log data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamConfigProperty := &kinesisStreamConfigProperty{
//   	roleArn: jsii.String("roleArn"),
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnRealtimeLogConfig_KinesisStreamConfigProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that CloudFront can use to send real-time log data to your Kinesis data stream.
	//
	// For more information the IAM role, see [Real-time log configuration IAM role](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) in the *Amazon CloudFront Developer Guide* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending real-time log data.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

