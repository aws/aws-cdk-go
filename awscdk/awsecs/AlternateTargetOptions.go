package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for AlternateTarget configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerRuleConfiguration listenerRuleConfiguration
//   var role role
//
//   alternateTargetOptions := &AlternateTargetOptions{
//   	Role: role,
//   	TestListener: listenerRuleConfiguration,
//   }
//
type AlternateTargetOptions struct {
	// The IAM role for the configuration.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The test listener configuration.
	// Default: - none.
	//
	TestListener ListenerRuleConfiguration `field:"optional" json:"testListener" yaml:"testListener"`
}

