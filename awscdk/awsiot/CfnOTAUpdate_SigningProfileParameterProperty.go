package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingProfileParameterProperty := &SigningProfileParameterProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   	Platform: jsii.String("platform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-signingprofileparameter.html
//
type CfnOTAUpdate_SigningProfileParameterProperty struct {
	// Certificate ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-signingprofileparameter.html#cfn-iot-otaupdate-signingprofileparameter-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The location of the code-signing certificate on your device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-signingprofileparameter.html#cfn-iot-otaupdate-signingprofileparameter-certificatepathondevice
	//
	CertificatePathOnDevice *string `field:"optional" json:"certificatePathOnDevice" yaml:"certificatePathOnDevice"`
	// The hardware platform of your device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-signingprofileparameter.html#cfn-iot-otaupdate-signingprofileparameter-platform
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

