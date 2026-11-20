//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeAliasPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeAliasPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRuntimeAliasPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRuntimeAliasPropsMixinParameters(props *CfnRuntimeAliasMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

