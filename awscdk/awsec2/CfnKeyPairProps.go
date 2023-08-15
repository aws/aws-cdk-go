package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnKeyPair`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyPairProps := &CfnKeyPairProps{
//   	KeyName: jsii.String("keyName"),
//
//   	// the properties below are optional
//   	KeyFormat: jsii.String("keyFormat"),
//   	KeyType: jsii.String("keyType"),
//   	PublicKeyMaterial: jsii.String("publicKeyMaterial"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html
//
type CfnKeyPairProps struct {
	// A unique name for the key pair.
	//
	// Constraints: Up to 255 ASCII characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html#cfn-ec2-keypair-keyname
	//
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The format of the key pair.
	//
	// Default: `pem`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html#cfn-ec2-keypair-keyformat
	//
	// Default: - "pem".
	//
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// The type of key pair. Note that ED25519 keys are not supported for Windows instances.
	//
	// If the `PublicKeyMaterial` property is specified, the `KeyType` property is ignored, and the key type is inferred from the `PublicKeyMaterial` value.
	//
	// Default: `rsa`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html#cfn-ec2-keypair-keytype
	//
	// Default: - "rsa".
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The public key material.
	//
	// The `PublicKeyMaterial` property is used to import a key pair. If this property is not specified, then a new key pair will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html#cfn-ec2-keypair-publickeymaterial
	//
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// The tags to apply to the key pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html#cfn-ec2-keypair-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

