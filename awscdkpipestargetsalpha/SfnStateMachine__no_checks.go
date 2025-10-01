//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SfnStateMachine) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_SfnStateMachine) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewSfnStateMachineParameters(stateMachine awsstepfunctions.IStateMachine, parameters *SfnStateMachineParameters) error {
	return nil
}

