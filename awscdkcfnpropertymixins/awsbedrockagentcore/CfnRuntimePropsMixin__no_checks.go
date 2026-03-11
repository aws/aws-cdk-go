//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRuntimePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRuntimePropsMixinParameters(props *CfnRuntimeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

