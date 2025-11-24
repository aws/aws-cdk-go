//go:build no_runtime_type_checking

package mixinsawslex

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBotAliasPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBotAliasPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBotAliasPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBotAliasPropsMixinParameters(props *CfnBotAliasMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

