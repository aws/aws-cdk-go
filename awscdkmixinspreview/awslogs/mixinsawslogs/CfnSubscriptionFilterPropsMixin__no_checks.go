//go:build no_runtime_type_checking

package mixinsawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubscriptionFilterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubscriptionFilterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubscriptionFilterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubscriptionFilterPropsMixinParameters(props *CfnSubscriptionFilterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

