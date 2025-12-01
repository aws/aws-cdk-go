//go:build no_runtime_type_checking

package previewawslambdamixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapacityProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCapacityProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCapacityProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCapacityProviderPropsMixinParameters(props *CfnCapacityProviderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

