//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFlowLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFlowLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFlowLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFlowLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

