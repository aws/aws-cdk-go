package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerLogsProperty := &AuthorizerLogsProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   	S3: &S3Property{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-authorizerlogs.html
//
type CfnCluster_AuthorizerLogsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-authorizerlogs.html#cfn-msk-cluster-authorizerlogs-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-authorizerlogs.html#cfn-msk-cluster-authorizerlogs-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-authorizerlogs.html#cfn-msk-cluster-authorizerlogs-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

