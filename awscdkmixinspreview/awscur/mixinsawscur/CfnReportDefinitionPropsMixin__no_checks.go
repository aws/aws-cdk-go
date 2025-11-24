//go:build no_runtime_type_checking

package mixinsawscur

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReportDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReportDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReportDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReportDefinitionPropsMixinParameters(props *CfnReportDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

