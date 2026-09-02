//go:build no_runtime_type_checking

package awscases

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCasePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCasePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCasePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCasePropsMixinParameters(props *CfnCaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

