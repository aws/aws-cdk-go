//go:build no_runtime_type_checking

package mixinsawsdatabrew

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRecipePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRecipePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRecipePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRecipePropsMixinParameters(props *CfnRecipeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

