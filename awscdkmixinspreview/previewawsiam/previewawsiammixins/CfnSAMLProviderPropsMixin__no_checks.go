//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSAMLProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSAMLProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSAMLProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSAMLProviderPropsMixinParameters(props *CfnSAMLProviderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

