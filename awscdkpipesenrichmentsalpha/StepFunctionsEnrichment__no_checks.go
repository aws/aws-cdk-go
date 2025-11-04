//go:build no_runtime_type_checking

package awscdkpipesenrichmentsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionsEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsEnrichment) validateGrantInvokeParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewStepFunctionsEnrichmentParameters(stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsEnrichmentProps) error {
	return nil
}

