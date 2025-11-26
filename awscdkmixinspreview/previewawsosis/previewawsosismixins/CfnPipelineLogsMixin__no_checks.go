//go:build no_runtime_type_checking

package previewawsosismixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPipelineLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPipelineLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPipelineLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPipelineLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

