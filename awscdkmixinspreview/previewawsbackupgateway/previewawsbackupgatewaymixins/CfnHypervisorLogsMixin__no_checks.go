//go:build no_runtime_type_checking

package previewawsbackupgatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHypervisorLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHypervisorLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHypervisorLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

