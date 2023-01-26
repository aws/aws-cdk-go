//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_Size) validateToGibibytesParameters(opts *SizeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Size) validateToKibibytesParameters(opts *SizeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Size) validateToMebibytesParameters(opts *SizeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Size) validateToPebibytesParameters(opts *SizeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Size) validateToTebibytesParameters(opts *SizeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateSize_GibibytesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateSize_KibibytesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateSize_MebibytesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateSize_PebibytesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateSize_TebibytesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

