//go:build no_runtime_type_checking

package mixinsawsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVirtualGatewayPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVirtualGatewayPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVirtualGatewayPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVirtualGatewayPropsMixinParameters(props *CfnVirtualGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

