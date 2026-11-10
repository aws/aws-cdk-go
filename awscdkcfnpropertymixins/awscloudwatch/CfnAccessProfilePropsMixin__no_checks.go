//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAccessProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAccessProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAccessProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAccessProfilePropsMixinParameters(props *CfnAccessProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

