//go:build no_runtime_type_checking

package previewawsb2bimixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTransformerLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTransformerLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTransformerLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

