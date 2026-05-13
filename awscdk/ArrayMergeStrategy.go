package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Strategies for merging arrays in L1 property mixins.
//
// Array elements are never deep-merged.
type ArrayMergeStrategy interface {
}

// The jsii proxy struct for ArrayMergeStrategy
type jsiiProxy_ArrayMergeStrategy struct {
	_ byte // padding
}

// Append source elements after the existing target elements.
//
// Example:
//
//
func ArrayMergeStrategy_Append() IArrayMergeStrategy {
	_init_.Initialize()

	var returns IArrayMergeStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ArrayMergeStrategy",
		"append",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Prepend source elements before the existing target elements.
//
// Example:
//
//
func ArrayMergeStrategy_Prepend() IArrayMergeStrategy {
	_init_.Initialize()

	var returns IArrayMergeStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ArrayMergeStrategy",
		"prepend",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Replace the target array entirely with the source array.
//
// Example:
//
//
func ArrayMergeStrategy_Replace() IArrayMergeStrategy {
	_init_.Initialize()

	var returns IArrayMergeStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ArrayMergeStrategy",
		"replace",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overwrite target elements positionally with source elements.
//
// Target elements beyond the source length are preserved.
//
// Example:
//
//
func ArrayMergeStrategy_ReplaceByIndex() IArrayMergeStrategy {
	_init_.Initialize()

	var returns IArrayMergeStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ArrayMergeStrategy",
		"replaceByIndex",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Match source and target elements by a shared key property.
//
// Matching target elements are replaced (not deep-merged).
// Unmatched source elements are appended.
//
// Example:
//
//
func ArrayMergeStrategy_ReplaceByKey(key *string) IArrayMergeStrategy {
	_init_.Initialize()

	if err := validateArrayMergeStrategy_ReplaceByKeyParameters(key); err != nil {
		panic(err)
	}
	var returns IArrayMergeStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ArrayMergeStrategy",
		"replaceByKey",
		[]interface{}{key},
		&returns,
	)

	return returns
}

