//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInternetGatewayPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInternetGatewayPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInternetGatewayPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInternetGatewayPropsMixinParameters(props *CfnInternetGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

