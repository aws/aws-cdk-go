//go:build no_runtime_type_checking

package previewawskafkaconnectmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConnectorLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConnectorLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConnectorLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConnectorLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

