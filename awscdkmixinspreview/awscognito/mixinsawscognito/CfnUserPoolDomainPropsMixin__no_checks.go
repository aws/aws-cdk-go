//go:build no_runtime_type_checking

package mixinsawscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolDomainPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolDomainPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPoolDomainPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPoolDomainPropsMixinParameters(props *CfnUserPoolDomainMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

