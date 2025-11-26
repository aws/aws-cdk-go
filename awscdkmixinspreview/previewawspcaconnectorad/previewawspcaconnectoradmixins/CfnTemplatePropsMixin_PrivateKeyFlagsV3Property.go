package previewawspcaconnectoradmixins


// Private key flags for v3 templates specify the client compatibility, if the private key can be exported, if user input is required when using a private key, and if an alternate signature algorithm should be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateKeyFlagsV3Property := &PrivateKeyFlagsV3Property{
//   	ClientVersion: jsii.String("clientVersion"),
//   	ExportableKey: jsii.Boolean(false),
//   	RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   	StrongKeyProtectionRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv3.html
//
type CfnTemplatePropsMixin_PrivateKeyFlagsV3Property struct {
	// Defines the minimum client compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv3.html#cfn-pcaconnectorad-template-privatekeyflagsv3-clientversion
	//
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// Allows the private key to be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv3.html#cfn-pcaconnectorad-template-privatekeyflagsv3-exportablekey
	//
	ExportableKey interface{} `field:"optional" json:"exportableKey" yaml:"exportableKey"`
	// Reguires the PKCS #1 v2.1 signature format for certificates. You should verify that your CA, objects, and applications can accept this signature format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv3.html#cfn-pcaconnectorad-template-privatekeyflagsv3-requirealternatesignaturealgorithm
	//
	RequireAlternateSignatureAlgorithm interface{} `field:"optional" json:"requireAlternateSignatureAlgorithm" yaml:"requireAlternateSignatureAlgorithm"`
	// Requirer user input when using the private key for enrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv3.html#cfn-pcaconnectorad-template-privatekeyflagsv3-strongkeyprotectionrequired
	//
	StrongKeyProtectionRequired interface{} `field:"optional" json:"strongKeyProtectionRequired" yaml:"strongKeyProtectionRequired"`
}

