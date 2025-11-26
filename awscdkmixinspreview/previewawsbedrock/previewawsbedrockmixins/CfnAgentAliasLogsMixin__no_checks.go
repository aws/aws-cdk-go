//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentAliasLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgentAliasLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgentAliasLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

