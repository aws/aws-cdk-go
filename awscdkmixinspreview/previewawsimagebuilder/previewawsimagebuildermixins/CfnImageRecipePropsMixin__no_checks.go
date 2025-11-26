//go:build no_runtime_type_checking

package previewawsimagebuildermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnImageRecipePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnImageRecipePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnImageRecipePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnImageRecipePropsMixinParameters(props *CfnImageRecipeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

