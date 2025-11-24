//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIntegrationAssociationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIntegrationAssociationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIntegrationAssociationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIntegrationAssociationPropsMixinParameters(props *CfnIntegrationAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

