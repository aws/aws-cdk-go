//go:build no_runtime_type_checking

package mixinsawsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGatewayResponsePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayResponsePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGatewayResponsePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGatewayResponsePropsMixinParameters(props *CfnGatewayResponseMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

