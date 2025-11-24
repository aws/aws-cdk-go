//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnThingPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnThingPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnThingPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnThingPropsMixinParameters(props *CfnThingMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

