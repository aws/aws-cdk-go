//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"
)

func validateWorkflowParameterValue_FromBooleanParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowParameterValue_FromIntegerParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowParameterValue_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowParameterValue_FromStringListParameters(values *[]*string) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateNewWorkflowParameterValueParameters(value *[]*string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

