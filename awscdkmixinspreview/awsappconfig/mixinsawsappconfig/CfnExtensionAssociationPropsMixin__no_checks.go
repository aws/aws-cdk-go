//go:build no_runtime_type_checking

package mixinsawsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExtensionAssociationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExtensionAssociationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExtensionAssociationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExtensionAssociationPropsMixinParameters(props *CfnExtensionAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

