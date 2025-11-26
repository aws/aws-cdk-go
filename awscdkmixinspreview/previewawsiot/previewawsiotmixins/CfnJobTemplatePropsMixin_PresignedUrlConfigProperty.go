package previewawsiotmixins


// Configuration for pre-signed S3 URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   presignedUrlConfigProperty := &PresignedUrlConfigProperty{
//   	ExpiresInSec: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html
//
type CfnJobTemplatePropsMixin_PresignedUrlConfigProperty struct {
	// How long (in seconds) pre-signed URLs are valid.
	//
	// Valid values are 60 - 3600, the default value is 3600 seconds. Pre-signed URLs are generated when Jobs receives an MQTT request for the job document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html#cfn-iot-jobtemplate-presignedurlconfig-expiresinsec
	//
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
	// The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored.
	//
	// The role must also grant permission for IoT to download the files.
	//
	// > For information about addressing the confused deputy problem, see [cross-service confused deputy prevention](https://docs.aws.amazon.com/iot/latest/developerguide/cross-service-confused-deputy-prevention.html) in the *AWS IoT Core developer guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html#cfn-iot-jobtemplate-presignedurlconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

