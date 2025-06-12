package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An object which serializes to the JSON `null` literal, and which can safely be passed across languages where `undefined` and `null` are not different.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonNull := cdk.JsonNull_INSTANCE()
//
type JsonNull interface {
	IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// Produce the Token's value at resolution time.
	Resolve(_ctx IResolveContext) interface{}
	// Obtains the JSON representation of this object (`null`).
	ToJSON() interface{}
	// Obtains the string representation of this object (`'null'`).
	ToString() *string
}

// The jsii proxy struct for JsonNull
type jsiiProxy_JsonNull struct {
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_JsonNull) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


func JsonNull_INSTANCE() JsonNull {
	_init_.Initialize()
	var returns JsonNull
	_jsii_.StaticGet(
		"aws-cdk-lib.JsonNull",
		"INSTANCE",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonNull) Resolve(_ctx IResolveContext) interface{} {
	if err := j.validateResolveParameters(_ctx); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_ctx},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonNull) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonNull) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

