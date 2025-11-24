//go:build no_runtime_type_checking

package mixinsawssqs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnQueuePolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnQueuePolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnQueuePolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnQueuePolicyPropsMixinParameters(props *CfnQueuePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

