//go:build no_runtime_type_checking

package previewawscleanroomsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMembershipLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMembershipLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMembershipLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

