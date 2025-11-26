//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAccessKeyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAccessKeyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAccessKeyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAccessKeyPropsMixinParameters(props *CfnAccessKeyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

