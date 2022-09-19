package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Number custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
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
type NumberAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for NumberAttribute
type jsiiProxy_NumberAttribute struct {
	jsiiProxy_ICustomAttribute
}

func NewNumberAttribute(props *NumberAttributeProps) NumberAttribute {
	_init_.Initialize()

	if err := validateNewNumberAttributeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NumberAttribute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.NumberAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNumberAttribute_Override(n NumberAttribute, props *NumberAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.NumberAttribute",
		[]interface{}{props},
		n,
	)
}

func (n *jsiiProxy_NumberAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		n,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

