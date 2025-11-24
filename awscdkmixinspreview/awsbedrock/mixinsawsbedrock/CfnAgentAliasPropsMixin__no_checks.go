//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentAliasPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgentAliasPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgentAliasPropsMixinParameters(props *CfnAgentAliasMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

