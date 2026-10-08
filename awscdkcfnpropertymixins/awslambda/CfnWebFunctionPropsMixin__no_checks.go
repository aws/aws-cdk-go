//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebFunctionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWebFunctionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWebFunctionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWebFunctionPropsMixinParameters(props *CfnWebFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

