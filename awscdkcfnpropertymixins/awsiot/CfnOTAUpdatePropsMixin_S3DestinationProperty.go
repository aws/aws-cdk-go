package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3DestinationProperty := &S3DestinationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-s3destination.html
//
type CfnOTAUpdatePropsMixin_S3DestinationProperty struct {
	// The S3 bucket that contains the updated firmware.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-s3destination.html#cfn-iot-otaupdate-s3destination-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-s3destination.html#cfn-iot-otaupdate-s3destination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

