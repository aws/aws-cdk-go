//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInAppTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInAppTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInAppTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInAppTemplatePropsMixinParameters(props *CfnInAppTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

