//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAIGuardrailPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAIGuardrailPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAIGuardrailPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAIGuardrailPropsMixinParameters(props *CfnAIGuardrailMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

