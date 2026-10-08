//go:build no_runtime_type_checking

package awsram

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPrincipalAssociationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPrincipalAssociationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPrincipalAssociationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPrincipalAssociationPropsMixinParameters(props *CfnPrincipalAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

