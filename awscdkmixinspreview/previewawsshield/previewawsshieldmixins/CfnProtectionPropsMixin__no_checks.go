//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProtectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProtectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProtectionPropsMixinParameters(props *CfnProtectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

