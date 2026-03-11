//go:build no_runtime_type_checking

package awsgreengrass

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFunctionDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFunctionDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFunctionDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFunctionDefinitionPropsMixinParameters(props *CfnFunctionDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

