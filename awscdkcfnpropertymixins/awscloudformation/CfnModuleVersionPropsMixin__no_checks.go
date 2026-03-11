//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnModuleVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnModuleVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnModuleVersionPropsMixinParameters(props *CfnModuleVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

