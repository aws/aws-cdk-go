//go:build no_runtime_type_checking

package previewawslogsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogAnomalyDetectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLogAnomalyDetectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLogAnomalyDetectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLogAnomalyDetectorPropsMixinParameters(props *CfnLogAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

