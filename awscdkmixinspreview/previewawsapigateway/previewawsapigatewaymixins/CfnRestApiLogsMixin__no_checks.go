//go:build no_runtime_type_checking

package previewawsapigatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRestApiLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRestApiLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRestApiLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

