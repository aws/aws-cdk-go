//go:build no_runtime_type_checking

package previewawspinpointmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAppPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAppPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAppPropsMixinParameters(props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

