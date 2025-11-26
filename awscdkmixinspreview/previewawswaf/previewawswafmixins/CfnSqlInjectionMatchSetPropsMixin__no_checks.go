//go:build no_runtime_type_checking

package previewawswafmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSqlInjectionMatchSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSqlInjectionMatchSetPropsMixinParameters(props *CfnSqlInjectionMatchSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

