package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the inputs for a step in the component document.
//
// Example:
//   phase := &ComponentDocumentPhase{
//   	Name: imagebuilder.ComponentPhaseName_BUILD,
//   	Steps: []ComponentDocumentStep{
//   		&ComponentDocumentStep{
//   			Name: jsii.String("configure-app"),
//   			Action: imagebuilder.ComponentAction_CREATE_FILE,
//   			Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
//   				"path": jsii.String("/etc/myapp/config.json"),
//   				"content": jsii.String("{\"env\": \"production\"}"),
//   			}),
//   		},
//   	},
//   }
//
// Experimental.
type ComponentStepInputs interface {
	// The rendered input value.
	// Experimental.
	Inputs() interface{}
}

// The jsii proxy struct for ComponentStepInputs
type jsiiProxy_ComponentStepInputs struct {
	_ byte // padding
}

func (j *jsiiProxy_ComponentStepInputs) Inputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponentStepInputs(input interface{}) ComponentStepInputs {
	_init_.Initialize()

	if err := validateNewComponentStepInputsParameters(input); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComponentStepInputs{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepInputs",
		[]interface{}{input},
		&j,
	)

	return &j
}

// Experimental.
func NewComponentStepInputs_Override(c ComponentStepInputs, input interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepInputs",
		[]interface{}{input},
		c,
	)
}

// Creates the input value from a list of input objects, for the component step.
// Experimental.
func ComponentStepInputs_FromList(inputsObjectList *[]*map[string]interface{}) ComponentStepInputs {
	_init_.Initialize()

	if err := validateComponentStepInputs_FromListParameters(inputsObjectList); err != nil {
		panic(err)
	}
	var returns ComponentStepInputs

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepInputs",
		"fromList",
		[]interface{}{inputsObjectList},
		&returns,
	)

	return returns
}

// Creates the input value from an object, for the component step.
// Experimental.
func ComponentStepInputs_FromObject(inputsObject *map[string]interface{}) ComponentStepInputs {
	_init_.Initialize()

	if err := validateComponentStepInputs_FromObjectParameters(inputsObject); err != nil {
		panic(err)
	}
	var returns ComponentStepInputs

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepInputs",
		"fromObject",
		[]interface{}{inputsObject},
		&returns,
	)

	return returns
}

