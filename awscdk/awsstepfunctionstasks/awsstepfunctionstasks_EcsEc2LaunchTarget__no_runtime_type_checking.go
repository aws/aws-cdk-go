//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsEc2LaunchTarget) validateBindParameters(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) error {
	return nil
}

func validateNewEcsEc2LaunchTargetParameters(options *EcsEc2LaunchTargetOptions) error {
	return nil
}

