//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroupTargetInput) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateLogGroupTargetInput_FromEventPathParameters(path *string) error {
	return nil
}

func validateLogGroupTargetInput_FromMultilineTextParameters(text *string) error {
	return nil
}

func validateLogGroupTargetInput_FromObjectParameters(obj interface{}) error {
	return nil
}

func validateLogGroupTargetInput_FromObjectV2Parameters(options *LogGroupTargetInputOptions) error {
	return nil
}

func validateLogGroupTargetInput_FromTextParameters(text *string) error {
	return nil
}

