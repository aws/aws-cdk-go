//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PropertyInjectors) validateForParameters(uniqueId *string) error {
	return nil
}

func validatePropertyInjectors_HasPropertyInjectorsParameters(x interface{}) error {
	return nil
}

func validatePropertyInjectors_OfParameters(scope constructs.IConstruct) error {
	return nil
}

