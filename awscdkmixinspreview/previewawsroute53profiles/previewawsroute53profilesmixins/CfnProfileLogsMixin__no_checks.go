//go:build no_runtime_type_checking

package previewawsroute53profilesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProfileLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnProfileLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnProfileLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnProfileLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

