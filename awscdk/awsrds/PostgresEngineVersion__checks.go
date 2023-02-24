//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validatePostgresEngineVersion_OfParameters(postgresFullVersion *string, postgresMajorVersion *string, postgresFeatures *PostgresEngineFeatures) error {
	if postgresFullVersion == nil {
		return fmt.Errorf("parameter postgresFullVersion is required, but nil was provided")
	}

	if postgresMajorVersion == nil {
		return fmt.Errorf("parameter postgresMajorVersion is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(postgresFeatures, func() string { return "parameter postgresFeatures" }); err != nil {
		return err
	}

	return nil
}

