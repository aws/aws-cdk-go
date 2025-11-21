//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"
)

func validateComponentStepInputs_FromListParameters(inputsObjectList *[]*map[string]interface{}) error {
	if inputsObjectList == nil {
		return fmt.Errorf("parameter inputsObjectList is required, but nil was provided")
	}

	return nil
}

func validateComponentStepInputs_FromObjectParameters(inputsObject *map[string]interface{}) error {
	if inputsObject == nil {
		return fmt.Errorf("parameter inputsObject is required, but nil was provided")
	}

	return nil
}

func validateNewComponentStepInputsParameters(input interface{}) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	return nil
}

