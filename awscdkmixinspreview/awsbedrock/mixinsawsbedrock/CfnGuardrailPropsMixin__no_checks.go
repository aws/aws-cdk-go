//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGuardrailPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGuardrailPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGuardrailPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGuardrailPropsMixinParameters(props *CfnGuardrailMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

