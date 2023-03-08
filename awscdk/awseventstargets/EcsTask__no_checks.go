//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsTask) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewEcsTaskParameters(props *EcsTaskProps) error {
	return nil
}

