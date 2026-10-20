//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBClusterEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBClusterEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBClusterEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBClusterEndpointPropsMixinParameters(props *CfnDBClusterEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

