//go:build no_runtime_type_checking

package previewawsrummixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppMonitorLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAppMonitorLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAppMonitorLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

