//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTopicPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTopicPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTopicPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTopicPolicyPropsMixinParameters(props *CfnTopicPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

