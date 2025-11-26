//go:build no_runtime_type_checking

package previewawsivsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStagePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStagePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStagePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStagePropsMixinParameters(props *CfnStageMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

