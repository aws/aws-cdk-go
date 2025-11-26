//go:build no_runtime_type_checking

package previewawsconfigmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigRulePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigRulePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigRulePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigRulePropsMixinParameters(props *CfnConfigRuleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

