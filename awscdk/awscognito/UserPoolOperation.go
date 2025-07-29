package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// User pool operations to which lambda triggers can be attached.
//
// Example:
//   authChallengeFn := lambda.NewFunction(this, jsii.String("authChallengeFn"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   })
//
//   userpool := cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	LambdaTriggers: &UserPoolTriggers{
//   		CreateAuthChallenge: authChallengeFn,
//   	},
//   })
//
//   userpool.AddTrigger(cognito.UserPoolOperation_USER_MIGRATION(), lambda.NewFunction(this, jsii.String("userMigrationFn"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_*FromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   }))
//
type UserPoolOperation interface {
	// The key to use in `CfnUserPool.LambdaConfigProperty`.
	OperationName() *string
}

// The jsii proxy struct for UserPoolOperation
type jsiiProxy_UserPoolOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_UserPoolOperation) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}


// A custom user pool operation.
func UserPoolOperation_Of(name *string) UserPoolOperation {
	_init_.Initialize()

	if err := validateUserPoolOperation_OfParameters(name); err != nil {
		panic(err)
	}
	var returns UserPoolOperation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func UserPoolOperation_CREATE_AUTH_CHALLENGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"CREATE_AUTH_CHALLENGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_EMAIL_SENDER() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"CUSTOM_EMAIL_SENDER",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_MESSAGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"CUSTOM_MESSAGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_SMS_SENDER() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"CUSTOM_SMS_SENDER",
		&returns,
	)
	return returns
}

func UserPoolOperation_DEFINE_AUTH_CHALLENGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"DEFINE_AUTH_CHALLENGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_POST_AUTHENTICATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"POST_AUTHENTICATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_POST_CONFIRMATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"POST_CONFIRMATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_AUTHENTICATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"PRE_AUTHENTICATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_SIGN_UP() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"PRE_SIGN_UP",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_TOKEN_GENERATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"PRE_TOKEN_GENERATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_TOKEN_GENERATION_CONFIG() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"PRE_TOKEN_GENERATION_CONFIG",
		&returns,
	)
	return returns
}

func UserPoolOperation_USER_MIGRATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"USER_MIGRATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_VERIFY_AUTH_CHALLENGE_RESPONSE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolOperation",
		"VERIFY_AUTH_CHALLENGE_RESPONSE",
		&returns,
	)
	return returns
}

