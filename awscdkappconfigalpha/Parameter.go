package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines a parameter for an extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   parameter := appconfig_alpha.Parameter_NotRequired(jsii.String("name"), jsii.String("value"), jsii.String("description"))
//
// Deprecated.
type Parameter interface {
	// The description of the parameter.
	// Deprecated.
	Description() *string
	// A boolean that indicates if the parameter is required or optional.
	// Deprecated.
	IsRequired() *bool
	// The name of the parameter.
	// Deprecated.
	Name() *string
	// The value of the parameter.
	// Deprecated.
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
// Deprecated.
func Parameter_NotRequired(name *string, value *string, description *string) Parameter {
	_init_.Initialize()

	if err := validateParameter_NotRequiredParameters(name); err != nil {
		panic(err)
	}
	var returns Parameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.Parameter",
		"notRequired",
		[]interface{}{name, value, description},
		&returns,
	)

	return returns
}

// A required parameter for an extension.
// Deprecated.
func Parameter_Required(name *string, value *string, description *string) Parameter {
	_init_.Initialize()

	if err := validateParameter_RequiredParameters(name, value); err != nil {
		panic(err)
	}
	var returns Parameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.Parameter",
		"required",
		[]interface{}{name, value, description},
		&returns,
	)

	return returns
}

