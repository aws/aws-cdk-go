//go:build no_runtime_type_checking

package previewawsbatchmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnJobDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnJobDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnJobDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnJobDefinitionPropsMixinParameters(props *CfnJobDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

