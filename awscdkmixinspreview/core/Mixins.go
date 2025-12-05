package core

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Main entry point for applying mixins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mixins := awscdkmixinspreview.Core.NewMixins()
//
// Experimental.
type Mixins interface {
}

// The jsii proxy struct for Mixins
type jsiiProxy_Mixins struct {
	_ byte // padding
}

// Experimental.
func NewMixins() Mixins {
	_init_.Initialize()

	j := jsiiProxy_Mixins{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.Mixins",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMixins_Override(m Mixins) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.Mixins",
		nil, // no parameters
		m,
	)
}

// Creates a MixinApplicator for the given scope.
// Experimental.
func Mixins_Of(scope constructs.IConstruct, selector IConstructSelector) MixinApplicator {
	_init_.Initialize()

	if err := validateMixins_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns MixinApplicator

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.Mixins",
		"of",
		[]interface{}{scope, selector},
		&returns,
	)

	return returns
}

