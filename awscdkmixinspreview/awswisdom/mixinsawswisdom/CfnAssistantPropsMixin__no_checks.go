//go:build no_runtime_type_checking

package mixinsawswisdom

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssistantPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAssistantPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAssistantPropsMixinParameters(props *CfnAssistantMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

