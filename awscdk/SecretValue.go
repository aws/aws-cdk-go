package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Work with secret values in the CDK.
//
// Constructs that need secrets will declare parameters of type `SecretValue`.
//
// The actual values of these secrets should not be committed to your
// repository, or even end up in the synthesized CloudFormation template. Instead, you should
// store them in an external system like AWS Secrets Manager or SSM Parameter
// Store, and you can reference them by calling `SecretValue.secretsManager()` or
// `SecretValue.ssmSecure()`.
//
// You can use `SecretValue.unsafePlainText()` to construct a `SecretValue` from a
// literal string, but doing so is highly discouraged.
//
// To make sure secret values don't accidentally end up in readable parts
// of your infrastructure definition (such as the environment variables
// of an AWS Lambda Function, where everyone who can read the function
// definition has access to the secret), using secret values directly is not
// allowed. You must pass them to constructs that accept `SecretValue`
// properties, which are guaranteed to use the value only in CloudFormation
// properties that are write-only.
//
// If you are sure that what you are doing is safe, you can call
// `secretValue.unsafeUnwrap()` to access the protected string of the secret
// value.
//
// (If you are writing something like an AWS Lambda Function and need to access
// a secret inside it, make the API call to `GetSecretValue` directly inside
// your Lamba's code, instead of using environment variables.)
//
// Example:
//   var myHostedZone iPublicHostedZone
//
//
//   ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
//   	Identity: ses.Identity_PublicHostedZone(myHostedZone),
//   	DkimIdentity: ses.DkimIdentity_ByoDkim(&ByoDkimOptions{
//   		PrivateKey: awscdk.SecretValue_SecretsManager(jsii.String("dkim-private-key")),
//   		PublicKey: jsii.String("...base64-encoded-public-key..."),
//   		Selector: jsii.String("selector"),
//   	}),
//   })
//
type SecretValue interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	// Type that the Intrinsic is expected to evaluate to.
	TypeHint() ResolutionTypeHint
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Resolve the secret.
	//
	// If the feature flag is not set, resolve as normal. Otherwise, throw a descriptive
	// error that the usage guard is missing.
	Resolve(context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
	// Convert an instance of this Token to a string list.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a list. We treat it the same as an explicit
	// stringification.
	ToStringList() *[]*string
	// Disable usage protection on this secret.
	//
	// Call this to indicate that you want to use the secret value held by this
	// object in an unchecked way. If you don't call this method, using the secret
	// value directly in a string context or as a property value somewhere will
	// produce an error.
	//
	// This method has 'unsafe' in the name on purpose! Make sure that the
	// construct property you are using the returned value in is does not end up
	// in a place in your AWS infrastructure where it could be read by anyone
	// unexpected.
	//
	// When in doubt, don't call this method and only pass the object to constructs that
	// accept `SecretValue` parameters.
	UnsafeUnwrap() *string
}

// The jsii proxy struct for SecretValue
type jsiiProxy_SecretValue struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_SecretValue) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretValue) TypeHint() ResolutionTypeHint {
	var returns ResolutionTypeHint
	_jsii_.Get(
		j,
		"typeHint",
		&returns,
	)
	return returns
}


// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
func NewSecretValue(protectedValue interface{}, options *IntrinsicProps) SecretValue {
	_init_.Initialize()

	if err := validateNewSecretValueParameters(protectedValue, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretValue{}

	_jsii_.Create(
		"aws-cdk-lib.SecretValue",
		[]interface{}{protectedValue, options},
		&j,
	)

	return &j
}

// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
func NewSecretValue_Override(s SecretValue, protectedValue interface{}, options *IntrinsicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.SecretValue",
		[]interface{}{protectedValue, options},
		s,
	)
}

// Obtain the secret value through a CloudFormation dynamic reference.
//
// If possible, use `SecretValue.ssmSecure` or `SecretValue.secretsManager` directly.
func SecretValue_CfnDynamicReference(ref CfnDynamicReference) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_CfnDynamicReferenceParameters(ref); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"cfnDynamicReference",
		[]interface{}{ref},
		&returns,
	)

	return returns
}

// Obtain the secret value through a CloudFormation parameter.
//
// Generally, this is not a recommended approach. AWS Secrets Manager is the
// recommended way to reference secrets.
func SecretValue_CfnParameter(param CfnParameter) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_CfnParameterParameters(param); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"cfnParameter",
		[]interface{}{param},
		&returns,
	)

	return returns
}

// Test whether an object is a SecretValue.
func SecretValue_IsSecretValue(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretValue_IsSecretValueParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"isSecretValue",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Construct a literal secret value for use with secret-aware constructs.
//
// Do not use this method for any secrets that you care about! The value
// will be visible to anyone who has access to the CloudFormation template
// (via the AWS Console, SDKs, or CLI).
//
// The only reasonable use case for using this method is when you are testing.
// Deprecated: Use `unsafePlainText()` instead.
func SecretValue_PlainText(secret *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_PlainTextParameters(secret); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"plainText",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

// Use a resource's output as secret value.
func SecretValue_ResourceAttribute(attr *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_ResourceAttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"resourceAttribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Creates a `SecretValue` with a value which is dynamically loaded from AWS Secrets Manager.
//
// If you rotate the value in the Secret, you must also change at least one property
// on the resource where you are using the secret, to force CloudFormation to re-read the secret.
func SecretValue_SecretsManager(secretId *string, options *SecretsManagerSecretOptions) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_SecretsManagerParameters(secretId, options); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"secretsManager",
		[]interface{}{secretId, options},
		&returns,
	)

	return returns
}

// Use a secret value stored from a Systems Manager (SSM) parameter.
//
// This secret source in only supported in a limited set of resources and
// properties. [Click here for the list of supported
// properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#template-parameters-dynamic-patterns-resources).
func SecretValue_SsmSecure(parameterName *string, version *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_SsmSecureParameters(parameterName); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"ssmSecure",
		[]interface{}{parameterName, version},
		&returns,
	)

	return returns
}

// Construct a literal secret value for use with secret-aware constructs.
//
// Do not use this method for any secrets that you care about! The value
// will be visible to anyone who has access to the CloudFormation template
// (via the AWS Console, SDKs, or CLI).
//
// The primary use case for using this method is when you are testing.
//
// The other use case where this is appropriate is when constructing a JSON secret.
// For example, a JSON secret might have multiple fields where only some are actual
// secret values.
//
// Example:
//   var secret secretValue
//
//   jsonSecret := map[string]secretValue{
//   	"username": awscdk.SecretValue_unsafePlainText(jsii.String("myUsername")),
//   	"password": secret,
//   }
//
func SecretValue_UnsafePlainText(secret *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_UnsafePlainTextParameters(secret); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"unsafePlainText",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) NewError(message *string) interface{} {
	if err := s.validateNewErrorParameters(message); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) Resolve(context IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) ToStringList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"toStringList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) UnsafeUnwrap() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"unsafeUnwrap",
		nil, // no parameters
		&returns,
	)

	return returns
}

