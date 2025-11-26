//go:build no_runtime_type_checking

package previewawsgreengrassmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCoreDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCoreDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCoreDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCoreDefinitionPropsMixinParameters(props *CfnCoreDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

