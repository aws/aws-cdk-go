//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMaintenanceWindowPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMaintenanceWindowPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMaintenanceWindowPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMaintenanceWindowPropsMixinParameters(props *CfnMaintenanceWindowMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

