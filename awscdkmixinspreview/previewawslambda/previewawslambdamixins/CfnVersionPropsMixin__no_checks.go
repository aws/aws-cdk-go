//go:build no_runtime_type_checking

package previewawslambdamixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVersionPropsMixinParameters(props *CfnVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

