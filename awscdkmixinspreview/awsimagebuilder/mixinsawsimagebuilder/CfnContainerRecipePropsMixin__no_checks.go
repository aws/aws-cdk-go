//go:build no_runtime_type_checking

package mixinsawsimagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContainerRecipePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContainerRecipePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContainerRecipePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContainerRecipePropsMixinParameters(props *CfnContainerRecipeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

