//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAIPromptPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAIPromptPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAIPromptPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAIPromptPropsMixinParameters(props *CfnAIPromptMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

