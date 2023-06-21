package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Lazily produce a value.
//
// Can be used to return a string, list or numeric value whose actual value
// will only be calculated later, during synthesis.
type Lazy interface {
}

// The jsii proxy struct for Lazy
type jsiiProxy_Lazy struct {
	_ byte // padding
}

// Defer the one-time calculation of an arbitrarily typed value to synthesis time.
//
// Use this if you want to render an object to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// The inner function will only be invoked one time and cannot depend on
// resolution context.
func Lazy_Any(producer IStableAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	if err := validateLazy_AnyParameters(producer, options); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"any",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a list value to synthesis time.
//
// Use this if you want to render a list to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `string[]` type and don't need
// the calculation to be deferred, use `Token.asList()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_List(producer IStableListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	if err := validateLazy_ListParameters(producer, options); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"list",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a number value to synthesis time.
//
// Use this if you want to render a number to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `number` type and don't need
// the calculation to be deferred, use `Token.asNumber()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_Number(producer IStableNumberProducer) *float64 {
	_init_.Initialize()

	if err := validateLazy_NumberParameters(producer); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"number",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a string value to synthesis time.
//
// Use this if you want to render a string to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `string` type and don't need
// the calculation to be deferred, use `Token.asString()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_String(producer IStableStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	if err := validateLazy_StringParameters(producer, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"string",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of an untyped value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.any()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedAny(producer IAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	if err := validateLazy_UncachedAnyParameters(producer, options); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedAny",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of a list value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.list()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedList(producer IListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	if err := validateLazy_UncachedListParameters(producer, options); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedList",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of a number value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.number()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedNumber(producer INumberProducer) *float64 {
	_init_.Initialize()

	if err := validateLazy_UncachedNumberParameters(producer); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedNumber",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Defer the calculation of a string value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.string()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedString(producer IStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	if err := validateLazy_UncachedStringParameters(producer, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedString",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

