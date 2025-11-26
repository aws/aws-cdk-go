//go:build no_runtime_type_checking

package previewawsivschatmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRoomLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRoomLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRoomLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRoomLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

