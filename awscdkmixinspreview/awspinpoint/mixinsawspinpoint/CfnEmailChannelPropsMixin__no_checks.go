//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEmailChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEmailChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEmailChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEmailChannelPropsMixinParameters(props *CfnEmailChannelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

