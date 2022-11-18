package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsEncryptionProperty := &hlsEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	constantInitializationVector: jsii.String("constantInitializationVector"),
//   	encryptionMethod: jsii.String("encryptionMethod"),
//   }
//
type CfnPackagingConfiguration_HlsEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.
	//
	// If you don't specify a constant initialization vector (IV), MediaPackage periodically rotates the IV.
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// HLS encryption type.
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
}

