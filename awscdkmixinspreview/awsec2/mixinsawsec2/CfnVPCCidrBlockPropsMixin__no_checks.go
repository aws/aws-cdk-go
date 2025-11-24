//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCCidrBlockPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCCidrBlockPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCCidrBlockPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCCidrBlockPropsMixinParameters(props *CfnVPCCidrBlockMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

