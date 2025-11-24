//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFlowAliasPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFlowAliasPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFlowAliasPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFlowAliasPropsMixinParameters(props *CfnFlowAliasMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

