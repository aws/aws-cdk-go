//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMeshPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMeshPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMeshPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMeshPropsMixinParameters(props *CfnMeshMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

