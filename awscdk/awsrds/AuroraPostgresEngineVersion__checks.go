//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAuroraPostgresEngineVersion_OfParameters(auroraPostgresFullVersion *string, auroraPostgresMajorVersion *string, auroraPostgresFeatures *AuroraPostgresEngineFeatures) error {
	if auroraPostgresFullVersion == nil {
		return fmt.Errorf("parameter auroraPostgresFullVersion is required, but nil was provided")
	}

	if auroraPostgresMajorVersion == nil {
		return fmt.Errorf("parameter auroraPostgresMajorVersion is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(auroraPostgresFeatures, func() string { return "parameter auroraPostgresFeatures" }); err != nil {
		return err
	}

	return nil
}

