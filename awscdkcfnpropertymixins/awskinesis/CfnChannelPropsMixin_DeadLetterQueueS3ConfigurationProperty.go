package awskinesis


// Configuration of the S3 bucket used to capture records that cannot be delivered to the primary destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   deadLetterQueueS3ConfigurationProperty := &DeadLetterQueueS3ConfigurationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html
//
type CfnChannelPropsMixin_DeadLetterQueueS3ConfigurationProperty struct {
	// The ARN of the S3 bucket for storing failed records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// Optional S3 key prefix under which error records are organized.
	//
	// When omitted, the service uses the default 'kinesis-channel/errors/<channelName>/<channelId>/'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-erroroutputprefix
	//
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The AWS account ID of the expected owner of the dead-letter queue S3 bucket.
	//
	// Used to verify bucket ownership before delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
}

