//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCEndpointServicePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCEndpointServicePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCEndpointServicePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCEndpointServicePropsMixinParameters(props *CfnVPCEndpointServiceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

