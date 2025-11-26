//go:build no_runtime_type_checking

package previewawsevsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEnvironmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEnvironmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEnvironmentPropsMixinParameters(props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

