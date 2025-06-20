//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IEcsLaunchTarget) validateBindParameters(task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	if launchTargetOptions == nil {
		return fmt.Errorf("parameter launchTargetOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(launchTargetOptions, func() string { return "parameter launchTargetOptions" }); err != nil {
		return err
	}

	return nil
}

