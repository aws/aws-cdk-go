//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogStreamPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLogStreamPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLogStreamPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLogStreamPropsMixinParameters(props *CfnLogStreamMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

