//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowParameterValue_FromBooleanParameters(value *bool) error {
	return nil
}

func validateWorkflowParameterValue_FromIntegerParameters(value *float64) error {
	return nil
}

func validateWorkflowParameterValue_FromStringParameters(value *string) error {
	return nil
}

func validateWorkflowParameterValue_FromStringListParameters(values *[]*string) error {
	return nil
}

func validateNewWorkflowParameterValueParameters(value *[]*string) error {
	return nil
}

