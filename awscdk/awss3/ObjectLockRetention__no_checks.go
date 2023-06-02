//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func validateObjectLockRetention_ComplianceParameters(duration awscdk.Duration) error {
	return nil
}

func validateObjectLockRetention_GovernanceParameters(duration awscdk.Duration) error {
	return nil
}

