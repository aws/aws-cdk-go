//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHookVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHookVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHookVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHookVersionPropsMixinParameters(props *CfnHookVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

