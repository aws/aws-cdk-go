//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgentLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgentLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgentLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

