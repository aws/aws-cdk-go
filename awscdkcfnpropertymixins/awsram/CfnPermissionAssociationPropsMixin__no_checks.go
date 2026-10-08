//go:build no_runtime_type_checking

package awsram

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPermissionAssociationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPermissionAssociationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPermissionAssociationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPermissionAssociationPropsMixinParameters(props *CfnPermissionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

