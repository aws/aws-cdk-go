//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCatalogPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCatalogPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCatalogPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCatalogPropsMixinParameters(props *CfnCatalogMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

