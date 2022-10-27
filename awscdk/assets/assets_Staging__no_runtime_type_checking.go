//go:build no_runtime_type_checking

package assets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Staging) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Staging) validateRelativeStagedPathParameters(stack awscdk.Stack) error {
	return nil
}

func (s *jsiiProxy_Staging) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStaging_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStagingParameters(scope awscdk.Construct, id *string, props *StagingProps) error {
	return nil
}

