//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPolicyPropsMixinParameters(props *CfnUserPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

