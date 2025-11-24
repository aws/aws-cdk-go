//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGuardrailVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGuardrailVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGuardrailVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGuardrailVersionPropsMixinParameters(props *CfnGuardrailVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

