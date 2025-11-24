//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Mixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (m *jsiiProxy_Mixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMixin_IsMixinParameters(x interface{}) error {
	return nil
}

