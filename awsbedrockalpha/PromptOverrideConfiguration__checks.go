//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validatePromptOverrideConfiguration_FromStepsParameters(steps *[]*PromptStepConfigBase) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}
	for idx_b7595e, v := range *steps {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter steps[%#v]", idx_b7595e) }); err != nil {
			return err
		}
	}

	return nil
}

func validatePromptOverrideConfiguration_WithCustomParserParameters(props *CustomParserProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

