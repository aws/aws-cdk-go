//go:build no_runtime_type_checking

package mixinsawssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEndpointConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEndpointConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEndpointConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEndpointConfigPropsMixinParameters(props *CfnEndpointConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

