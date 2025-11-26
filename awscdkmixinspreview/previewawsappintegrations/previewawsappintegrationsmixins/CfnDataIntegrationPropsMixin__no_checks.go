//go:build no_runtime_type_checking

package previewawsappintegrationsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDataIntegrationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDataIntegrationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDataIntegrationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDataIntegrationPropsMixinParameters(props *CfnDataIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

