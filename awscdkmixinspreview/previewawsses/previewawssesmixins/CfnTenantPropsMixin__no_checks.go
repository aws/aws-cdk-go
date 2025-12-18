//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTenantPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTenantPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTenantPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTenantPropsMixinParameters(props *CfnTenantMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

