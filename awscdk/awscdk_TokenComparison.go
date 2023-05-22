// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An enum-like class that represents the result of comparing two Tokens.
//
// The return type of {@link Token.compareStrings}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tokenComparison := monocdk.TokenComparison_BOTH_UNRESOLVED()
//
// Experimental.
type TokenComparison interface {
}

// The jsii proxy struct for TokenComparison
type jsiiProxy_TokenComparison struct {
	_ byte // padding
}

func TokenComparison_BOTH_UNRESOLVED() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"monocdk.TokenComparison",
		"BOTH_UNRESOLVED",
		&returns,
	)
	return returns
}

func TokenComparison_DIFFERENT() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"monocdk.TokenComparison",
		"DIFFERENT",
		&returns,
	)
	return returns
}

func TokenComparison_ONE_UNRESOLVED() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"monocdk.TokenComparison",
		"ONE_UNRESOLVED",
		&returns,
	)
	return returns
}

func TokenComparison_SAME() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"monocdk.TokenComparison",
		"SAME",
		&returns,
	)
	return returns
}

