//go:build no_runtime_type_checking

package mixinsawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAccountPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAccountPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAccountPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAccountPolicyPropsMixinParameters(props *CfnAccountPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

