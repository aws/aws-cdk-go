//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StateMachineGrants) validateActionsParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateExecutionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateRedriveExecutionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateStartExecutionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateStartSyncExecutionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachineGrants) validateTaskResponseParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateStateMachineGrants_FromStateMachineParameters(resource interfacesawsstepfunctions.IStateMachineRef) error {
	return nil
}

