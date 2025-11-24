//go:build no_runtime_type_checking

package mixinsawsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigurationProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigurationProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigurationProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigurationProfilePropsMixinParameters(props *CfnConfigurationProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

