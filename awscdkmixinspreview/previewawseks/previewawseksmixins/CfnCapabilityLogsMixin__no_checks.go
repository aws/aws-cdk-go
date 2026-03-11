//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCapabilityLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCapabilityLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

