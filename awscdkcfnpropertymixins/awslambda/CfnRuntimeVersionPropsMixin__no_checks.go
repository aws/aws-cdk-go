//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRuntimeVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRuntimeVersionPropsMixinParameters(props *CfnRuntimeVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

