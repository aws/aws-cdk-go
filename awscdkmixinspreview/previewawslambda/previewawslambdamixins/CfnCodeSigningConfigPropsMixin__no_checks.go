//go:build no_runtime_type_checking

package previewawslambdamixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeSigningConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCodeSigningConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCodeSigningConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCodeSigningConfigPropsMixinParameters(props *CfnCodeSigningConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

