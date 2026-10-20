//go:build no_runtime_type_checking

package awsworkspaces

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDirectoryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDirectoryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDirectoryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDirectoryPropsMixinParameters(props *CfnDirectoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

