//go:build no_runtime_type_checking

package previewawslogsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIntegrationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIntegrationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIntegrationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIntegrationPropsMixinParameters(props *CfnIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

