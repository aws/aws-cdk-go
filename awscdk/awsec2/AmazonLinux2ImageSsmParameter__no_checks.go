//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmazonLinux2ImageSsmParameter) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateAmazonLinux2ImageSsmParameter_SsmParameterNameParameters(props *AmazonLinux2ImageSsmParameterProps) error {
	return nil
}

func validateNewAmazonLinux2ImageSsmParameterParameters(props *AmazonLinux2ImageSsmParameterProps) error {
	return nil
}

