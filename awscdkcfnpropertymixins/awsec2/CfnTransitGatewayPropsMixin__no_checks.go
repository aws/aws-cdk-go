//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTransitGatewayPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTransitGatewayPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTransitGatewayPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTransitGatewayPropsMixinParameters(props *CfnTransitGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

