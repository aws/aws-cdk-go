package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The array that the Map state will iterate over.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var array interface{}
//
//   provideItems := awscdk.Aws_stepfunctions.ProvideItems_JsonArray([]interface{}{
//   	array,
//   })
//
type ProvideItems interface {
	// The array that the Map state will iterate over.
	Items() interface{}
}

// The jsii proxy struct for ProvideItems
type jsiiProxy_ProvideItems struct {
	_ byte // padding
}

func (j *jsiiProxy_ProvideItems) Items() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}


func NewProvideItems_Override(p ProvideItems) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ProvideItems",
		nil, // no parameters
		p,
	)
}

// Use a JSON array as Map state items.
//
// Example value: `[1, "{% $two %}", 3]`.
func ProvideItems_JsonArray(array *[]interface{}) ProvideItems {
	_init_.Initialize()

	if err := validateProvideItems_JsonArrayParameters(array); err != nil {
		panic(err)
	}
	var returns ProvideItems

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.ProvideItems",
		"jsonArray",
		[]interface{}{array},
		&returns,
	)

	return returns
}

// Use a JSONata expression as Map state items.
//
// Example value: `{% $states.input.items %}`
func ProvideItems_Jsonata(jsonataExpression *string) ProvideItems {
	_init_.Initialize()

	if err := validateProvideItems_JsonataParameters(jsonataExpression); err != nil {
		panic(err)
	}
	var returns ProvideItems

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.ProvideItems",
		"jsonata",
		[]interface{}{jsonataExpression},
		&returns,
	)

	return returns
}

