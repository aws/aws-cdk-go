//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContactFlowPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContactFlowPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContactFlowPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContactFlowPropsMixinParameters(props *CfnContactFlowMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

