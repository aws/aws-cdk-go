//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGatewayPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGatewayPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGatewayPropsMixinParameters(props *CfnGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

