//go:build !no_runtime_type_checking

package awssns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSubscriptionFilter_NumericFilterParameters(numericConditions *NumericConditions) error {
	if numericConditions == nil {
		return fmt.Errorf("parameter numericConditions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(numericConditions, func() string { return "parameter numericConditions" }); err != nil {
		return err
	}

	return nil
}

func validateSubscriptionFilter_StringFilterParameters(stringConditions *StringConditions) error {
	if stringConditions == nil {
		return fmt.Errorf("parameter stringConditions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(stringConditions, func() string { return "parameter stringConditions" }); err != nil {
		return err
	}

	return nil
}

