//go:build no_runtime_type_checking

package previewawswafmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnByteMatchSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnByteMatchSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnByteMatchSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnByteMatchSetPropsMixinParameters(props *CfnByteMatchSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

