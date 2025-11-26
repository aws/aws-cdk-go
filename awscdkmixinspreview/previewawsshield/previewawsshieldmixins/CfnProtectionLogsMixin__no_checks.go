//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProtectionLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProtectionLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProtectionLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

