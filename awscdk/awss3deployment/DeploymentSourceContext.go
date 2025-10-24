package awss3deployment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Bind context for ISources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   deploymentSourceContext := &DeploymentSourceContext{
//   	HandlerRole: role,
//   }
//
type DeploymentSourceContext struct {
	// The role for the handler.
	HandlerRole awsiam.IRole `field:"required" json:"handlerRole" yaml:"handlerRole"`
}

