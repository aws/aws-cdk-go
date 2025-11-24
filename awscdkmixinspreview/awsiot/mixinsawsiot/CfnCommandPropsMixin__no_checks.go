//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCommandPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCommandPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCommandPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCommandPropsMixinParameters(props *CfnCommandMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

