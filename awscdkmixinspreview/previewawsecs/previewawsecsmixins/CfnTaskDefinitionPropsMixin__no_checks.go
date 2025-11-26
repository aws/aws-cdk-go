//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTaskDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTaskDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTaskDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTaskDefinitionPropsMixinParameters(props *CfnTaskDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

