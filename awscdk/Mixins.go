package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Main entry point for applying mixins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   mixins := cdk.NewMixins()
//
type Mixins interface {
}

// The jsii proxy struct for Mixins
type jsiiProxy_Mixins struct {
	_ byte // padding
}

func NewMixins() Mixins {
	_init_.Initialize()

	j := jsiiProxy_Mixins{}

	_jsii_.Create(
		"aws-cdk-lib.Mixins",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewMixins_Override(m Mixins) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Mixins",
		nil, // no parameters
		m,
	)
}

// Creates a MixinApplicator for the given scope.
func Mixins_Of(scope constructs.IConstruct, selector IConstructSelector) MixinApplicator {
	_init_.Initialize()

	if err := validateMixins_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns MixinApplicator

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Mixins",
		"of",
		[]interface{}{scope, selector},
		&returns,
	)

	return returns
}

