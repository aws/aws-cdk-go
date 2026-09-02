package awsiot


// Represents a file to stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamFileProperty := &StreamFileProperty{
//   	FileId: jsii.Number(123),
//   	S3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-streamfile.html
//
type CfnStreamPropsMixin_StreamFileProperty struct {
	// The file ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-streamfile.html#cfn-iot-stream-streamfile-fileid
	//
	FileId *float64 `field:"optional" json:"fileId" yaml:"fileId"`
	// The location of the file in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-stream-streamfile.html#cfn-iot-stream-streamfile-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

