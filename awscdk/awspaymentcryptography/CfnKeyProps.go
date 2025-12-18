package awspaymentcryptography

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyProps := &CfnKeyProps{
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
//
//   	// the properties below are optional
//   	DeriveKeyUsage: jsii.String("deriveKeyUsage"),
//   	Enabled: jsii.Boolean(false),
//   	KeyCheckValueAlgorithm: jsii.String("keyCheckValueAlgorithm"),
//   	ReplicationRegions: []*string{
//   		jsii.String("replicationRegions"),
//   	},
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
type CfnKeyProps struct {
	// Specifies whether the key is exportable.
	//
	// This data is immutable after the key is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-exportable
	//
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key.
	//
	// This data is immutable after the key is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-keyattributes
	//
	KeyAttributes interface{} `field:"required" json:"keyAttributes" yaml:"keyAttributes"`
	// The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-derivekeyusage
	//
	DeriveKeyUsage *string `field:"optional" json:"deriveKeyUsage" yaml:"deriveKeyUsage"`
	// Specifies whether the key is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV).
	//
	// It is used to validate the key integrity.
	//
	// For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-keycheckvaluealgorithm
	//
	KeyCheckValueAlgorithm *string `field:"optional" json:"keyCheckValueAlgorithm" yaml:"keyCheckValueAlgorithm"`
	// The list of AWS Regions to remove from the key's replication configuration.
	//
	// The key will no longer be available for cryptographic operations in these regions after removal. Ensure no active operations depend on the key in these regions before removal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-replicationregions
	//
	ReplicationRegions *[]*string `field:"optional" json:"replicationRegions" yaml:"replicationRegions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-key.html#cfn-paymentcryptography-key-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

