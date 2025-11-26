//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVerifiedAccessInstanceLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVerifiedAccessInstanceLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVerifiedAccessInstanceLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVerifiedAccessInstanceLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

