//go:build no_runtime_type_checking

package mixinsawsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReportPlanPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReportPlanPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReportPlanPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReportPlanPropsMixinParameters(props *CfnReportPlanMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

