// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
//   // Read the secret from Secrets Manager
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("GitHub_Source"),
//   	owner: jsii.String("awslabs"),
//   	repo: jsii.String("aws-cdk"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	output: sourceOutput,
//   	branch: jsii.String("develop"),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
// Experimental.
type SecretValue interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	// Experimental.
	CreationStack() *[]*string
	// Creates a throwable Error object that contains the token creation stack trace.
	// Experimental.
	NewError(message *string) interface{}
	// Resolve the secret.
	//
	// If the feature flag is not set, resolve as normal. Otherwise, throw a descriptive
	// error that the usage guard is missing.
	// Experimental.
	Resolve(context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	// Experimental.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	// Experimental.
	ToString() *string
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
	// Experimental.
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


// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
// Experimental.
func NewSecretValue(protectedValue interface{}, options *IntrinsicProps) SecretValue {
	_init_.Initialize()

	if err := validateNewSecretValueParameters(protectedValue, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretValue{}

	_jsii_.Create(
		"monocdk.SecretValue",
		[]interface{}{protectedValue, options},
		&j,
	)

	return &j
}

// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
// Experimental.
func NewSecretValue_Override(s SecretValue, protectedValue interface{}, options *IntrinsicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.SecretValue",
		[]interface{}{protectedValue, options},
		s,
	)
}

// Obtain the secret value through a CloudFormation dynamic reference.
//
// If possible, use `SecretValue.ssmSecure` or `SecretValue.secretsManager` directly.
// Experimental.
func SecretValue_CfnDynamicReference(ref CfnDynamicReference) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_CfnDynamicReferenceParameters(ref); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
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
// Experimental.
func SecretValue_CfnParameter(param CfnParameter) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_CfnParameterParameters(param); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
		"cfnParameter",
		[]interface{}{param},
		&returns,
	)

	return returns
}

// Test whether an object is a SecretValue.
// Experimental.
func SecretValue_IsSecretValue(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretValue_IsSecretValueParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
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
		"monocdk.SecretValue",
		"plainText",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

// Use a resource's output as secret value.
// Experimental.
func SecretValue_ResourceAttribute(attr *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_ResourceAttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
		"resourceAttribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Creates a `SecretValue` with a value which is dynamically loaded from AWS Secrets Manager.
// Experimental.
func SecretValue_SecretsManager(secretId *string, options *SecretsManagerSecretOptions) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_SecretsManagerParameters(secretId, options); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
		"secretsManager",
		[]interface{}{secretId, options},
		&returns,
	)

	return returns
}

// Use a secret value stored from a Systems Manager (SSM) parameter.
// Experimental.
func SecretValue_SsmSecure(parameterName *string, version *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_SsmSecureParameters(parameterName); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
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
// The only reasonable use case for using this method is when you are testing.
// Experimental.
func SecretValue_UnsafePlainText(secret *string) SecretValue {
	_init_.Initialize()

	if err := validateSecretValue_UnsafePlainTextParameters(secret); err != nil {
		panic(err)
	}
	var returns SecretValue

	_jsii_.StaticInvoke(
		"monocdk.SecretValue",
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

