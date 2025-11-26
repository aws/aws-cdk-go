//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDRTAccessPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDRTAccessPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDRTAccessPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDRTAccessPropsMixinParameters(props *CfnDRTAccessMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

