//go:build no_runtime_type_checking

package mixinsawslex

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBotPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBotPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBotPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBotPropsMixinParameters(props *CfnBotMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

