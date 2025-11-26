//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssistantLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAssistantLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAssistantLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

