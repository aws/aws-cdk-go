//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPushTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPushTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPushTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPushTemplatePropsMixinParameters(props *CfnPushTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

