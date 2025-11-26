//go:build no_runtime_type_checking

package previewawsopsworksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstancePropsMixinParameters(props *CfnInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

