//go:build no_runtime_type_checking

package awscassandra

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTypePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTypePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTypePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTypePropsMixinParameters(props *CfnTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

