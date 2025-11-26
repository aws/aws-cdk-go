//go:build no_runtime_type_checking

package previewawsnotificationsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNotificationHubPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNotificationHubPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNotificationHubPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNotificationHubPropsMixinParameters(props *CfnNotificationHubMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

