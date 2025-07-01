package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsEncryptionProperty := &HlsEncryptionProperty{
//   	SpekeKeyProvider: &SpekeKeyProviderProperty{
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SystemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   		EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   	EncryptionMethod: jsii.String("encryptionMethod"),
//   	KeyRotationIntervalSeconds: jsii.Number(123),
//   	RepeatExtXKey: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html
//
type CfnOriginEndpoint_HlsEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html#cfn-mediapackage-originendpoint-hlsencryption-spekekeyprovider
	//
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html#cfn-mediapackage-originendpoint-hlsencryption-constantinitializationvector
	//
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// HLS encryption type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html#cfn-mediapackage-originendpoint-hlsencryption-encryptionmethod
	//
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html#cfn-mediapackage-originendpoint-hlsencryption-keyrotationintervalseconds
	//
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// Repeat the `EXT-X-KEY` directive for every media segment.
	//
	// This might result in an increase in client requests to the DRM server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlsencryption.html#cfn-mediapackage-originendpoint-hlsencryption-repeatextxkey
	//
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
}

