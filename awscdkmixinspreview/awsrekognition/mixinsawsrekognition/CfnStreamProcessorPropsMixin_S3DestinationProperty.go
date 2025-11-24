package mixinsawsrekognition


// The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.
//
// These results include the name of the stream processor resource, the session ID of the stream processing session, and labeled timestamps and bounding boxes for detected labels. For more information, see [S3Destination](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_S3Destination) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DestinationProperty := &S3DestinationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html
//
type CfnStreamProcessorPropsMixin_S3DestinationProperty struct {
	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name of a stream processor's exports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html#cfn-rekognition-streamprocessor-s3destination-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Describes the destination Amazon Simple Storage Service (Amazon S3) object keys of a stream processor's exports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html#cfn-rekognition-streamprocessor-s3destination-objectkeyprefix
	//
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

