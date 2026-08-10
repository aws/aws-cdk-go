//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserDefinedFunctionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserDefinedFunctionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserDefinedFunctionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserDefinedFunctionPropsMixinParameters(props *CfnUserDefinedFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

