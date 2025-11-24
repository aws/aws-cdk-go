//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBProxyEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBProxyEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBProxyEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBProxyEndpointPropsMixinParameters(props *CfnDBProxyEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

