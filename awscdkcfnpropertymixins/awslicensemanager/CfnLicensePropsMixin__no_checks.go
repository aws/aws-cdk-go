//go:build no_runtime_type_checking

package awslicensemanager

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLicensePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLicensePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLicensePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLicensePropsMixinParameters(props *CfnLicenseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

