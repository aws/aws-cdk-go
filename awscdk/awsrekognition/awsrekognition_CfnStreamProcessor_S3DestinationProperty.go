package awsrekognition


// The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.
//
// These results include the name of the stream processor resource, the session ID of the stream processing session, and labeled timestamps and bounding boxes for detected labels. For more information, see [S3Destination](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_S3Destination) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &s3DestinationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   }
//
type CfnStreamProcessor_S3DestinationProperty struct {
	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name of a stream processor's exports.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Describes the destination Amazon Simple Storage Service (Amazon S3) object keys of a stream processor's exports.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

