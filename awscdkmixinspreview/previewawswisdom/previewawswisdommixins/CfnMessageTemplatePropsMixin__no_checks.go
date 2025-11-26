//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMessageTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMessageTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMessageTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMessageTemplatePropsMixinParameters(props *CfnMessageTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

