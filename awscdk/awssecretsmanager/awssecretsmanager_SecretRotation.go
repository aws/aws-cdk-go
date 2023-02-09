package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
// Experimental.
type SecretRotation interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for SecretRotation
type jsiiProxy_SecretRotation struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_SecretRotation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecretRotation(scope constructs.Construct, id *string, props *SecretRotationProps) SecretRotation {
	_init_.Initialize()

	if err := validateNewSecretRotationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretRotation{}

	_jsii_.Create(
		"monocdk.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecretRotation_Override(s SecretRotation, scope constructs.Construct, id *string, props *SecretRotationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func SecretRotation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretRotation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.SecretRotation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretRotation) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretRotation) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SecretRotation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretRotation) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretRotation) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

func (s *jsiiProxy_SecretRotation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

