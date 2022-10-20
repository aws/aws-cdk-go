//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Synthetics
package awscdksyntheticsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewRuntimeParameters(name *string, family RuntimeFamily) error {
	return nil
}

