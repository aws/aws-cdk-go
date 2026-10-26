//go:build no_runtime_type_checking

package previewawsgameliftmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGameSessionQueueLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGameSessionQueueLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGameSessionQueueLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGameSessionQueueLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

