//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateComponentStepIfCondition_FromObjectParameters(ifObject *map[string]interface{}) error {
	return nil
}

func validateNewComponentStepIfConditionParameters(ifCondition interface{}) error {
	return nil
}

