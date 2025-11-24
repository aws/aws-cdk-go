//go:build no_runtime_type_checking

package mixinsawsce

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCostCategoryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCostCategoryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCostCategoryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCostCategoryPropsMixinParameters(props *CfnCostCategoryMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

