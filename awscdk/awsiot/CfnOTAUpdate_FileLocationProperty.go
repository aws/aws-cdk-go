package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileLocationProperty := &FileLocationProperty{
//   	S3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   	Stream: &StreamProperty{
//   		FileId: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-filelocation.html
//
type CfnOTAUpdate_FileLocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-filelocation.html#cfn-iot-otaupdate-filelocation-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-filelocation.html#cfn-iot-otaupdate-filelocation-stream
	//
	Stream interface{} `field:"optional" json:"stream" yaml:"stream"`
}

