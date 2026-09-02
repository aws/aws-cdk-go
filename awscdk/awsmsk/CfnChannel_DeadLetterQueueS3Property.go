package awsmsk


// Dead letter queue S3 configuration of the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterQueueS3Property := &DeadLetterQueueS3Property{
//   	BucketArn: jsii.String("bucketArn"),
//   	ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//
//   	// the properties below are optional
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-deadletterqueues3.html
//
type CfnChannel_DeadLetterQueueS3Property struct {
	// The ARN of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-deadletterqueues3.html#cfn-msk-channel-deadletterqueues3-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The error output prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-deadletterqueues3.html#cfn-msk-channel-deadletterqueues3-erroroutputprefix
	//
	ErrorOutputPrefix *string `field:"required" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// Optional 12-digit AWS account ID expected to own the dead-letter S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-deadletterqueues3.html#cfn-msk-channel-deadletterqueues3-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
}

