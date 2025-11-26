package previewawsiotmixins


// The Amazon S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-s3location.html
//
type CfnSoftwarePackageVersionPropsMixin_S3LocationProperty struct {
	// The S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-s3location.html#cfn-iot-softwarepackageversion-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-s3location.html#cfn-iot-softwarepackageversion-s3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-s3location.html#cfn-iot-softwarepackageversion-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

