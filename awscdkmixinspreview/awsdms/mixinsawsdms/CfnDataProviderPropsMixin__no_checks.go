//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDataProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDataProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDataProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDataProviderPropsMixinParameters(props *CfnDataProviderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

