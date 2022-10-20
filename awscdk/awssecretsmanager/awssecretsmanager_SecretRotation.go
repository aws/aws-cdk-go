package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Secret rotation for a service or database.
//
// Example:
//   var mySecret secret
//   var myDatabase iConnectable
//   var myVpc vpc
//
//
//   secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &secretRotationProps{
//   	application: secretsmanager.secretRotationApplication_MYSQL_ROTATION_SINGLE_USER(),
//   	 // MySQL single user scheme
//   	secret: mySecret,
//   	target: myDatabase,
//   	 // a Connectable
//   	vpc: myVpc,
//   	 // The VPC where the secret rotation application will be deployed
//   	excludeCharacters: jsii.String(" %+:;{}"),
//   })
//
type SecretRotation interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SecretRotation
type jsiiProxy_SecretRotation struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SecretRotation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSecretRotation(scope constructs.Construct, id *string, props *SecretRotationProps) SecretRotation {
	_init_.Initialize()

	if err := validateNewSecretRotationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretRotation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecretRotation_Override(s SecretRotation, scope constructs.Construct, id *string, props *SecretRotationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SecretRotation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretRotation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretRotation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

