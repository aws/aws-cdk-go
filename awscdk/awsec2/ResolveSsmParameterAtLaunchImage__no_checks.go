//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResolveSsmParameterAtLaunchImage) validateGetImageParameters(_arg constructs.Construct) error {
	return nil
}

func validateNewResolveSsmParameterAtLaunchImageParameters(parameterName *string, props *SsmParameterImageOptions) error {
	return nil
}

