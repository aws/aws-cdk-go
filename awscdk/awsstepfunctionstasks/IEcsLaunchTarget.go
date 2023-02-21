package awsstepfunctionstasks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An Amazon ECS launch type determines the type of infrastructure on which your tasks and services are hosted.
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html
//
type IEcsLaunchTarget interface {
	// called when the ECS launch target is configured on RunTask.
	Bind(task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy for IEcsLaunchTarget
type jsiiProxy_IEcsLaunchTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IEcsLaunchTarget) Bind(task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	if err := i.validateBindParameters(task, launchTargetOptions); err != nil {
		panic(err)
	}
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{task, launchTargetOptions},
		&returns,
	)

	return returns
}

