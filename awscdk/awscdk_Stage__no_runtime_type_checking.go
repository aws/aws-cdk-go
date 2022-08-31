//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stage) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Stage) validateSynthParameters(options *StageSynthesisOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStage_IsStageParameters(x interface{}) error {
	return nil
}

func validateStage_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStageParameters(scope constructs.Construct, id *string, props *StageProps) error {
	return nil
}

