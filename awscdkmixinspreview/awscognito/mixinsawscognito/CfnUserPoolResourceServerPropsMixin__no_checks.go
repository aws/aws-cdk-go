//go:build no_runtime_type_checking

package mixinsawscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolResourceServerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolResourceServerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPoolResourceServerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPoolResourceServerPropsMixinParameters(props *CfnUserPoolResourceServerMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

