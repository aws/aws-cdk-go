package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Base class for the model side of context providers.
//
// Instances of this class communicate with context provider plugins in the 'cdk
// toolkit' via context variables (input), outputting specialized queries for
// more context variables (output).
//
// ContextProvider needs access to a Construct to hook into the context mechanism.
// Experimental.
type ContextProvider interface {
}

// The jsii proxy struct for ContextProvider
type jsiiProxy_ContextProvider struct {
	_ byte // padding
}

// Returns: the context key or undefined if a key cannot be rendered (due to tokens used in any of the props).
// Experimental.
func ContextProvider_GetKey(scope constructs.Construct, options *GetContextKeyOptions) *GetContextKeyResult {
	_init_.Initialize()

	if err := validateContextProvider_GetKeyParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *GetContextKeyResult

	_jsii_.StaticInvoke(
		"monocdk.ContextProvider",
		"getKey",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Experimental.
func ContextProvider_GetValue(scope constructs.Construct, options *GetContextValueOptions) *GetContextValueResult {
	_init_.Initialize()

	if err := validateContextProvider_GetValueParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *GetContextValueResult

	_jsii_.StaticInvoke(
		"monocdk.ContextProvider",
		"getValue",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

