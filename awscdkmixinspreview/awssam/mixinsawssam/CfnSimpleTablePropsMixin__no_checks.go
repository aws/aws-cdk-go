//go:build no_runtime_type_checking

package mixinsawssam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSimpleTablePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSimpleTablePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSimpleTablePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSimpleTablePropsMixinParameters(props *CfnSimpleTableMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

