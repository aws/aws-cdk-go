package previewawspaymentcryptographymixins


// The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key.
//
// This data is immutable after the key is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyAttributesProperty := &KeyAttributesProperty{
//   	KeyAlgorithm: jsii.String("keyAlgorithm"),
//   	KeyClass: jsii.String("keyClass"),
//   	KeyModesOfUse: &KeyModesOfUseProperty{
//   		Decrypt: jsii.Boolean(false),
//   		DeriveKey: jsii.Boolean(false),
//   		Encrypt: jsii.Boolean(false),
//   		Generate: jsii.Boolean(false),
//   		NoRestrictions: jsii.Boolean(false),
//   		Sign: jsii.Boolean(false),
//   		Unwrap: jsii.Boolean(false),
//   		Verify: jsii.Boolean(false),
//   		Wrap: jsii.Boolean(false),
//   	},
//   	KeyUsage: jsii.String("keyUsage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html
//
type CfnKeyPropsMixin_KeyAttributesProperty struct {
	// The key algorithm to be use during creation of an AWS Payment Cryptography key.
	//
	// For symmetric keys, AWS Payment Cryptography supports `AES` and `TDES` algorithms. For asymmetric keys, AWS Payment Cryptography supports `RSA` and `ECC_NIST` algorithms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyalgorithm
	//
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// The type of AWS Payment Cryptography key to create, which determines the classiﬁcation of the cryptographic method and whether AWS Payment Cryptography key contains a symmetric key or an asymmetric key pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyclass
	//
	KeyClass *string `field:"optional" json:"keyClass" yaml:"keyClass"`
	// The list of cryptographic operations that you can perform using the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keymodesofuse
	//
	KeyModesOfUse interface{} `field:"optional" json:"keyModesOfUse" yaml:"keyModesOfUse"`
	// The cryptographic usage of an AWS Payment Cryptography key as deﬁned in section A.5.2 of the TR-31 spec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyusage
	//
	KeyUsage *string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
}

