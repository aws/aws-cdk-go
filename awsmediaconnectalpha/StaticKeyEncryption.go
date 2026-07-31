package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Static key encryption/decryption configuration for Zixi protocol sources and outputs, and flow entitlements.
//
// The secret must live in the same AWS account and Region as the resource (source,
// output, or entitlement) that uses it. MediaConnect does not support cross-account or
// cross-Region secrets.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   entitlement := awsmediaconnectalpha.NewFlowEntitlement(stack, jsii.String("MyEntitlement"), &FlowEntitlementProps{
//   	Flow: flow,
//   	Description: jsii.String("Grant partner access to live feed"),
//   	Subscribers: []*string{
//   		jsii.String("111122223333"),
//   	},
//   	Encryption: &StaticKeyEncryption{
//   		Role: *Role,
//   		Secret: *Secret,
//   		Algorithm: awsmediaconnectalpha.EncryptionAlgorithm_AES256(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/cross-service-confused-deputy-prevention.html
//
// Experimental.
type StaticKeyEncryption struct {
	// The encryption algorithm to use.
	// Experimental.
	Algorithm EncryptionAlgorithm `field:"required" json:"algorithm" yaml:"algorithm"`
	// Secrets Manager secret containing the static encryption key.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// IAM role that MediaConnect assumes to access the Secrets Manager secret.
	//
	// If provided, the role is used as-is; you must grant it the necessary permissions
	// yourself.
	// Default: - a scoped role is auto-created with read access to the secret (including
	// `kms:Decrypt` for a customer-managed key) and a confused-deputy trust condition. See
	// the **Encryption** section of the module README for the generated trust policy.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

