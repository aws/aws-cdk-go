//go:build no_runtime_type_checking

package mixinsawskafkaconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomPluginPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCustomPluginPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomPluginPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCustomPluginPropsMixinParameters(props *CfnCustomPluginMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

