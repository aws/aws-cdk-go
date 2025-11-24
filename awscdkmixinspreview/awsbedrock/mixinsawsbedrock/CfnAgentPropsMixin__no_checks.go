//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgentPropsMixinParameters(props *CfnAgentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

