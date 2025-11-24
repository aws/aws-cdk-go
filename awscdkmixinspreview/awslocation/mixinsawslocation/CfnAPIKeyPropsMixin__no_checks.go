//go:build no_runtime_type_checking

package mixinsawslocation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAPIKeyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAPIKeyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAPIKeyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAPIKeyPropsMixinParameters(props *CfnAPIKeyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

