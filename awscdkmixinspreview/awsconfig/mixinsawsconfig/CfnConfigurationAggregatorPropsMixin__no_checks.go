//go:build no_runtime_type_checking

package mixinsawsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigurationAggregatorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigurationAggregatorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigurationAggregatorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigurationAggregatorPropsMixinParameters(props *CfnConfigurationAggregatorMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

