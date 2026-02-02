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
//   var constructSelector IConstructSelector
//
//   mixinApplicator := awscdkmixinspreview.Core.NewMixinApplicator(this, constructSelector)
//
// Experimental.
type MixinApplicator interface {
	// Returns the successful mixin applications.
	// Experimental.
	Report() *[]*MixinApplication
	// The constructs that match the selector in the given scope.
	// Experimental.
	SelectedConstructs() *[]constructs.IConstruct
	// Applies a mixin to selected constructs.
	// Experimental.
	Apply(mixins ...IMixin) MixinApplicator
	// Requires all selected constructs to support the applied mixins.
	//
	// Will only check for future call of `apply()`.
	// Set this before calling `apply()` to take effect.
	//
	// Example:
	//   awscdkmixinspreview.Mixins_Of(scope).RequireAll().Apply(NewMyMixin())
	//
	// Experimental.
	RequireAll() MixinApplicator
	// Requires at least one mixin to be successfully applied.
	//
	// Will only check for future call of `apply()`.
	// Set this before calling `apply()` to take effect.
	//
	// Example:
	//   awscdkmixinspreview.Mixins_Of(scope).RequireAny().Apply(NewMyMixin())
	//
	// Experimental.
	RequireAny() MixinApplicator
}

// The jsii proxy struct for MixinApplicator
type jsiiProxy_MixinApplicator struct {
	_ byte // padding
}

func (j *jsiiProxy_MixinApplicator) Report() *[]*MixinApplication {
	var returns *[]*MixinApplication
	_jsii_.Get(
		j,
		"report",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MixinApplicator) SelectedConstructs() *[]constructs.IConstruct {
	var returns *[]constructs.IConstruct
	_jsii_.Get(
		j,
		"selectedConstructs",
		&returns,
	)
	return returns
}


// Experimental.
func NewMixinApplicator(scope constructs.IConstruct, selector IConstructSelector) MixinApplicator {
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
func NewMixinApplicator_Override(m MixinApplicator, scope constructs.IConstruct, selector IConstructSelector) {
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

func (m *jsiiProxy_MixinApplicator) RequireAll() MixinApplicator {
	var returns MixinApplicator

	_jsii_.Invoke(
		m,
		"requireAll",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MixinApplicator) RequireAny() MixinApplicator {
	var returns MixinApplicator

	_jsii_.Invoke(
		m,
		"requireAny",
		nil, // no parameters
		&returns,
	)

	return returns
}

