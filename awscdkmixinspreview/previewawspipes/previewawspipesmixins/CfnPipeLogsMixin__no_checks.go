//go:build no_runtime_type_checking

package previewawspipesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPipeLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPipeLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPipeLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPipeLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

