//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogAlarmPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLogAlarmPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLogAlarmPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLogAlarmPropsMixinParameters(props *CfnLogAlarmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

