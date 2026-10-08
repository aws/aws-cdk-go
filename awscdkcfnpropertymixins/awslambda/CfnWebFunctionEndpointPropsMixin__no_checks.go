//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebFunctionEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWebFunctionEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWebFunctionEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWebFunctionEndpointPropsMixinParameters(props *CfnWebFunctionEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

