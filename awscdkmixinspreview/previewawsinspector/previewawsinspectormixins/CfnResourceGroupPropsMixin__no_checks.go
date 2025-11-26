//go:build no_runtime_type_checking

package previewawsinspectormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourceGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceGroupPropsMixinParameters(props *CfnResourceGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

