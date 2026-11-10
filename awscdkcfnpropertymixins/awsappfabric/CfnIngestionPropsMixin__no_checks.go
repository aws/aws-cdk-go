//go:build no_runtime_type_checking

package awsappfabric

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIngestionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIngestionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIngestionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIngestionPropsMixinParameters(props *CfnIngestionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

