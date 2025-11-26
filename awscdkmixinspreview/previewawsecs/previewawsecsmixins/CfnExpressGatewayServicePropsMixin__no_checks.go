//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExpressGatewayServicePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExpressGatewayServicePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExpressGatewayServicePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExpressGatewayServicePropsMixinParameters(props *CfnExpressGatewayServiceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

