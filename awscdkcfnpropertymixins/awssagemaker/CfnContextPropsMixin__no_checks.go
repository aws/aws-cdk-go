//go:build no_runtime_type_checking

package awssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContextPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContextPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContextPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContextPropsMixinParameters(props *CfnContextMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

