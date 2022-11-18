package awsmediapackage


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionContractConfigurationProperty := &encryptionContractConfigurationProperty{
//   	presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   	presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   }
//
type CfnPackagingConfiguration_EncryptionContractConfigurationProperty struct {
	// `CfnPackagingConfiguration.EncryptionContractConfigurationProperty.PresetSpeke20Audio`.
	PresetSpeke20Audio *string `field:"required" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// `CfnPackagingConfiguration.EncryptionContractConfigurationProperty.PresetSpeke20Video`.
	PresetSpeke20Video *string `field:"required" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

