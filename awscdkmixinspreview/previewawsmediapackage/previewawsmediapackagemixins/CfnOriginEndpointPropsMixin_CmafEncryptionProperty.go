package previewawsmediapackagemixins


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafEncryptionProperty := &CmafEncryptionProperty{
//   	ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   	EncryptionMethod: jsii.String("encryptionMethod"),
//   	KeyRotationIntervalSeconds: jsii.Number(123),
//   	SpekeKeyProvider: &SpekeKeyProviderProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SystemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafencryption.html
//
type CfnOriginEndpointPropsMixin_CmafEncryptionProperty struct {
	// An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks.
	//
	// If you don't specify a value, then AWS Elemental MediaPackage creates the constant initialization vector (IV).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafencryption.html#cfn-mediapackage-originendpoint-cmafencryption-constantinitializationvector
	//
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The encryption method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafencryption.html#cfn-mediapackage-originendpoint-cmafencryption-encryptionmethod
	//
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafencryption.html#cfn-mediapackage-originendpoint-cmafencryption-keyrotationintervalseconds
	//
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// Parameters for the SPEKE key provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafencryption.html#cfn-mediapackage-originendpoint-cmafencryption-spekekeyprovider
	//
	SpekeKeyProvider interface{} `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

