//go:build no_runtime_type_checking

package mixinsawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDeliveryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDeliveryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDeliveryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDeliveryPropsMixinParameters(props *CfnDeliveryMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

