package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
)

// Properties for a MediaConnect input.
//
// Example:
//   var role IRole
//   var flow IFlowRef
//
//
//   awsmedialivealpha.InputConfiguration_MediaConnect(&MediaConnectInputProps{
//   	Flows: []IFlowRef{
//   		flow,
//   	},
//   	Role: Role,
//   })
//
// Experimental.
type MediaConnectInputProps struct {
	// The MediaConnect flows to use as sources (one, or two for pipeline redundancy).
	// Experimental.
	Flows *[]interfacesawsmediaconnect.IFlowRef `field:"required" json:"flows" yaml:"flows"`
	// The IAM role MediaLive uses to manage the output it adds to the flow for this input.
	//
	// When omitted, a role is created and granted the required permissions.
	// When you provide a role, no permissions are added — you own all the permissions it needs.
	// Default: - a role is created with the medialive.amazonaws.com service principal
	//
	// [disable-awslint:prefer-ref-interface].
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

