//go:build no_runtime_type_checking

package awsdrs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLaunchConfigurationTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLaunchConfigurationTemplatePropsMixinParameters(props *CfnLaunchConfigurationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

