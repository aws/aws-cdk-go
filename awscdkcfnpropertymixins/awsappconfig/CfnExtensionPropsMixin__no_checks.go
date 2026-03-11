//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExtensionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExtensionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExtensionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExtensionPropsMixinParameters(props *CfnExtensionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

