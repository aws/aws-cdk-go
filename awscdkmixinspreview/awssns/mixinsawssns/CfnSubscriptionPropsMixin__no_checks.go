//go:build no_runtime_type_checking

package mixinsawssns

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubscriptionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubscriptionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubscriptionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubscriptionPropsMixinParameters(props *CfnSubscriptionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

