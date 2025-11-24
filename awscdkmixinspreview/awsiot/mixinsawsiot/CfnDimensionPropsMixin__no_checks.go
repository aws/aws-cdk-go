//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDimensionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDimensionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDimensionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDimensionPropsMixinParameters(props *CfnDimensionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

