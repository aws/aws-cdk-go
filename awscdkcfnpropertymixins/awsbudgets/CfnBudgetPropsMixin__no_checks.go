//go:build no_runtime_type_checking

package awsbudgets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBudgetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBudgetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBudgetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBudgetPropsMixinParameters(props *CfnBudgetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

