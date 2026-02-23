//go:build no_runtime_type_checking

package previewawsdevopsagentmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentSpaceLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgentSpaceLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgentSpaceLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgentSpaceLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

