//go:build no_runtime_type_checking

package previewawsivsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPublicKeyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPublicKeyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPublicKeyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPublicKeyPropsMixinParameters(props *CfnPublicKeyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

