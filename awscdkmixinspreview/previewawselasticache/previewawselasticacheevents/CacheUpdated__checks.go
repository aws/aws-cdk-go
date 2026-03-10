//go:build !no_runtime_type_checking

package previewawselasticacheevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCacheUpdated_CacheUpdatedPatternParameters(options *CacheUpdated_CacheUpdatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

