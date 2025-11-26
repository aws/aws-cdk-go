//go:build no_runtime_type_checking

package previewawsconnectmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

