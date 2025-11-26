package previewawspaymentcryptographymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnKeyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKeyMixinProps := &CfnKeyMixinProps{
//   	DeriveKeyUsage: jsii.String("deriveKeyUsage"),
//   	Enabled: jsii.Boolean(false),
//   	Exportable: jsii.Boolean(false),
//   	KeyAttributes: &KeyAttributesProperty{
//   		KeyAlgorithm: jsii.String("keyAlgorithm"),
//   		KeyClass: jsii.String("keyClass"),
//   		KeyModesOfUse: &KeyModesOfUseProperty{
//   			Decrypt: jsii.Boolean(false),
//   			DeriveKey: jsii.Boolean(false),
//   			Encrypt: jsii.Boolean(false),
//   			Generate: jsii.Boolean(false),
//   			NoRestrictions: jsii.Boolean(false),
//   			Sign: jsii.Boolean(false),
//   			Unwrap: jsii.Boolean(false),
//   			Verify: jsii.Boolean(false),
//   			Wrap: jsii.Boolean(false),
//   		},
//   		KeyUsage: jsii.String("keyUsage"),
//   	},
//   	KeyCheckValueAlgorithm: jsii.String("keyCheckValueAlgorithm"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html
//
type CfnKeyMixinProps struct {
	// The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-derivekeyusage
	//
	DeriveKeyUsage *string `field:"optional" json:"deriveKeyUsage" yaml:"deriveKeyUsage"`
	// Specifies whether the key is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key is exportable.
	//
	// This data is immutable after the key is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-exportable
	//
	Exportable interface{} `field:"optional" json:"exportable" yaml:"exportable"`
	// The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key.
	//
	// This data is immutable after the key is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-keyattributes
	//
	KeyAttributes interface{} `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV).
	//
	// It is used to validate the key integrity.
	//
	// For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-keycheckvaluealgorithm
	//
	KeyCheckValueAlgorithm *string `field:"optional" json:"keyCheckValueAlgorithm" yaml:"keyCheckValueAlgorithm"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

