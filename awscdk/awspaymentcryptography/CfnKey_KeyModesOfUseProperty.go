package awspaymentcryptography


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnKey_KeyModesOfUseProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-decrypt
	//
	// Default: - false.
	//
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-derivekey
	//
	// Default: - false.
	//
	DeriveKey interface{} `field:"optional" json:"deriveKey" yaml:"deriveKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-encrypt
	//
	// Default: - false.
	//
	Encrypt interface{} `field:"optional" json:"encrypt" yaml:"encrypt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-generate
	//
	// Default: - false.
	//
	Generate interface{} `field:"optional" json:"generate" yaml:"generate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-norestrictions
	//
	// Default: - false.
	//
	NoRestrictions interface{} `field:"optional" json:"noRestrictions" yaml:"noRestrictions"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-verify
	//
	// Default: - false.
	//
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keymodesofuse.html#cfn-paymentcryptography-key-keymodesofuse-wrap
	//
	// Default: - false.
	//
	Wrap interface{} `field:"optional" json:"wrap" yaml:"wrap"`
}

