//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Expiration) validateIsAfterParameters(t Duration) error {
	return nil
}

func (e *jsiiProxy_Expiration) validateIsBeforeParameters(t Duration) error {
	return nil
}

func validateExpiration_AfterParameters(t Duration) error {
	return nil
}

func validateExpiration_AtDateParameters(d *time.Time) error {
	return nil
}

func validateExpiration_AtTimestampParameters(t *float64) error {
	return nil
}

func validateExpiration_FromStringParameters(s *string) error {
	return nil
}

