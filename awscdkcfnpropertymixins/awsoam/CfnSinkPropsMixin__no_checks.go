//go:build no_runtime_type_checking

package awsoam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSinkPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSinkPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSinkPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSinkPropsMixinParameters(props *CfnSinkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

