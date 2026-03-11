//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOIDCProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnOIDCProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnOIDCProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnOIDCProviderPropsMixinParameters(props *CfnOIDCProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

