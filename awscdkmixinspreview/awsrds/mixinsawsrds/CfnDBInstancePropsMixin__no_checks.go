//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBInstancePropsMixinParameters(props *CfnDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

