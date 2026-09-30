package awsmgn


// SSM command configuration for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorSsmCommandConfigProperty := &ConnectorSsmCommandConfigProperty{
//   	CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	CloudWatchOutputEnabled: jsii.Boolean(false),
//   	OutputS3BucketName: jsii.String("outputS3BucketName"),
//   	S3OutputEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html
//
type CfnConnectorPropsMixin_ConnectorSsmCommandConfigProperty struct {
	// The CloudWatch Logs group name for SSM command output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-cloudwatchloggroupname
	//
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Whether SSM command output is sent to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-cloudwatchoutputenabled
	//
	CloudWatchOutputEnabled interface{} `field:"optional" json:"cloudWatchOutputEnabled" yaml:"cloudWatchOutputEnabled"`
	// The S3 bucket name for SSM command output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-outputs3bucketname
	//
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// Whether SSM command output is stored in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-connector-connectorssmcommandconfig.html#cfn-mgn-connector-connectorssmcommandconfig-s3outputenabled
	//
	S3OutputEnabled interface{} `field:"optional" json:"s3OutputEnabled" yaml:"s3OutputEnabled"`
}

