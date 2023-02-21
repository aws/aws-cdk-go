//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDistribution) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDistribution) validateGrantCreateInvalidationParameters(identity awsiam.IGrantable) error {
	return nil
}

