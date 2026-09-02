package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a MediaConnect router input.
//
// Example:
//   // Redundant pipelines with explicit AZs
//   awsmedialivealpha.InputConfiguration_MediaConnectRouter(&MediaConnectRouterInputProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//   })
//
// Experimental.
type MediaConnectRouterInputProps struct {
	// The availability zones for the router input pipelines.
	//
	// Provide one AZ for a single pipeline, or two for redundant pipelines.
	// If omitted, defaults to the stack's first availability zone (single pipeline).
	// Default: - single pipeline using the stack's first availability zone.
	//
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The Secrets Manager secret for custom encryption.
	//
	// If not provided, automatic encryption is used.
	// Default: - automatic encryption.
	//
	// Experimental.
	EncryptionSecret awssecretsmanager.ISecret `field:"optional" json:"encryptionSecret" yaml:"encryptionSecret"`
}

