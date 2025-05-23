package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Key Signing Key for a Route 53 Hosted Zone.
//
// Example:
//   var hostedZone hostedZone
//   var kmsKey key
//
//   route53.NewKeySigningKey(this, jsii.String("KeySigningKey"), &KeySigningKeyProps{
//   	HostedZone: HostedZone,
//   	KmsKey: KmsKey,
//   	KeySigningKeyName: jsii.String("ksk"),
//   	Status: route53.KeySigningKeyStatus_ACTIVE,
//   })
//
type KeySigningKey interface {
	awscdk.Resource
	IKeySigningKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The hosted zone that the key signing key signs.
	HostedZone() IHostedZone
	// The ID of the key signing key, derived from the hosted zone ID and its name.
	KeySigningKeyId() *string
	// The name of the key signing key.
	KeySigningKeyName() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for KeySigningKey
type jsiiProxy_KeySigningKey struct {
	internal.Type__awscdkResource
	jsiiProxy_IKeySigningKey
}

func (j *jsiiProxy_KeySigningKey) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) HostedZone() IHostedZone {
	var returns IHostedZone
	_jsii_.Get(
		j,
		"hostedZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) KeySigningKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySigningKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) KeySigningKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySigningKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeySigningKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewKeySigningKey(scope constructs.Construct, id *string, props *KeySigningKeyProps) KeySigningKey {
	_init_.Initialize()

	if err := validateNewKeySigningKeyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeySigningKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKeySigningKey_Override(k KeySigningKey, scope constructs.Construct, id *string, props *KeySigningKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		[]interface{}{scope, id, props},
		k,
	)
}

// Imports a key signing key from its attributes.
func KeySigningKey_FromKeySigningKeyAttributes(scope constructs.Construct, id *string, attrs *KeySigningKeyAttributes) IKeySigningKey {
	_init_.Initialize()

	if err := validateKeySigningKey_FromKeySigningKeyAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IKeySigningKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		"fromKeySigningKeyAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
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
func KeySigningKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeySigningKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func KeySigningKey_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKeySigningKey_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func KeySigningKey_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKeySigningKey_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func KeySigningKey_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.KeySigningKey",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeySigningKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := k.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (k *jsiiProxy_KeySigningKey) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeySigningKey) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := k.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeySigningKey) GetResourceNameAttribute(nameAttr *string) *string {
	if err := k.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeySigningKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

