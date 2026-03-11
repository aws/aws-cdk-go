//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScheduledQueryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnScheduledQueryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnScheduledQueryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnScheduledQueryPropsMixinParameters(props *CfnScheduledQueryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

