package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Token subclass that represents values intrinsic to the target document language.
//
// WARNING: this class should not be externally exposed, but is currently visible
// because of a limitation of jsii (https://github.com/aws/jsii/issues/524).
//
// This class will disappear in a future release and should not be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   intrinsic := cdk.NewIntrinsic(value, &IntrinsicProps{
//   	StackTrace: jsii.Boolean(false),
//   	TypeHint: cdk.ResolutionTypeHint_STRING,
//   })
//
type Intrinsic interface {
	IResolvable
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
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

// The jsii proxy struct for Intrinsic
type jsiiProxy_Intrinsic struct {
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_Intrinsic) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Intrinsic) TypeHint() ResolutionTypeHint {
	var returns ResolutionTypeHint
	_jsii_.Get(
		j,
		"typeHint",
		&returns,
	)
	return returns
}


func NewIntrinsic(value interface{}, options *IntrinsicProps) Intrinsic {
	_init_.Initialize()

	if err := validateNewIntrinsicParameters(value, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Intrinsic{}

	_jsii_.Create(
		"aws-cdk-lib.Intrinsic",
		[]interface{}{value, options},
		&j,
	)

	return &j
}

func NewIntrinsic_Override(i Intrinsic, value interface{}, options *IntrinsicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Intrinsic",
		[]interface{}{value, options},
		i,
	)
}

func (i *jsiiProxy_Intrinsic) NewError(message *string) interface{} {
	if err := i.validateNewErrorParameters(message); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) Resolve(_context IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) ToStringList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"toStringList",
		nil, // no parameters
		&returns,
	)

	return returns
}

