package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizerLogsProperty := &AuthorizerLogsProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		DeliveryStream: jsii.String("deliveryStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	S3: &S3Property{
//   		Bucket: jsii.String("bucket"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-authorizerlogs.html
//
type CfnClusterPropsMixin_AuthorizerLogsProperty struct {
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

