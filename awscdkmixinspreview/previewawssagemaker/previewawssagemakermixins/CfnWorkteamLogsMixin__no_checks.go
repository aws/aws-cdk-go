//go:build no_runtime_type_checking

package previewawssagemakermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkteamLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkteamLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkteamLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkteamLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

