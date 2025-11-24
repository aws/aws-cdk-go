//go:build no_runtime_type_checking

package mixinsawsfrauddetector

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDetectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDetectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDetectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDetectorPropsMixinParameters(props *CfnDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

