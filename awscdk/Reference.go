package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// An intrinsic Token that represents a reference to a construct.
//
// References are recorded.
type Reference interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	DisplayName() *string
	Target() constructs.IConstruct
	// Type that the Intrinsic is expected to evaluate to.
	TypeHint() ResolutionTypeHint
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Produce the Token's value at resolution time.
	Resolve(_context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
	// Convert an instance of this Token to a string list.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a list. We treat it the same as an explicit
	// stringification.
	ToStringList() *[]*string
}

// The jsii proxy struct for Reference
type jsiiProxy_Reference struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_Reference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Reference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Reference) Target() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Reference) TypeHint() ResolutionTypeHint {
	var returns ResolutionTypeHint
	_jsii_.Get(
		j,
		"typeHint",
		&returns,
	)
	return returns
}


func NewReference_Override(r Reference, value interface{}, target constructs.IConstruct, displayName *string, typeHint ResolutionTypeHint) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Reference",
		[]interface{}{value, target, displayName, typeHint},
		r,
	)
}

// Check whether this is actually a Reference.
func Reference_IsReference(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReference_IsReferenceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Reference",
		"isReference",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) NewError(message *string) interface{} {
	if err := r.validateNewErrorParameters(message); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) Resolve(_context IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) ToStringList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"toStringList",
		nil, // no parameters
		&returns,
	)

	return returns
}

