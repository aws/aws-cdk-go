//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func validateStepFunctionsIntegration_StartExecutionParameters(stateMachine awsstepfunctions.IStateMachine, options *StepFunctionsExecutionIntegrationOptions) error {
	return nil
}

