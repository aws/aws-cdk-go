//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLogGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLogGroupPropsMixinParameters(props *CfnLogGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

