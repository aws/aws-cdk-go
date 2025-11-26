//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFirewallLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFirewallLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

