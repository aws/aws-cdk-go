//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnManagedPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnManagedPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnManagedPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnManagedPolicyPropsMixinParameters(props *CfnManagedPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

