//go:build no_runtime_type_checking

package awsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnThingGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnThingGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnThingGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnThingGroupPropsMixinParameters(props *CfnThingGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

