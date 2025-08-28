//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_Signals) validateDoRenderParameters(options *SignalsOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Signals) validateRenderCreationPolicyParameters(renderOptions *RenderSignalsOptions) error {
	if renderOptions == nil {
		return fmt.Errorf("parameter renderOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(renderOptions, func() string { return "parameter renderOptions" }); err != nil {
		return err
	}

	return nil
}

func validateSignals_WaitForAllParameters(options *SignalsOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSignals_WaitForCountParameters(count *float64, options *SignalsOptions) error {
	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSignals_WaitForMinCapacityParameters(options *SignalsOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

