//go:build no_runtime_type_checking

package mixinsawssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSecretPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSecretPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSecretPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSecretPropsMixinParameters(props *CfnSecretMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

