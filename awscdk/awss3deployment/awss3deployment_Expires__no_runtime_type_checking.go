//go:build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func validateExpires_AfterParameters(t awscdk.Duration) error {
	return nil
}

func validateExpires_AtDateParameters(d *time.Time) error {
	return nil
}

func validateExpires_AtTimestampParameters(t *float64) error {
	return nil
}

func validateExpires_FromStringParameters(s *string) error {
	return nil
}

