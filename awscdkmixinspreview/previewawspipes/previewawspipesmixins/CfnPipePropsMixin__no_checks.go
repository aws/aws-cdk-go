//go:build no_runtime_type_checking

package previewawspipesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPipePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPipePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPipePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPipePropsMixinParameters(props *CfnPipeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

