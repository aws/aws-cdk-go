//go:build no_runtime_type_checking

package previewawsconfigmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigurationRecorderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigurationRecorderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigurationRecorderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigurationRecorderPropsMixinParameters(props *CfnConfigurationRecorderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

