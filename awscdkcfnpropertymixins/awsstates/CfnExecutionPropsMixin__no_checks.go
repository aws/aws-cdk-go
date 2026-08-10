//go:build no_runtime_type_checking

package awsstates

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExecutionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExecutionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExecutionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExecutionPropsMixinParameters(props *CfnExecutionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

