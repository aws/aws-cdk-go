//go:build no_runtime_type_checking

package previewawsnotificationsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventRulePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventRulePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventRulePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventRulePropsMixinParameters(props *CfnEventRuleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

