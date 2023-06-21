//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Product) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Product) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	return nil
}

func (p *jsiiProxy_Product) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Product) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateProduct_FromProductArnParameters(scope constructs.Construct, id *string, productArn *string) error {
	return nil
}

func validateProduct_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProduct_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateProduct_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewProductParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

