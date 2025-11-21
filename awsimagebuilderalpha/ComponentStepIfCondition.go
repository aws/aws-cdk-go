package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an `if` condition in the component document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   var ifObject interface{}
//
//   componentStepIfCondition := imagebuilder_alpha.ComponentStepIfCondition_FromObject(map[string]interface{}{
//   	"ifObjectKey": ifObject,
//   })
//
// Experimental.
type ComponentStepIfCondition interface {
	// The rendered input value.
	// Experimental.
	IfCondition() interface{}
}

// The jsii proxy struct for ComponentStepIfCondition
type jsiiProxy_ComponentStepIfCondition struct {
	_ byte // padding
}

func (j *jsiiProxy_ComponentStepIfCondition) IfCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifCondition",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponentStepIfCondition(ifCondition interface{}) ComponentStepIfCondition {
	_init_.Initialize()

	if err := validateNewComponentStepIfConditionParameters(ifCondition); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComponentStepIfCondition{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepIfCondition",
		[]interface{}{ifCondition},
		&j,
	)

	return &j
}

// Experimental.
func NewComponentStepIfCondition_Override(c ComponentStepIfCondition, ifCondition interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepIfCondition",
		[]interface{}{ifCondition},
		c,
	)
}

// Creates the `if` value from an object, for the component step.
// Experimental.
func ComponentStepIfCondition_FromObject(ifObject *map[string]interface{}) ComponentStepIfCondition {
	_init_.Initialize()

	if err := validateComponentStepIfCondition_FromObjectParameters(ifObject); err != nil {
		panic(err)
	}
	var returns ComponentStepIfCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepIfCondition",
		"fromObject",
		[]interface{}{ifObject},
		&returns,
	)

	return returns
}

