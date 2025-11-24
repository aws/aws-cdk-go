//go:build no_runtime_type_checking

package mixinsawsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDomainNameV2PropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDomainNameV2PropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDomainNameV2PropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDomainNameV2PropsMixinParameters(props *CfnDomainNameV2MixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

