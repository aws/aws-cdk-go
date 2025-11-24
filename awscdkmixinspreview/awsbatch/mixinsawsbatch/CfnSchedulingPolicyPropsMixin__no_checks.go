//go:build no_runtime_type_checking

package mixinsawsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSchedulingPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSchedulingPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSchedulingPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSchedulingPolicyPropsMixinParameters(props *CfnSchedulingPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

