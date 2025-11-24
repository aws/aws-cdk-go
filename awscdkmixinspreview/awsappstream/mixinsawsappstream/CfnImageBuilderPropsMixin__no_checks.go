//go:build no_runtime_type_checking

package mixinsawsappstream

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnImageBuilderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnImageBuilderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnImageBuilderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnImageBuilderPropsMixinParameters(props *CfnImageBuilderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

