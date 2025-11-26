//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEIPPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEIPPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEIPPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEIPPropsMixinParameters(props *CfnEIPMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

