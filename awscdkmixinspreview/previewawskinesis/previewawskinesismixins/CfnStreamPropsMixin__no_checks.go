//go:build no_runtime_type_checking

package previewawskinesismixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStreamPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStreamPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStreamPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStreamPropsMixinParameters(props *CfnStreamMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

