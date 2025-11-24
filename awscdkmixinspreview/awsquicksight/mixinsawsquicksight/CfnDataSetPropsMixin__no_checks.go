//go:build no_runtime_type_checking

package mixinsawsquicksight

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDataSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDataSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDataSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDataSetPropsMixinParameters(props *CfnDataSetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

