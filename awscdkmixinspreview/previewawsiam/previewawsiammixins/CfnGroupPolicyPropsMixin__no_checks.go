//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGroupPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGroupPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGroupPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGroupPolicyPropsMixinParameters(props *CfnGroupPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

