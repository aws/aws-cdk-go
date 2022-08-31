//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OriginGroup) validateBindParameters(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewOriginGroupParameters(props *OriginGroupProps) error {
	return nil
}

