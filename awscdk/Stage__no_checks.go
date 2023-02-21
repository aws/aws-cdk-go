//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stage) validateSynthParameters(options *StageSynthesisOptions) error {
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

