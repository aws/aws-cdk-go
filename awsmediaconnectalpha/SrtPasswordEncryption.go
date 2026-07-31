package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// SRT password encryption/decryption configuration for SRT Listener and SRT Caller sources and outputs on flows.
//
// The secret must live in the same AWS account and Region as the resource (source or
// output) that uses it. MediaConnect does not support cross-account or cross-Region
// secrets.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   // SRT Caller output with encryption
//   output := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("EncryptedOutput"), &FlowOutputProps{
//   	Flow: flow,
//   	Description: jsii.String("Encrypted SRT output"),
//   	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
//   		Destination: jsii.String("203.0.113.100"),
//   		Port: jsii.Number(7000),
//   		Encryption: &SrtPasswordEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/cross-service-confused-deputy-prevention.html
//
// Experimental.
type SrtPasswordEncryption struct {
	// Secrets Manager secret containing the SRT passphrase.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// IAM role that MediaConnect assumes to access the Secrets Manager secret.
	//
	// If provided, the role is used as-is; you must grant it the necessary permissions
	// yourself.
	// Default: - a scoped role is auto-created with read access to the secret and a
	// confused-deputy trust condition. See the **Encryption** section of the module README
	// for the generated trust policy.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

