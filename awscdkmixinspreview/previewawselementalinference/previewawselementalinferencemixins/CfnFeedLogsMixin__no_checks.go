//go:build no_runtime_type_checking

package previewawselementalinferencemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFeedLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFeedLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFeedLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFeedLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

