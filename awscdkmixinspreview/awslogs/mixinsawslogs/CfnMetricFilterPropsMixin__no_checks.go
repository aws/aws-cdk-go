//go:build no_runtime_type_checking

package mixinsawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMetricFilterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMetricFilterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMetricFilterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMetricFilterPropsMixinParameters(props *CfnMetricFilterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

