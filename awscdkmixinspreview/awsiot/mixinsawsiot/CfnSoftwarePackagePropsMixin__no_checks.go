//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSoftwarePackagePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSoftwarePackagePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSoftwarePackagePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSoftwarePackagePropsMixinParameters(props *CfnSoftwarePackageMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

