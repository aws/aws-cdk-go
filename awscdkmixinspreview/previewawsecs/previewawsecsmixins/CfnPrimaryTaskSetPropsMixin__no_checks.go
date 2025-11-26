//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPrimaryTaskSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPrimaryTaskSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPrimaryTaskSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPrimaryTaskSetPropsMixinParameters(props *CfnPrimaryTaskSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

