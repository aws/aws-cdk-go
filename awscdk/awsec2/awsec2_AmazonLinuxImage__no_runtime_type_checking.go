//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmazonLinuxImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateAmazonLinuxImage_SsmParameterNameParameters(props *AmazonLinuxImageProps) error {
	return nil
}

func validateNewAmazonLinuxImageParameters(props *AmazonLinuxImageProps) error {
	return nil
}

