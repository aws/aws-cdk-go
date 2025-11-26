//go:build no_runtime_type_checking

package previewawsmskmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnClusterLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnClusterLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnClusterLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

