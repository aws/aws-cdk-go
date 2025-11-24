//go:build no_runtime_type_checking

package mixinsawscassandra

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTablePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTablePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTablePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTablePropsMixinParameters(props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

