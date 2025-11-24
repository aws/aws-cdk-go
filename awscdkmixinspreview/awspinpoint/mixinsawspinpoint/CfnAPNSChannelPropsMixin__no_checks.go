//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAPNSChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAPNSChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAPNSChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAPNSChannelPropsMixinParameters(props *CfnAPNSChannelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

