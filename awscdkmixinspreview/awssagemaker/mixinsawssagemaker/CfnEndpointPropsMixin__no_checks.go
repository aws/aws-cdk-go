//go:build no_runtime_type_checking

package mixinsawssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEndpointPropsMixinParameters(props *CfnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

