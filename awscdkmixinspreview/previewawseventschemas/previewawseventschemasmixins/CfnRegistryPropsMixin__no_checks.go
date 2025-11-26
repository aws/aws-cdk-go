//go:build no_runtime_type_checking

package previewawseventschemasmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRegistryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRegistryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRegistryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRegistryPropsMixinParameters(props *CfnRegistryMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

