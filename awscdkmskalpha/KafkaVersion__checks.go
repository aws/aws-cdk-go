//go:build !no_runtime_type_checking

package awscdkmskalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateKafkaVersion_OfParameters(version *string, features *KafkaVersionFeatures) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(features, func() string { return "parameter features" }); err != nil {
		return err
	}

	return nil
}

