package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines a parameter for an extension.
//
// Example:
//   var fn function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   	Parameters: []parameter{
//   		appconfig.*parameter_Required(jsii.String("testParam"), jsii.String("true")),
//   		appconfig.*parameter_NotRequired(jsii.String("testNotRequiredParam")),
//   	},
//   })
//
type Parameter interface {
	// The description of the parameter.
	Description() *string
	// A boolean that indicates if the parameter is required or optional.
	IsRequired() *bool
	// The name of the parameter.
	Name() *string
	// The value of the parameter.
	Value() *string
}

// The jsii proxy struct for Parameter
type jsiiProxy_Parameter struct {
	_ byte // padding
}

func (j *jsiiProxy_Parameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parameter) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parameter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parameter) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// An optional parameter for an extension.
func Parameter_NotRequired(name *string, value *string, description *string) Parameter {
	_init_.Initialize()

	if err := validateParameter_NotRequiredParameters(name); err != nil {
		panic(err)
	}
	var returns Parameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.Parameter",
		"notRequired",
		[]interface{}{name, value, description},
		&returns,
	)

	return returns
}

// A required parameter for an extension.
func Parameter_Required(name *string, value *string, description *string) Parameter {
	_init_.Initialize()

	if err := validateParameter_RequiredParameters(name, value); err != nil {
		panic(err)
	}
	var returns Parameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.Parameter",
		"required",
		[]interface{}{name, value, description},
		&returns,
	)

	return returns
}

