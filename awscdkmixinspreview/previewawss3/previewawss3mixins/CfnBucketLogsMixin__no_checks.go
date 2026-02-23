//go:build no_runtime_type_checking

package previewawss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBucketLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBucketLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBucketLogsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBucketLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	return nil
}

