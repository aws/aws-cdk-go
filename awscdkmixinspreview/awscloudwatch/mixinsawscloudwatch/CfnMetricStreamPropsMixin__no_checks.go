//go:build no_runtime_type_checking

package mixinsawscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMetricStreamPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMetricStreamPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMetricStreamPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMetricStreamPropsMixinParameters(props *CfnMetricStreamMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

