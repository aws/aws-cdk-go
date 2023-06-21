package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The type returned from `IProject#enableBatchBuilds`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   batchBuildConfig := &BatchBuildConfig{
//   	Role: role,
//   }
//
type BatchBuildConfig struct {
	// The IAM batch service Role of this Project.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

