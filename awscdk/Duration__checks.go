//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_Duration) validateMinusParameters(rhs Duration) error {
	if rhs == nil {
		return fmt.Errorf("parameter rhs is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Duration) validatePlusParameters(rhs Duration) error {
	if rhs == nil {
		return fmt.Errorf("parameter rhs is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Duration) validateToDaysParameters(opts *TimeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_Duration) validateToHoursParameters(opts *TimeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_Duration) validateToMillisecondsParameters(opts *TimeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_Duration) validateToMinutesParameters(opts *TimeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_Duration) validateToSecondsParameters(opts *TimeConversionOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateDuration_DaysParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateDuration_HoursParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateDuration_MillisParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateDuration_MinutesParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateDuration_ParseParameters(duration *string) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func validateDuration_SecondsParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

