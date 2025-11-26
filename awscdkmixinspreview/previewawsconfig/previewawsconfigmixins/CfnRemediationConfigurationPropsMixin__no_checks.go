//go:build no_runtime_type_checking

package previewawsconfigmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRemediationConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRemediationConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRemediationConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRemediationConfigurationPropsMixinParameters(props *CfnRemediationConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

