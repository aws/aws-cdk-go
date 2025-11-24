//go:build no_runtime_type_checking

package mixinsawscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAnomalyDetectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAnomalyDetectorPropsMixinParameters(props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

