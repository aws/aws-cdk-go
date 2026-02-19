//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_MetadataExtractionProcessor) validateBindParameters(scope constructs.Construct, options *DataProcessorBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateMetadataExtractionProcessor_Jq16Parameters(query *map[string]*string) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	return nil
}

func validateNewMetadataExtractionProcessorParameters(options *MetadataExtractionProcessorOptions, keys *[]*string) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	if keys == nil {
		return fmt.Errorf("parameter keys is required, but nil was provided")
	}

	return nil
}

