//go:build no_runtime_type_checking

package previewawsbackupmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTieringConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTieringConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTieringConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTieringConfigurationPropsMixinParameters(props *CfnTieringConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

