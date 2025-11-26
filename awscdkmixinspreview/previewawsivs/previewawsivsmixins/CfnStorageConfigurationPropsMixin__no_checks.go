//go:build no_runtime_type_checking

package previewawsivsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStorageConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStorageConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStorageConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStorageConfigurationPropsMixinParameters(props *CfnStorageConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

