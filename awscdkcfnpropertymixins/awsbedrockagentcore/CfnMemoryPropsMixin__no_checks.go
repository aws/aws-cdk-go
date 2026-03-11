//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMemoryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMemoryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMemoryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMemoryPropsMixinParameters(props *CfnMemoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

