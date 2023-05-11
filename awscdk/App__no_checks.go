//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateSynthParameters(options *StageSynthesisOptions) error {
	return nil
}

func validateApp_IsAppParameters(obj interface{}) error {
	return nil
}

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApp_IsStageParameters(x interface{}) error {
	return nil
}

func validateApp_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAppParameters(props *AppProps) error {
	return nil
}

