package previewawsmediapackagemixins


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsEncryptionProperty := &HlsEncryptionProperty{
//   	ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   	EncryptionMethod: jsii.String("encryptionMethod"),
//   	SpekeKeyProvider: &SpekeKeyProviderProperty{
//   		EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		SystemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsencryption.html
//
type CfnPackagingConfigurationPropsMixin_HlsEncryptionProperty struct {
	// A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.
	//
	// If you don't specify a constant initialization vector (IV), AWS Elemental MediaPackage periodically rotates the IV.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsencryption.html#cfn-mediapackage-packagingconfiguration-hlsencryption-constantinitializationvector
	//
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// HLS encryption type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsencryption.html#cfn-mediapackage-packagingconfiguration-hlsencryption-encryptionmethod
	//
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Parameters for the SPEKE key provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlsencryption.html#cfn-mediapackage-packagingconfiguration-hlsencryption-spekekeyprovider
	//
	SpekeKeyProvider interface{} `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

