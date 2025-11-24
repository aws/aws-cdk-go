//go:build no_runtime_type_checking

package mixinsawseks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFargateProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFargateProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFargateProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFargateProfilePropsMixinParameters(props *CfnFargateProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

