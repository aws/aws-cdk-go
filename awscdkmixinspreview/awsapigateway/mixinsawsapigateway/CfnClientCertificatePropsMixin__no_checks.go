//go:build no_runtime_type_checking

package mixinsawsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClientCertificatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnClientCertificatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnClientCertificatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnClientCertificatePropsMixinParameters(props *CfnClientCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

