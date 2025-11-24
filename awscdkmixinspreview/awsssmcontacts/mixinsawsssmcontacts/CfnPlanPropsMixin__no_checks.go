//go:build no_runtime_type_checking

package mixinsawsssmcontacts

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlanPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPlanPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPlanPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPlanPropsMixinParameters(props *CfnPlanMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

