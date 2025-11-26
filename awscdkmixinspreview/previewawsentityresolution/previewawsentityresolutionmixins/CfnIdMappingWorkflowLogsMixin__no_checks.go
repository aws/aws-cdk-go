//go:build no_runtime_type_checking

package previewawsentityresolutionmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIdMappingWorkflowLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIdMappingWorkflowLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIdMappingWorkflowLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIdMappingWorkflowLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

