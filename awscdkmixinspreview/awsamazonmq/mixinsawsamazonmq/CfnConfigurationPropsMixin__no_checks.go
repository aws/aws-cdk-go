//go:build no_runtime_type_checking

package mixinsawsamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigurationPropsMixinParameters(props *CfnConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

