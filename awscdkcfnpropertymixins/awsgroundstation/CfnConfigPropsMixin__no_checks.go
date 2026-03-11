//go:build no_runtime_type_checking

package awsgroundstation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConfigPropsMixinParameters(props *CfnConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

