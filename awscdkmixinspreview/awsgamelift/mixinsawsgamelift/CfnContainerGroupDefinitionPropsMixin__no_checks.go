//go:build no_runtime_type_checking

package mixinsawsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContainerGroupDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContainerGroupDefinitionPropsMixinParameters(props *CfnContainerGroupDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

