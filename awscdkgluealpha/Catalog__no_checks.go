//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Catalog) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (c *jsiiProxy_Catalog) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Catalog) validateConfigureEncryptionParameters(options *CatalogEncryptionOptions) error {
	return nil
}

func (c *jsiiProxy_Catalog) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Catalog) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCatalog_EncryptAccountParameters(scope constructs.Construct, options *CatalogEncryptionOptions) error {
	return nil
}

func validateCatalog_ForAccountParameters(scope constructs.Construct) error {
	return nil
}

func validateCatalog_FromCatalogArnParameters(scope constructs.Construct, id *string, catalogArn *string) error {
	return nil
}

func validateCatalog_FromCatalogIdParameters(scope constructs.Construct, id *string, catalogId *string) error {
	return nil
}

func validateCatalog_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCatalog_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCatalog_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCatalogParameters(scope constructs.Construct, id *string, props *CatalogProps) error {
	return nil
}

