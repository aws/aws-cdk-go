package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configure how Cognito sends emails.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	email: cognito.userPoolEmail.withSES(&userPoolSESOptions{
//   		fromEmail: jsii.String("noreply@myawesomeapp.com"),
//   		fromName: jsii.String("Awesome App"),
//   		replyTo: jsii.String("support@myawesomeapp.com"),
//   	}),
//   })
//
// Experimental.
type UserPoolEmail interface {
}

// The jsii proxy struct for UserPoolEmail
type jsiiProxy_UserPoolEmail struct {
	_ byte // padding
}

// Experimental.
func NewUserPoolEmail_Override(u UserPoolEmail) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolEmail",
		nil, // no parameters
		u,
	)
}

// Send email using Cognito.
// Experimental.
func UserPoolEmail_WithCognito(replyTo *string) UserPoolEmail {
	_init_.Initialize()

	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolEmail",
		"withCognito",
		[]interface{}{replyTo},
		&returns,
	)

	return returns
}

// Send email using SES.
// Experimental.
func UserPoolEmail_WithSES(options *UserPoolSESOptions) UserPoolEmail {
	_init_.Initialize()

	if err := validateUserPoolEmail_WithSESParameters(options); err != nil {
		panic(err)
	}
	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolEmail",
		"withSES",
		[]interface{}{options},
		&returns,
	)

	return returns
}

