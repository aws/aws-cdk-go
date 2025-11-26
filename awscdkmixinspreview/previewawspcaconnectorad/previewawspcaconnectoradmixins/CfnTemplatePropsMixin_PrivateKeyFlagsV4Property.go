package previewawspcaconnectoradmixins


// Private key flags for v4 templates specify the client compatibility, if the private key can be exported, if user input is required when using a private key, if an alternate signature algorithm should be used, and if certificates are renewed using the same private key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateKeyFlagsV4Property := &PrivateKeyFlagsV4Property{
//   	ClientVersion: jsii.String("clientVersion"),
//   	ExportableKey: jsii.Boolean(false),
//   	RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   	RequireSameKeyRenewal: jsii.Boolean(false),
//   	StrongKeyProtectionRequired: jsii.Boolean(false),
//   	UseLegacyProvider: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html
//
type CfnTemplatePropsMixin_PrivateKeyFlagsV4Property struct {
	// Defines the minimum client compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-clientversion
	//
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// Allows the private key to be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-exportablekey
	//
	ExportableKey interface{} `field:"optional" json:"exportableKey" yaml:"exportableKey"`
	// Requires the PKCS #1 v2.1 signature format for certificates. You should verify that your CA, objects, and applications can accept this signature format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-requirealternatesignaturealgorithm
	//
	RequireAlternateSignatureAlgorithm interface{} `field:"optional" json:"requireAlternateSignatureAlgorithm" yaml:"requireAlternateSignatureAlgorithm"`
	// Renew certificate using the same private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-requiresamekeyrenewal
	//
	RequireSameKeyRenewal interface{} `field:"optional" json:"requireSameKeyRenewal" yaml:"requireSameKeyRenewal"`
	// Require user input when using the private key for enrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-strongkeyprotectionrequired
	//
	StrongKeyProtectionRequired interface{} `field:"optional" json:"strongKeyProtectionRequired" yaml:"strongKeyProtectionRequired"`
	// Specifies the cryptographic service provider category used to generate private keys.
	//
	// Set to TRUE to use Legacy Cryptographic Service Providers and FALSE to use Key Storage Providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-uselegacyprovider
	//
	UseLegacyProvider interface{} `field:"optional" json:"useLegacyProvider" yaml:"useLegacyProvider"`
}

