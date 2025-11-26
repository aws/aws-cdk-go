package previewawspcaconnectoradmixins


// Private key flags for v2 templates specify the client compatibility, if the private key can be exported, and if user input is required when using a private key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateKeyFlagsV2Property := &PrivateKeyFlagsV2Property{
//   	ClientVersion: jsii.String("clientVersion"),
//   	ExportableKey: jsii.Boolean(false),
//   	StrongKeyProtectionRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv2.html
//
type CfnTemplatePropsMixin_PrivateKeyFlagsV2Property struct {
	// Defines the minimum client compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv2.html#cfn-pcaconnectorad-template-privatekeyflagsv2-clientversion
	//
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// Allows the private key to be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv2.html#cfn-pcaconnectorad-template-privatekeyflagsv2-exportablekey
	//
	ExportableKey interface{} `field:"optional" json:"exportableKey" yaml:"exportableKey"`
	// Require user input when using the private key for enrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv2.html#cfn-pcaconnectorad-template-privatekeyflagsv2-strongkeyprotectionrequired
	//
	StrongKeyProtectionRequired interface{} `field:"optional" json:"strongKeyProtectionRequired" yaml:"strongKeyProtectionRequired"`
}

