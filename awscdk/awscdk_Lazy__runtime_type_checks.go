//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateLazy_AnyParameters(producer IStableAnyProducer, options *LazyAnyValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLazy_ListParameters(producer IStableListProducer, options *LazyListValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLazy_NumberParameters(producer IStableNumberProducer) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	return nil
}

func validateLazy_StringParameters(producer IStableStringProducer, options *LazyStringValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLazy_UncachedAnyParameters(producer IAnyProducer, options *LazyAnyValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLazy_UncachedListParameters(producer IListProducer, options *LazyListValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLazy_UncachedNumberParameters(producer INumberProducer) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	return nil
}

func validateLazy_UncachedStringParameters(producer IStringProducer, options *LazyStringValueOptions) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

