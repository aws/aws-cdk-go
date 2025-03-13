//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ISynthesisSession) validateSetAssemblyParameters(val cxapi.CloudAssemblyBuilder) error {
	return nil
}

func (j *jsiiProxy_ISynthesisSession) validateSetOutdirParameters(val *string) error {
	return nil
}

