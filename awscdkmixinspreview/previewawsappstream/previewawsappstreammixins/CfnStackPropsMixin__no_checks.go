//go:build no_runtime_type_checking

package previewawsappstreammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStackPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStackPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStackPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStackPropsMixinParameters(props *CfnStackMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

