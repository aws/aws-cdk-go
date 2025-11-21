//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"
)

func validateComponentStepIfCondition_FromObjectParameters(ifObject *map[string]interface{}) error {
	if ifObject == nil {
		return fmt.Errorf("parameter ifObject is required, but nil was provided")
	}

	return nil
}

func validateNewComponentStepIfConditionParameters(ifCondition interface{}) error {
	if ifCondition == nil {
		return fmt.Errorf("parameter ifCondition is required, but nil was provided")
	}

	return nil
}

