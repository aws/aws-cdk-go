//go:build no_runtime_type_checking

package previewawsvpclatticemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServiceLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServiceLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServiceLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServiceLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

