//go:build no_runtime_type_checking

package mixinsawssupportapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAccountAliasPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAccountAliasPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAccountAliasPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAccountAliasPropsMixinParameters(props *CfnAccountAliasMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

