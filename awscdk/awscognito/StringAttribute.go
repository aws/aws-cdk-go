package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The String custom attribute type.
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
type StringAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for StringAttribute
type jsiiProxy_StringAttribute struct {
	jsiiProxy_ICustomAttribute
}

func NewStringAttribute(props *StringAttributeProps) StringAttribute {
	_init_.Initialize()

	if err := validateNewStringAttributeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringAttribute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.StringAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStringAttribute_Override(s StringAttribute, props *StringAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.StringAttribute",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StringAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		s,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

