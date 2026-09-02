//go:build no_runtime_type_checking

package awsdrs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSourceNetworkPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSourceNetworkPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSourceNetworkPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSourceNetworkPropsMixinParameters(props *CfnSourceNetworkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

