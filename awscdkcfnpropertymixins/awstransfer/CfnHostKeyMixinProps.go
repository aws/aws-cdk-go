package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHostKeyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHostKeyMixinProps := &CfnHostKeyMixinProps{
//   	Description: jsii.String("description"),
//   	HostKeyBody: jsii.String("hostKeyBody"),
//   	KmsKey: jsii.String("kmsKey"),
//   	ServerId: jsii.String("serverId"),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html
//
type CfnHostKeyMixinProps struct {
	// The text description for this host key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The private key portion of an SSH key pair.
	//
	// Transfer Family accepts RSA, ECDSA, and ED25519 keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-hostkeybody
	//
	HostKeyBody *string `field:"optional" json:"hostKeyBody" yaml:"hostKeyBody"`
	// The Amazon Resource Name (ARN) of a customer managed key used to encrypt the Secrets Manager secret that holds the host key.
	//
	// Applies only when StorageType is SECRETS_MANAGER. If omitted, the AWS managed key is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The identifier of the server that contains the host key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-serverid
	//
	ServerId *string `field:"optional" json:"serverId" yaml:"serverId"`
	// The type of storage used for the host key.
	//
	// SERVICE_MANAGED keys are stored within the Transfer Family service. SECRETS_MANAGER keys are stored in an AWS Secrets Manager secret in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Key-value pairs that can be used to group and search for host keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-hostkey.html#cfn-transfer-hostkey-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

