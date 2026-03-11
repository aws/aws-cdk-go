//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTransitGatewayRoutePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTransitGatewayRoutePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTransitGatewayRoutePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTransitGatewayRoutePropsMixinParameters(props *CfnTransitGatewayRouteMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

