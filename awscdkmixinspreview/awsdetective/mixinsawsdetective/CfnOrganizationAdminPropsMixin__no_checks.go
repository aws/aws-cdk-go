//go:build no_runtime_type_checking

package mixinsawsdetective

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOrganizationAdminPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnOrganizationAdminPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnOrganizationAdminPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnOrganizationAdminPropsMixinParameters(props *CfnOrganizationAdminMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

