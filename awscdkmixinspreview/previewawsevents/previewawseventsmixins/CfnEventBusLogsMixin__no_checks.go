//go:build no_runtime_type_checking

package previewawseventsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventBusLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventBusLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventBusLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

