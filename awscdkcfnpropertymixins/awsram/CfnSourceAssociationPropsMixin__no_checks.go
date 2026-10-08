//go:build no_runtime_type_checking

package awsram

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSourceAssociationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSourceAssociationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSourceAssociationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSourceAssociationPropsMixinParameters(props *CfnSourceAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

