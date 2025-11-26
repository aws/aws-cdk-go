//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRolePolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRolePolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRolePolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRolePolicyPropsMixinParameters(props *CfnRolePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

