//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMigrationProjectPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMigrationProjectPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMigrationProjectPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMigrationProjectPropsMixinParameters(props *CfnMigrationProjectMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

