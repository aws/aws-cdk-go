package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configure how Cognito sends emails.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	Email: cognito.UserPoolEmail_WithSES(&UserPoolSESOptions{
//   		FromEmail: jsii.String("noreply@myawesomeapp.com"),
//   		FromName: jsii.String("Awesome App"),
//   		ReplyTo: jsii.String("support@myawesomeapp.com"),
//   	}),
//   })
//
type UserPoolEmail interface {
}

// The jsii proxy struct for UserPoolEmail
type jsiiProxy_UserPoolEmail struct {
	_ byte // padding
}

func NewUserPoolEmail_Override(u UserPoolEmail) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPoolEmail",
		nil, // no parameters
		u,
	)
}

// Send email using Cognito.
func UserPoolEmail_WithCognito(replyTo *string) UserPoolEmail {
	_init_.Initialize()

	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolEmail",
		"withCognito",
		[]interface{}{replyTo},
		&returns,
	)

	return returns
}

// Send email using SES.
func UserPoolEmail_WithSES(options *UserPoolSESOptions) UserPoolEmail {
	_init_.Initialize()

	if err := validateUserPoolEmail_WithSESParameters(options); err != nil {
		panic(err)
	}
	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolEmail",
		"withSES",
		[]interface{}{options},
		&returns,
	)

	return returns
}

