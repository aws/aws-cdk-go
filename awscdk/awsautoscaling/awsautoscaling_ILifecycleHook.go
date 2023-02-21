package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// A basic lifecycle hook object.
// Experimental.
type ILifecycleHook interface {
	awscdk.IResource
	// The role for the lifecycle hook to execute.
	// Experimental.
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

