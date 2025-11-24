//go:build no_runtime_type_checking

package mixinsawskafkaconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkerConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkerConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkerConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkerConfigurationPropsMixinParameters(props *CfnWorkerConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

