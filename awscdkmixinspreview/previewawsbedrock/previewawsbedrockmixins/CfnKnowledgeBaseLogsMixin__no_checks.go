//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnKnowledgeBaseLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnKnowledgeBaseLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnKnowledgeBaseLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

