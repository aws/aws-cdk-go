//go:build no_runtime_type_checking

package mixinsawslex

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBotVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBotVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBotVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBotVersionPropsMixinParameters(props *CfnBotVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

