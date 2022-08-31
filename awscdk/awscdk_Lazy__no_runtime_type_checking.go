//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateLazy_AnyParameters(producer IStableAnyProducer, options *LazyAnyValueOptions) error {
	return nil
}

func validateLazy_AnyValueParameters(producer IAnyProducer, options *LazyAnyValueOptions) error {
	return nil
}

func validateLazy_ListParameters(producer IStableListProducer, options *LazyListValueOptions) error {
	return nil
}

func validateLazy_ListValueParameters(producer IListProducer, options *LazyListValueOptions) error {
	return nil
}

func validateLazy_NumberParameters(producer IStableNumberProducer) error {
	return nil
}

func validateLazy_NumberValueParameters(producer INumberProducer) error {
	return nil
}

func validateLazy_StringParameters(producer IStableStringProducer, options *LazyStringValueOptions) error {
	return nil
}

func validateLazy_StringValueParameters(producer IStringProducer, options *LazyStringValueOptions) error {
	return nil
}

func validateLazy_UncachedAnyParameters(producer IAnyProducer, options *LazyAnyValueOptions) error {
	return nil
}

func validateLazy_UncachedListParameters(producer IListProducer, options *LazyListValueOptions) error {
	return nil
}

func validateLazy_UncachedNumberParameters(producer INumberProducer) error {
	return nil
}

func validateLazy_UncachedStringParameters(producer IStringProducer, options *LazyStringValueOptions) error {
	return nil
}

