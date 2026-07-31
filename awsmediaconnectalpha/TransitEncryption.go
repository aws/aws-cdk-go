package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Transit encryption configuration for router integrations — securing the link between a router and a flow or a MediaLive channel/input.
//
// Uses AWS Secrets Manager for key
// management.
//
// The secret must live in the same AWS account and Region as the consuming resource.
// MediaConnect does not support cross-account or cross-Region secrets.
//
// **Trust-policy scope on routers.** Router I/O ids are service-generated (unknown at synth
// time), and pinning the live ARN would create a CloudFormation dependency cycle — so the
// auto-created role pins `aws:SourceArn` to a wildcarded ARN (`arn:...:routerInput:*` /
// `arn:...:routerOutput:*`) plus `aws:SourceAccount`. To pin a tighter trust policy, supply
// your own `role`.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//   var existingRouterOutput RouterOutput
//
//
//   // Flow output to router with transit encryption
//   routerOutput := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("RouterOutput"), &FlowOutputProps{
//   	Flow: flow,
//   	Output: awsmediaconnectalpha.OutputConfiguration_Router(&RouterTransitConfig{
//   		Encryption: &TransitEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
//   // Flow source from router with transit encryption
//   flowFromRouter := awsmediaconnectalpha.NewFlow(stack, jsii.String("FlowFromRouter"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Router(&RouterSource{
//   		RouterOutput: existingRouterOutput,
//   		Decryption: &TransitEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/cross-service-confused-deputy-prevention.html
//
// Experimental.
type TransitEncryption struct {
	// Secrets Manager secret containing the transit encryption key.
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

