package core

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Applies mixins to constructs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var constructSelector ConstructSelector
//
//   mixinApplicator := awscdkmixinspreview.Core.NewMixinApplicator(this, constructSelector)
//
// Experimental.
type MixinApplicator interface {
	// Applies a mixin to selected constructs.
	// Experimental.
	Apply(mixins ...IMixin) MixinApplicator
	// Applies a mixin and requires that it be applied to all constructs.
	// Experimental.
	MustApply(mixins ...IMixin) MixinApplicator
}

// The jsii proxy struct for MixinApplicator
type jsiiProxy_MixinApplicator struct {
	_ byte // padding
}

// Experimental.
func NewMixinApplicator(scope constructs.IConstruct, selector ConstructSelector) MixinApplicator {
	_init_.Initialize()

	if err := validateNewMixinApplicatorParameters(scope); err != nil {
		panic(err)
	}
	j := jsiiProxy_MixinApplicator{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.MixinApplicator",
		[]interface{}{scope, selector},
		&j,
	)

	return &j
}

// Experimental.
func NewMixinApplicator_Override(m MixinApplicator, scope constructs.IConstruct, selector ConstructSelector) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.MixinApplicator",
		[]interface{}{scope, selector},
		m,
	)
}

func (m *jsiiProxy_MixinApplicator) Apply(mixins ...IMixin) MixinApplicator {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns MixinApplicator

	_jsii_.Invoke(
		m,
		"apply",
		args,
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MixinApplicator) MustApply(mixins ...IMixin) MixinApplicator {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns MixinApplicator

	_jsii_.Invoke(
		m,
		"mustApply",
		args,
		&returns,
	)

	return returns
}

