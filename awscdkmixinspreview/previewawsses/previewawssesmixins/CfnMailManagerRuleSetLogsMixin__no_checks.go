//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMailManagerRuleSetLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerRuleSetLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMailManagerRuleSetLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMailManagerRuleSetLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

