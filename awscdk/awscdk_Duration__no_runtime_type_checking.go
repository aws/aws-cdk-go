//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Duration) validateMinusParameters(rhs Duration) error {
	return nil
}

func (d *jsiiProxy_Duration) validatePlusParameters(rhs Duration) error {
	return nil
}

func (d *jsiiProxy_Duration) validateToDaysParameters(opts *TimeConversionOptions) error {
	return nil
}

func (d *jsiiProxy_Duration) validateToHoursParameters(opts *TimeConversionOptions) error {
	return nil
}

func (d *jsiiProxy_Duration) validateToMillisecondsParameters(opts *TimeConversionOptions) error {
	return nil
}

func (d *jsiiProxy_Duration) validateToMinutesParameters(opts *TimeConversionOptions) error {
	return nil
}

func (d *jsiiProxy_Duration) validateToSecondsParameters(opts *TimeConversionOptions) error {
	return nil
}

func validateDuration_DaysParameters(amount *float64) error {
	return nil
}

func validateDuration_HoursParameters(amount *float64) error {
	return nil
}

func validateDuration_MillisParameters(amount *float64) error {
	return nil
}

func validateDuration_MinutesParameters(amount *float64) error {
	return nil
}

func validateDuration_ParseParameters(duration *string) error {
	return nil
}

func validateDuration_SecondsParameters(amount *float64) error {
	return nil
}

