//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContactFlowModulePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContactFlowModulePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContactFlowModulePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContactFlowModulePropsMixinParameters(props *CfnContactFlowModuleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

