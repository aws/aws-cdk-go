//go:build no_runtime_type_checking

package awssso

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnApplicationProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnApplicationProviderPropsMixinParameters(props *CfnApplicationProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

