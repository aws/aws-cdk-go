//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTaskSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTaskSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTaskSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTaskSetPropsMixinParameters(props *CfnTaskSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

