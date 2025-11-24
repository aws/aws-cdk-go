//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDataMigrationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDataMigrationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDataMigrationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDataMigrationPropsMixinParameters(props *CfnDataMigrationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

