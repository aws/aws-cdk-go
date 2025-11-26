//go:build no_runtime_type_checking

package previewawsvpclatticemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceConfigurationLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourceConfigurationLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceConfigurationLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceConfigurationLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

