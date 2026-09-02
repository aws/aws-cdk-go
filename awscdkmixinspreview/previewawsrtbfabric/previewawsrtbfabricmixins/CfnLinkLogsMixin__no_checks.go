//go:build no_runtime_type_checking

package previewawsrtbfabricmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLinkLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLinkLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLinkLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLinkLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

