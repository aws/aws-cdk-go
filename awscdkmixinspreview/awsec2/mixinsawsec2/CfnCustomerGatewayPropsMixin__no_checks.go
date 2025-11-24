//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomerGatewayPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCustomerGatewayPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomerGatewayPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCustomerGatewayPropsMixinParameters(props *CfnCustomerGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

