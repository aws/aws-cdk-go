//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMailManagerIngressPointLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMailManagerIngressPointLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMailManagerIngressPointLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

