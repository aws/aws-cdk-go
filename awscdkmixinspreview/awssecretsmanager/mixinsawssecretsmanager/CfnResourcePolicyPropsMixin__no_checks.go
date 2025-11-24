//go:build no_runtime_type_checking

package mixinsawssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourcePolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourcePolicyPropsMixinParameters(props *CfnResourcePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

