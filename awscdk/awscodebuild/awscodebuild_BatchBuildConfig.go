package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The type returned from {@link IProject#enableBatchBuilds}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   batchBuildConfig := &batchBuildConfig{
//   	role: role,
//   }
//
// Experimental.
type BatchBuildConfig struct {
	// The IAM batch service Role of this Project.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

