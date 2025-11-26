//go:build no_runtime_type_checking

package previewawsgreengrassmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGroupVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGroupVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGroupVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGroupVersionPropsMixinParameters(props *CfnGroupVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

