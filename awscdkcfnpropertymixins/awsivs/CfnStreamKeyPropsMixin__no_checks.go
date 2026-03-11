//go:build no_runtime_type_checking

package awsivs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStreamKeyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStreamKeyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStreamKeyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStreamKeyPropsMixinParameters(props *CfnStreamKeyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

