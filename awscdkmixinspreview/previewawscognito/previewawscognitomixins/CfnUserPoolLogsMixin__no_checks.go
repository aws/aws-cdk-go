//go:build no_runtime_type_checking

package previewawscognitomixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserPoolLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserPoolLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

