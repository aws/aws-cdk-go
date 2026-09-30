package awsmgn


// SSM command configuration for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorSsmCommandConfigProperty := &ConnectorSsmCommandConfigProperty{
//   	CloudWatchOutputEnabled: jsii.Boolean(false),
//   	S3OutputEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	OutputS3BucketName: jsii.String("outputS3BucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html
//
type CfnConnector_ConnectorSsmCommandConfigProperty struct {
	// Whether SSM command output is sent to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-cloudwatchoutputenabled
	//
	CloudWatchOutputEnabled interface{} `field:"required" json:"cloudWatchOutputEnabled" yaml:"cloudWatchOutputEnabled"`
	// Whether SSM command output is stored in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-s3outputenabled
	//
	S3OutputEnabled interface{} `field:"required" json:"s3OutputEnabled" yaml:"s3OutputEnabled"`
	// The CloudWatch Logs group name for SSM command output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-cloudwatchloggroupname
	//
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// The S3 bucket name for SSM command output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-outputs3bucketname
	//
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
}

