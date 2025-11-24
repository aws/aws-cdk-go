package mixinsawspaymentcryptography


// The list of cryptographic operations that you can perform using the key.
//
// The modes of use are deﬁned in section A.5.3 of the TR-31 spec.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyModesOfUseProperty := &KeyModesOfUseProperty{
//   	Decrypt: jsii.Boolean(false),
//   	DeriveKey: jsii.Boolean(false),
//   	Encrypt: jsii.Boolean(false),
//   	Generate: jsii.Boolean(false),
//   	NoRestrictions: jsii.Boolean(false),
//   	Sign: jsii.Boolean(false),
//   	Unwrap: jsii.Boolean(false),
//   	Verify: jsii.Boolean(false),
//   	Wrap: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html
//
type CfnKeyPropsMixin_KeyModesOfUseProperty struct {
	// Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-decrypt
	//
	// Default: - false.
	//
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-derivekey
	//
	// Default: - false.
	//
	DeriveKey interface{} `field:"optional" json:"deriveKey" yaml:"deriveKey"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-encrypt
	//
	// Default: - false.
	//
	Encrypt interface{} `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-generate
	//
	// Default: - false.
	//
	Generate interface{} `field:"optional" json:"generate" yaml:"generate"`
	// Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by `KeyUsage` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-norestrictions
	//
	// Default: - false.
	//
	NoRestrictions interface{} `field:"optional" json:"noRestrictions" yaml:"noRestrictions"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used for signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-sign
	//
	// Default: - false.
	//
	Sign interface{} `field:"optional" json:"sign" yaml:"sign"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-unwrap
	//
	// Default: - false.
	//
	Unwrap interface{} `field:"optional" json:"unwrap" yaml:"unwrap"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-verify
	//
	// Default: - false.
	//
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-wrap
	//
	// Default: - false.
	//
	Wrap interface{} `field:"optional" json:"wrap" yaml:"wrap"`
}

