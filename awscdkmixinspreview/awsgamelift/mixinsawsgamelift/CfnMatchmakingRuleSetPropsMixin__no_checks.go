//go:build no_runtime_type_checking

package mixinsawsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMatchmakingRuleSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMatchmakingRuleSetPropsMixinParameters(props *CfnMatchmakingRuleSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

