//go:build no_runtime_type_checking

package mixinsawsgreengrass

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLoggerDefinitionVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLoggerDefinitionVersionPropsMixinParameters(props *CfnLoggerDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

