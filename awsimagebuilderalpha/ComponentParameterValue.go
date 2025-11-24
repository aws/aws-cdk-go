package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The parameter value for a component parameter.
//
// Example:
//   parameterizedComponent := imagebuilder.Component_FromComponentName(this, jsii.String("ParameterizedComponent"), jsii.String("my-parameterized-component"))
//
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ParameterizedImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	Components: []ComponentConfiguration{
//   		&ComponentConfiguration{
//   			Component: parameterizedComponent,
//   			Parameters: map[string]ComponentParameterValue{
//   				"environment": imagebuilder.ComponentParameterValue_fromString(jsii.String("production")),
//   				"version": imagebuilder.ComponentParameterValue_fromString(jsii.String("1.0.0")),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ComponentParameterValue interface {
	// The rendered parameter value.
	// Experimental.
	Value() *[]*string
}

// The jsii proxy struct for ComponentParameterValue
type jsiiProxy_ComponentParameterValue struct {
	_ byte // padding
}

func (j *jsiiProxy_ComponentParameterValue) Value() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponentParameterValue(value *[]*string) ComponentParameterValue {
	_init_.Initialize()

	if err := validateNewComponentParameterValueParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComponentParameterValue{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentParameterValue",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewComponentParameterValue_Override(c ComponentParameterValue, value *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentParameterValue",
		[]interface{}{value},
		c,
	)
}

// The value of the parameter as a string.
// Experimental.
func ComponentParameterValue_FromString(value *string) ComponentParameterValue {
	_init_.Initialize()

	if err := validateComponentParameterValue_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ComponentParameterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentParameterValue",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

