//go:build no_runtime_type_checking

package previewawstransfermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServerPropsMixinParameters(props *CfnServerMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

