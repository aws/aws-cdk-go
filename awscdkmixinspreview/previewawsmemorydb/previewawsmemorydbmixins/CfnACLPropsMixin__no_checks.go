//go:build no_runtime_type_checking

package previewawsmemorydbmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnACLPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnACLPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnACLPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnACLPropsMixinParameters(props *CfnACLMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

