//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProtectionGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProtectionGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProtectionGroupPropsMixinParameters(props *CfnProtectionGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

