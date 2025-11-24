//go:build no_runtime_type_checking

package mixinsawsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServerCertificatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServerCertificatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServerCertificatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServerCertificatePropsMixinParameters(props *CfnServerCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

