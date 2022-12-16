package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafEncryptionProperty := &cmafEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	constantInitializationVector: jsii.String("constantInitializationVector"),
//   	encryptionMethod: jsii.String("encryptionMethod"),
//   	keyRotationIntervalSeconds: jsii.Number(123),
//   }
//
type CfnOriginEndpoint_CmafEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// `CfnOriginEndpoint.CmafEncryptionProperty.EncryptionMethod`.
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

