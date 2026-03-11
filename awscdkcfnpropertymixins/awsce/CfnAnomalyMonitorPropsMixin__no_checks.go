//go:build no_runtime_type_checking

package awsce

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAnomalyMonitorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAnomalyMonitorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAnomalyMonitorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAnomalyMonitorPropsMixinParameters(props *CfnAnomalyMonitorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

