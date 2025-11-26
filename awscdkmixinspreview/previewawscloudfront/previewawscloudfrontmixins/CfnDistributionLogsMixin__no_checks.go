//go:build no_runtime_type_checking

package previewawscloudfrontmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDistributionLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDistributionLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDistributionLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

