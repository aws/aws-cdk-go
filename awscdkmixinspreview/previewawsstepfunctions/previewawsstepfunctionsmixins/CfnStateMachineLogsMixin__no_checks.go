//go:build no_runtime_type_checking

package previewawsstepfunctionsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStateMachineLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStateMachineLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStateMachineLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStateMachineLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

