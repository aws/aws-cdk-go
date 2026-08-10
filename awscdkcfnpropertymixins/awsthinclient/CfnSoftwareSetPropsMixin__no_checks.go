//go:build no_runtime_type_checking

package awsthinclient

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSoftwareSetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSoftwareSetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSoftwareSetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSoftwareSetPropsMixinParameters(props *CfnSoftwareSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

