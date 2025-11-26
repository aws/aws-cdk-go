//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGatewayLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGatewayLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGatewayLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

