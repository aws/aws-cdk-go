//go:build no_runtime_type_checking

package mixinsawsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAlarmModelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAlarmModelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAlarmModelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAlarmModelPropsMixinParameters(props *CfnAlarmModelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

