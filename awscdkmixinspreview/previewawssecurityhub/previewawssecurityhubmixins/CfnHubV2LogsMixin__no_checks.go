//go:build no_runtime_type_checking

package previewawssecurityhubmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHubV2LogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHubV2LogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHubV2LogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHubV2LogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

