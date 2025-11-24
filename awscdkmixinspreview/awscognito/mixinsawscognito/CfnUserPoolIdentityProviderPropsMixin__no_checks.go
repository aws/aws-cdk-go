//go:build no_runtime_type_checking

package mixinsawscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPoolIdentityProviderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPoolIdentityProviderPropsMixinParameters(props *CfnUserPoolIdentityProviderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

