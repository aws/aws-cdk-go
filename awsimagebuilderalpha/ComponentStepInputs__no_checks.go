//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateComponentStepInputs_FromListParameters(inputsObjectList *[]*map[string]interface{}) error {
	return nil
}

func validateComponentStepInputs_FromObjectParameters(inputsObject *map[string]interface{}) error {
	return nil
}

func validateNewComponentStepInputsParameters(input interface{}) error {
	return nil
}

