//go:build no_runtime_type_checking

package mixinsawsemr

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceFleetConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceFleetConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceFleetConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceFleetConfigPropsMixinParameters(props *CfnInstanceFleetConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

