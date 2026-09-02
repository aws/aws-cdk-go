package awsiot


// The location of the file in S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-s3location.html
//
type CfnStreamPropsMixin_S3LocationProperty struct {
	// The S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-s3location.html#cfn-iot-stream-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-s3location.html#cfn-iot-stream-s3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 bucket version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-s3location.html#cfn-iot-stream-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

