//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNetworkConnectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNetworkConnectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNetworkConnectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNetworkConnectorPropsMixinParameters(props *CfnNetworkConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

