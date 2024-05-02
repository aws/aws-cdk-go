package awspaymentcryptography


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnKey_KeyAttributesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyalgorithm
	//
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyclass
	//
	KeyClass *string `field:"required" json:"keyClass" yaml:"keyClass"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keymodesofuse
	//
	KeyModesOfUse interface{} `field:"required" json:"keyModesOfUse" yaml:"keyModesOfUse"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-keyattributes.html#cfn-paymentcryptography-key-keyattributes-keyusage
	//
	KeyUsage *string `field:"required" json:"keyUsage" yaml:"keyUsage"`
}

