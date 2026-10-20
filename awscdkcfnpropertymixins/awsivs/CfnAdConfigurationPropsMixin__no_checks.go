//go:build no_runtime_type_checking

package awsivs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAdConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAdConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAdConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAdConfigurationPropsMixinParameters(props *CfnAdConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

