//go:build no_runtime_type_checking

package previewawsapsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkspaceLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkspaceLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkspaceLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkspaceLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

