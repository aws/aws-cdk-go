//go:build no_runtime_type_checking

package mixinsawsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAutoScalingConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAutoScalingConfigurationPropsMixinParameters(props *CfnAutoScalingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

