package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration for a deployment lifecycle hook target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   deploymentLifecycleHookTargetConfig := &DeploymentLifecycleHookTargetConfig{
//   	LifecycleStages: []DeploymentLifecycleStage{
//   		awscdk.Aws_ecs.DeploymentLifecycleStage_RECONCILE_SERVICE,
//   	},
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Role: role,
//   }
//
type DeploymentLifecycleHookTargetConfig struct {
	// The lifecycle stages when this hook should be executed.
	LifecycleStages *[]DeploymentLifecycleStage `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// The ARN of the target resource.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The IAM role that grants permissions to invoke the target.
	// Default: - a role will be created automatically.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

