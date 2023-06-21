package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Boolean custom attribute type.
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
type BooleanAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for BooleanAttribute
type jsiiProxy_BooleanAttribute struct {
	jsiiProxy_ICustomAttribute
}

func NewBooleanAttribute(props *CustomAttributeProps) BooleanAttribute {
	_init_.Initialize()

	if err := validateNewBooleanAttributeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BooleanAttribute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.BooleanAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBooleanAttribute_Override(b BooleanAttribute, props *CustomAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.BooleanAttribute",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_BooleanAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		b,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

