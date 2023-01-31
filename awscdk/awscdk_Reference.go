// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An intrinsic Token that represents a reference to a construct.
//
// References are recorded.
// Experimental.
type Reference interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DisplayName() *string
	// Experimental.
	Target() IConstruct
	// Creates a throwable Error object that contains the token creation stack trace.
	// Experimental.
	NewError(message *string) interface{}
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	// Experimental.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	// Experimental.
	ToString() *string
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

func (j *jsiiProxy_Reference) Target() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


// Experimental.
func NewReference_Override(r Reference, value interface{}, target IConstruct, displayName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.Reference",
		[]interface{}{value, target, displayName},
		r,
	)
}

// Check whether this is actually a Reference.
// Experimental.
func Reference_IsReference(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReference_IsReferenceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.Reference",
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

