//go:build no_runtime_type_checking

package mixinsawsdatazone

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserProfilePropsMixinParameters(props *CfnUserProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

