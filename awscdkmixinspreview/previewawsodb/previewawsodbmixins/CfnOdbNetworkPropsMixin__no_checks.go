//go:build no_runtime_type_checking

package previewawsodbmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOdbNetworkPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnOdbNetworkPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnOdbNetworkPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnOdbNetworkPropsMixinParameters(props *CfnOdbNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

