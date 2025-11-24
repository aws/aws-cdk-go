//go:build no_runtime_type_checking

package mixinsawsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLicenseEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLicenseEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLicenseEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLicenseEndpointPropsMixinParameters(props *CfnLicenseEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

