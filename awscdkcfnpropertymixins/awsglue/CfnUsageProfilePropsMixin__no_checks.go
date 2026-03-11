//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUsageProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUsageProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUsageProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUsageProfilePropsMixinParameters(props *CfnUsageProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

