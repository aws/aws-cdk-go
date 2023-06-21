package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An experimental class used to specify an initial secret value for a Secret.
//
// The class wraps a simple string (or JSON representation) in order to provide some safety checks and warnings
// about the dangers of using plaintext strings as initial secret seed values via CDK/CloudFormation.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   accessKey := iam.NewAccessKey(this, jsii.String("AccessKey"), &AccessKeyProps{
//   	User: User,
//   })
//   secretValue := secretsmanager.SecretStringValueBeta1_FromToken(jSON.stringify(map[string]interface{}{
//   	"username": user.userName,
//   	"database": jsii.String("foo"),
//   	"password": accessKey.secretAccessKey.unsafeUnwrap(),
//   }))
//
// Deprecated: Use `cdk.SecretValue` instead.
type SecretStringValueBeta1 interface {
	// Returns the secret value.
	// Deprecated: Use `cdk.SecretValue` instead.
	SecretValue() *string
}

// The jsii proxy struct for SecretStringValueBeta1
type jsiiProxy_SecretStringValueBeta1 struct {
	_ byte // padding
}

// Creates a `SecretValueValueBeta1` from a string value coming from a Token.
//
// The intent is to enable creating secrets from references (e.g., `Ref`, `Fn::GetAtt`) from other resources.
// This might be the direct output of another Construct, or the output of a Custom Resource.
// This method throws if it determines the input is an unsafe plaintext string.
//
// For example:
//
// ```ts
// // Creates a new IAM user, access and secret keys, and stores the secret access key in a Secret.
// const user = new iam.User(this, 'User');
// const accessKey = new iam.AccessKey(this, 'AccessKey', { user });
// const secret = new secretsmanager.Secret(this, 'Secret', {
// 	secretStringValue: accessKey.secretAccessKey,
// });
// ```
//
// The secret may also be embedded in a string representation of a JSON structure:
//
// ```ts
// const user = new iam.User(this, 'User');
// const accessKey = new iam.AccessKey(this, 'AccessKey', { user });
// const secretValue = secretsmanager.SecretStringValueBeta1.fromToken(JSON.stringify({
//   username: user.userName,
//   database: 'foo',
//   password: accessKey.secretAccessKey.unsafeUnwrap(),
// }));
// ```
//
// Note that the value being a Token does *not* guarantee safety. For example, a Lazy-evaluated string
// (e.g., `Lazy.string({ produce: () => 'myInsecurePassword' }))`) is a Token, but as the output is
// ultimately a plaintext string, and so insecure.
// Deprecated: Use `cdk.SecretValue` instead.
func SecretStringValueBeta1_FromToken(secretValueFromToken *string) SecretStringValueBeta1 {
	_init_.Initialize()

	if err := validateSecretStringValueBeta1_FromTokenParameters(secretValueFromToken); err != nil {
		panic(err)
	}
	var returns SecretStringValueBeta1

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1",
		"fromToken",
		[]interface{}{secretValueFromToken},
		&returns,
	)

	return returns
}

// Creates a `SecretStringValueBeta1` from a plaintext value.
//
// This approach is inherently unsafe, as the secret value may be visible in your source control repository
// and will also appear in plaintext in the resulting CloudFormation template, including in the AWS Console or APIs.
// Usage of this method is discouraged, especially for production workloads.
// Deprecated: Use `cdk.SecretValue` instead.
func SecretStringValueBeta1_FromUnsafePlaintext(secretValue *string) SecretStringValueBeta1 {
	_init_.Initialize()

	if err := validateSecretStringValueBeta1_FromUnsafePlaintextParameters(secretValue); err != nil {
		panic(err)
	}
	var returns SecretStringValueBeta1

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1",
		"fromUnsafePlaintext",
		[]interface{}{secretValue},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretStringValueBeta1) SecretValue() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"secretValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

