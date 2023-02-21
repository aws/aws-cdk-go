//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StateMachineFragment) validateNextParameters(next IChainable) error {
	return nil
}

func (s *jsiiProxy_StateMachineFragment) validateToSingleStateParameters(options *SingleStateOptions) error {
	return nil
}

func validateStateMachineFragment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStateMachineFragmentParameters(scope constructs.Construct, id *string) error {
	return nil
}

