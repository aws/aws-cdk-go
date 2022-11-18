//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationProduct) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudFormationProduct) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationProduct) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudFormationProduct) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCloudFormationProduct_FromProductArnParameters(scope constructs.Construct, id *string, productArn *string) error {
	return nil
}

func validateCloudFormationProduct_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudFormationProduct_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudFormationProduct_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudFormationProductParameters(scope constructs.Construct, id *string, props *CloudFormationProductProps) error {
	return nil
}

