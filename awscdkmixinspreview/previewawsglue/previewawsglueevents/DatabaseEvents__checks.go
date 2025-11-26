//go:build !no_runtime_type_checking

package previewawsglueevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

func (d *jsiiProxy_DatabaseEvents) validateGlueDataCatalogDatabaseStateChangePatternParameters(options *DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatabaseEvents) validateGlueDataCatalogTableStateChangePatternParameters(options *DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDatabaseEvents_FromDatabaseParameters(databaseRef interfacesawsglue.IDatabaseRef) error {
	if databaseRef == nil {
		return fmt.Errorf("parameter databaseRef is required, but nil was provided")
	}

	return nil
}

