//go:build no_runtime_type_checking

package previewawssecurityhubmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHubLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHubLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHubLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHubLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

