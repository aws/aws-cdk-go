//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (c *jsiiProxy_CatalogBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CatalogBase) validateConfigureEncryptionParameters(options *CatalogEncryptionOptions) error {
	return nil
}

func (c *jsiiProxy_CatalogBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CatalogBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCatalogBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCatalogBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCatalogBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCatalogBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

