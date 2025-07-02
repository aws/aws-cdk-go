//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ValidationResults) validateCollectParameters(result ValidationResult) error {
	return nil
}

func (v *jsiiProxy_ValidationResults) validateWrapParameters(message *string) error {
	return nil
}

func (j *jsiiProxy_ValidationResults) validateSetResultsParameters(val *[]ValidationResult) error {
	return nil
}

