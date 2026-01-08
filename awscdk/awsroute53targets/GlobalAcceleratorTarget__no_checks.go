//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobalAcceleratorTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	return nil
}

func validateNewGlobalAcceleratorTargetParameters(accelerator interfacesawsglobalaccelerator.IAcceleratorRef) error {
	return nil
}

