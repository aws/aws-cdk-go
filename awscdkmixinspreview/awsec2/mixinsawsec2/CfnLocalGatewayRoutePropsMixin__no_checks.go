//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLocalGatewayRoutePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLocalGatewayRoutePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLocalGatewayRoutePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLocalGatewayRoutePropsMixinParameters(props *CfnLocalGatewayRouteMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

