//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventSubscriptionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventSubscriptionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventSubscriptionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventSubscriptionPropsMixinParameters(props *CfnEventSubscriptionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

