//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubnetCidrBlockPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubnetCidrBlockPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubnetCidrBlockPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubnetCidrBlockPropsMixinParameters(props *CfnSubnetCidrBlockMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

