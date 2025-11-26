//go:build no_runtime_type_checking

package previewawsentityresolutionmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMatchingWorkflowLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMatchingWorkflowLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMatchingWorkflowLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMatchingWorkflowLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

