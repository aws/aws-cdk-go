package mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Strategy for handling nested properties in L1 property mixins.
// Experimental.
type PropertyMergeStrategy interface {
}

// The jsii proxy struct for PropertyMergeStrategy
type jsiiProxy_PropertyMergeStrategy struct {
	_ byte // padding
}

// Deep merges nested objects from source into target.
//
// When both the existing and new value for a key are plain objects,
// their properties are merged recursively. Primitives, arrays, and
// mismatched types are overridden by the source value.
// Experimental.
func PropertyMergeStrategy_Combine() awscdk.IMergeStrategy {
	_init_.Initialize()

	var returns awscdk.IMergeStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.mixins.PropertyMergeStrategy",
		"combine",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Replaces existing property values on the target with the values from the source.
//
// Each allowed key is copied from source to target as-is, without
// inspecting nested objects. Any previous value on the target is discarded.
// Experimental.
func PropertyMergeStrategy_Override() awscdk.IMergeStrategy {
	_init_.Initialize()

	var returns awscdk.IMergeStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.mixins.PropertyMergeStrategy",
		"override",
		nil, // no parameters
		&returns,
	)

	return returns
}

