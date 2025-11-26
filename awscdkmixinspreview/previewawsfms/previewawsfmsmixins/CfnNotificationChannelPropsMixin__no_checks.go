//go:build no_runtime_type_checking

package previewawsfmsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNotificationChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNotificationChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNotificationChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNotificationChannelPropsMixinParameters(props *CfnNotificationChannelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

