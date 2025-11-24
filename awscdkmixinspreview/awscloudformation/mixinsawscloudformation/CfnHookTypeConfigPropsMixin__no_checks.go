//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHookTypeConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHookTypeConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHookTypeConfigPropsMixinParameters(props *CfnHookTypeConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

