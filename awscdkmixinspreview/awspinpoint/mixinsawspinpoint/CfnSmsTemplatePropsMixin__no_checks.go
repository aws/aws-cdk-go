//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSmsTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSmsTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSmsTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSmsTemplatePropsMixinParameters(props *CfnSmsTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

