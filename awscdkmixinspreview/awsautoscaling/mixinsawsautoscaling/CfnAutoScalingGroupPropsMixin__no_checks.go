//go:build no_runtime_type_checking

package mixinsawsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAutoScalingGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAutoScalingGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAutoScalingGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAutoScalingGroupPropsMixinParameters(props *CfnAutoScalingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

