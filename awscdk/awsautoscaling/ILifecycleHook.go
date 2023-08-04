package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A basic lifecycle hook object.
type ILifecycleHook interface {
	awscdk.IResource
	// The role for the lifecycle hook to execute.
	// Default: - A default role is created if 'notificationTarget' is specified.
	// Otherwise, no role is created.
	//
	Role() awsiam.IRole
}

// The jsii proxy for ILifecycleHook
type jsiiProxy_ILifecycleHook struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILifecycleHook) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

