//go:build no_runtime_type_checking

package previewawstransfermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebAppPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWebAppPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWebAppPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWebAppPropsMixinParameters(props *CfnWebAppMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

