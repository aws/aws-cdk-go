//go:build no_runtime_type_checking

package mixinsawstimestream

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDatabasePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDatabasePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDatabasePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDatabasePropsMixinParameters(props *CfnDatabaseMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

