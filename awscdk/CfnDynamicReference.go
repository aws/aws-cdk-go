package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// References a dynamically retrieved value.
//
// This is a Construct so that subclasses will (eventually) be able to attach
// metadata to themselves without having to change call signatures.
//
// Example:
//   awscdk.NewCfnDynamicReference(awscdk.CfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html
//
type CfnDynamicReference interface {
	Intrinsic
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

// The jsii proxy struct for CfnDynamicReference
type jsiiProxy_CfnDynamicReference struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_CfnDynamicReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDynamicReference) TypeHint() ResolutionTypeHint {
	var returns ResolutionTypeHint
	_jsii_.Get(
		j,
		"typeHint",
		&returns,
	)
	return returns
}


func NewCfnDynamicReference(service CfnDynamicReferenceService, key *string) CfnDynamicReference {
	_init_.Initialize()

	if err := validateNewCfnDynamicReferenceParameters(service, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDynamicReference{}

	_jsii_.Create(
		"aws-cdk-lib.CfnDynamicReference",
		[]interface{}{service, key},
		&j,
	)

	return &j
}

func NewCfnDynamicReference_Override(c CfnDynamicReference, service CfnDynamicReferenceService, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnDynamicReference",
		[]interface{}{service, key},
		c,
	)
}

func (c *jsiiProxy_CfnDynamicReference) NewError(message *string) interface{} {
	if err := c.validateNewErrorParameters(message); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) Resolve(_context IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) ToStringList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"toStringList",
		nil, // no parameters
		&returns,
	)

	return returns
}

