//go:build no_runtime_type_checking

package mixinsawscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolClientPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolClientPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPoolClientPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPoolClientPropsMixinParameters(props *CfnUserPoolClientMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

