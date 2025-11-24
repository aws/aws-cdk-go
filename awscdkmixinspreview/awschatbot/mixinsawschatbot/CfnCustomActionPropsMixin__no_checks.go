//go:build no_runtime_type_checking

package mixinsawschatbot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomActionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCustomActionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomActionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCustomActionPropsMixinParameters(props *CfnCustomActionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

