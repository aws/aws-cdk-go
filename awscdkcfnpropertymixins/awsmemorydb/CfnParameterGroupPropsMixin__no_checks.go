//go:build no_runtime_type_checking

package awsmemorydb

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnParameterGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnParameterGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnParameterGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnParameterGroupPropsMixinParameters(props *CfnParameterGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

