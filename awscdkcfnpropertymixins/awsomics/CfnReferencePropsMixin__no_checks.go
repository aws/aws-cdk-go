//go:build no_runtime_type_checking

package awsomics

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReferencePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReferencePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReferencePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReferencePropsMixinParameters(props *CfnReferenceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

