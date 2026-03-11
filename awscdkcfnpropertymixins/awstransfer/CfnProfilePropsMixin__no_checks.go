//go:build no_runtime_type_checking

package awstransfer

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProfilePropsMixinParameters(props *CfnProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

