//go:build no_runtime_type_checking

package mixinsawsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScheduleGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnScheduleGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnScheduleGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnScheduleGroupPropsMixinParameters(props *CfnScheduleGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

