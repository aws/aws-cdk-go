//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationTemplate) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateCloudFormationTemplate_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateCloudFormationTemplate_FromProductStackParameters(productStack ProductStack) error {
	return nil
}

func validateCloudFormationTemplate_FromUrlParameters(url *string) error {
	return nil
}

