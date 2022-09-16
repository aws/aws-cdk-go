//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha

import (
	"fmt"
)

func validateRuntime_OfParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

