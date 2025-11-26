package previewawsmediapackagemixins


// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that provides encryption keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spekeKeyProviderProperty := &SpekeKeyProviderProperty{
//   	EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   		PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   		PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SystemIds: []*string{
//   		jsii.String("systemIds"),
//   	},
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.html
//
type CfnPackagingConfigurationPropsMixin_SpekeKeyProviderProperty struct {
	// Use `encryptionContractConfiguration` to configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.html#cfn-mediapackage-packagingconfiguration-spekekeyprovider-encryptioncontractconfiguration
	//
	EncryptionContractConfiguration interface{} `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
	// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.html#cfn-mediapackage-packagingconfiguration-spekekeyprovider-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.html#cfn-mediapackage-packagingconfiguration-spekekeyprovider-systemids
	//
	SystemIds *[]*string `field:"optional" json:"systemIds" yaml:"systemIds"`
	// URL for the key provider's key retrieval API endpoint.
	//
	// Must start with https://.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.html#cfn-mediapackage-packagingconfiguration-spekekeyprovider-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

