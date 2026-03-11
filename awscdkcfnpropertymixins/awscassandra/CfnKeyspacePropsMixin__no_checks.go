//go:build no_runtime_type_checking

package awscassandra

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnKeyspacePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnKeyspacePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnKeyspacePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnKeyspacePropsMixinParameters(props *CfnKeyspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

