//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceConnectEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceConnectEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceConnectEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceConnectEndpointPropsMixinParameters(props *CfnInstanceConnectEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

