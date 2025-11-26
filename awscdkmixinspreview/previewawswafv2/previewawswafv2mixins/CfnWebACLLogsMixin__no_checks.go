//go:build no_runtime_type_checking

package previewawswafv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebACLLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWebACLLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWebACLLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

