//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCPeeringConnectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCPeeringConnectionPropsMixinParameters(props *CfnVPCPeeringConnectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

