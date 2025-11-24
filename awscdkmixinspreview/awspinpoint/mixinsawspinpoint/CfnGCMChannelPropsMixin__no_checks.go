//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGCMChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGCMChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGCMChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGCMChannelPropsMixinParameters(props *CfnGCMChannelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

