package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Result of binding a lifecycle hook to a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   lifecycleHookTargetConfig := &LifecycleHookTargetConfig{
//   	CreatedRole: role,
//   	NotificationTargetArn: jsii.String("notificationTargetArn"),
//   }
//
type LifecycleHookTargetConfig struct {
	// The IRole that was used to bind the lifecycle hook to the target.
	CreatedRole awsiam.IRole `field:"required" json:"createdRole" yaml:"createdRole"`
	// The targetArn that the lifecycle hook was bound to.
	NotificationTargetArn *string `field:"required" json:"notificationTargetArn" yaml:"notificationTargetArn"`
}

