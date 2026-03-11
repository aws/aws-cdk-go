//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProfilePermissionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProfilePermissionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProfilePermissionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProfilePermissionPropsMixinParameters(props *CfnProfilePermissionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

