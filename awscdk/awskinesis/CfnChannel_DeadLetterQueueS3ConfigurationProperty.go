package awskinesis


// Configuration of the S3 bucket used to capture records that cannot be delivered to the primary destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterQueueS3ConfigurationProperty := &DeadLetterQueueS3ConfigurationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//
//   	// the properties below are optional
//   	ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html
//
type CfnChannel_DeadLetterQueueS3ConfigurationProperty struct {
	// The ARN of the S3 bucket for storing failed records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The AWS account ID of the expected owner of the dead-letter queue S3 bucket.
	//
	// Used to verify bucket ownership before delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"required" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Optional S3 key prefix under which error records are organized.
	//
	// When omitted, the service uses the default 'kinesis-channel/errors/<channelName>/<channelId>/'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-deadletterqueues3configuration.html#cfn-kinesis-channel-deadletterqueues3configuration-erroroutputprefix
	//
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
}

