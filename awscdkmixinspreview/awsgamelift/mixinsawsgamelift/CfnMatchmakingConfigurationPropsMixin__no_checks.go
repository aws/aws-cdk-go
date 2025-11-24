//go:build no_runtime_type_checking

package mixinsawsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMatchmakingConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMatchmakingConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMatchmakingConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMatchmakingConfigurationPropsMixinParameters(props *CfnMatchmakingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

