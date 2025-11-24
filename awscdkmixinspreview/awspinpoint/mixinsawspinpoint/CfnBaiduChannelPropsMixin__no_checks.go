//go:build no_runtime_type_checking

package mixinsawspinpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBaiduChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBaiduChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBaiduChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBaiduChannelPropsMixinParameters(props *CfnBaiduChannelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

