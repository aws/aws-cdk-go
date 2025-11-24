//go:build no_runtime_type_checking

package mixinsawscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCompositeAlarmPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCompositeAlarmPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCompositeAlarmPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCompositeAlarmPropsMixinParameters(props *CfnCompositeAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

