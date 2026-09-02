//go:build no_runtime_type_checking

package awschime

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnChannelFlowPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnChannelFlowPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnChannelFlowPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnChannelFlowPropsMixinParameters(props *CfnChannelFlowMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

