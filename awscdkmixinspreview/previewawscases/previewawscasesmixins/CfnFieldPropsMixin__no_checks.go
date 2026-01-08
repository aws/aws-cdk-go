//go:build no_runtime_type_checking

package previewawscasesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFieldPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFieldPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFieldPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFieldPropsMixinParameters(props *CfnFieldMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

