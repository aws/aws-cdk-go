//go:build no_runtime_type_checking

package awssso

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTrustedTokenIssuerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTrustedTokenIssuerPropsMixinParameters(props *CfnTrustedTokenIssuerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

