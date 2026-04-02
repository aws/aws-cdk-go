//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateStartTag_OfParameters(timeOffset *float64, options *StartTagOptions) error {
	if timeOffset == nil {
		return fmt.Errorf("parameter timeOffset is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateStartTag_WithPreciseParameters(timeOffset *float64) error {
	if timeOffset == nil {
		return fmt.Errorf("parameter timeOffset is required, but nil was provided")
	}

	return nil
}

