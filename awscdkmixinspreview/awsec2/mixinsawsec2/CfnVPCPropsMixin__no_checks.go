//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCPropsMixinParameters(props *CfnVPCMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

