package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The DateTime custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	StandardAttributes: &StandardAttributes{
//   		Fullname: &StandardAttribute{
//   			Required: jsii.Boolean(true),
//   			Mutable: jsii.Boolean(false),
//   		},
//   		Address: &StandardAttribute{
//   			Required: jsii.Boolean(false),
//   			Mutable: jsii.Boolean(true),
//   		},
//   	},
//   	CustomAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
type DateTimeAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for DateTimeAttribute
type jsiiProxy_DateTimeAttribute struct {
	jsiiProxy_ICustomAttribute
}

func NewDateTimeAttribute(props *CustomAttributeProps) DateTimeAttribute {
	_init_.Initialize()

	if err := validateNewDateTimeAttributeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DateTimeAttribute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.DateTimeAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDateTimeAttribute_Override(d DateTimeAttribute, props *CustomAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.DateTimeAttribute",
		[]interface{}{props},
		d,
	)
}

func (d *jsiiProxy_DateTimeAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		d,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

