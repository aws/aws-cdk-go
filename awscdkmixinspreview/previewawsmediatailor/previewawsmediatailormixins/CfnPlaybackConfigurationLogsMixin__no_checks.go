//go:build no_runtime_type_checking

package previewawsmediatailormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlaybackConfigurationLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPlaybackConfigurationLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPlaybackConfigurationLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

