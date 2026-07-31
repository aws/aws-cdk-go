package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// SRT encryption configuration for router inputs and outputs (SRT Listener and SRT Caller).
//
// Uses AWS Secrets Manager for key management. Distinct from
// {@link SrtPasswordEncryption}, which is used for flow sources and outputs.
//
// The secret must live in the same AWS account and Region as the router I/O that uses
// it. MediaConnect does not support cross-account or cross-Region secrets.
//
// **Trust-policy scope on routers.** Router I/O ids are service-generated (unknown at synth
// time), and pinning the live ARN would create a CloudFormation dependency cycle — so the
// auto-created role pins `aws:SourceArn` to a wildcarded ARN (`arn:...:routerInput:*` /
// `arn:...:routerOutput:*`) plus `aws:SourceAccount`. To pin a tighter trust policy, supply
// your own `role`.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//   var role IRole
//   var secret ISecret
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("EncryptedOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("encrypted-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtCaller(&SrtCallerOutputProtocolProps{
//   			DestinationAddress: jsii.String("203.0.113.100"),
//   			DestinationPort: jsii.Number(9001),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   			EncryptionConfiguration: &RouterSrtEncryption{
//   				Role: *Role,
//   				Secret: *Secret,
//   			},
//   		}),
//   		NetworkInterface: networkInterface,
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/cross-service-confused-deputy-prevention.html
//
// Experimental.
type RouterSrtEncryption struct {
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

