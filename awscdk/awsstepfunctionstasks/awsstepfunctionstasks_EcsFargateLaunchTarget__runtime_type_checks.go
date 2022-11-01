//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (e *jsiiProxy_EcsFargateLaunchTarget) validateBindParameters(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) error {
	if _task == nil {
		return fmt.Errorf("parameter _task is required, but nil was provided")
	}

	if launchTargetOptions == nil {
		return fmt.Errorf("parameter launchTargetOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(launchTargetOptions, func() string { return "parameter launchTargetOptions" }); err != nil {
		return err
	}

	return nil
}

func validateNewEcsFargateLaunchTargetParameters(options *EcsFargateLaunchTargetOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

