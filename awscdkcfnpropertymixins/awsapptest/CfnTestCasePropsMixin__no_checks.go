//go:build no_runtime_type_checking

package awsapptest

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTestCasePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTestCasePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTestCasePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTestCasePropsMixinParameters(props *CfnTestCaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

