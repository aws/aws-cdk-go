//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DkimIdentity) validateBindParameters(emailIdentity EmailIdentity) error {
	return nil
}

func validateDkimIdentity_ByoDkimParameters(options *ByoDkimOptions) error {
	return nil
}

