package core

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract base class for mixins that provides default implementations.
// Experimental.
type Mixin interface {
	IMixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for Mixin
type jsiiProxy_Mixin struct {
	jsiiProxy_IMixin
}

// Experimental.
func NewMixin_Override(m Mixin) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.Mixin",
		nil, // no parameters
		m,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func Mixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.Mixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := m.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mixin) Supports(construct constructs.IConstruct) *bool {
	if err := m.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		m,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

