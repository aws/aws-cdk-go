//go:build no_runtime_type_checking

package previewawsapsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScraperLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnScraperLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnScraperLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnScraperLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

