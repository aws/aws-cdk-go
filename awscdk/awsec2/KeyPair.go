package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

// An EC2 Key Pair.
//
// Example:
//   keyPair := ec2.KeyPair_FromKeyPairAttributes(this, jsii.String("KeyPair"), &KeyPairAttributes{
//   	KeyPairName: jsii.String("the-keypair-name"),
//   	Type: ec2.KeyPairType_RSA,
//   })
//
type KeyPair interface {
	awscdk.Resource
	IKeyPair
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The format of the key pair.
	Format() KeyPairFormat
	// Whether the key material was imported.
	//
	// Keys with imported material do not have their private key material stored
	// or returned automatically.
	HasImportedMaterial() *bool
	// The fingerprint of the key pair.
	KeyPairFingerprint() *string
	// The unique ID of the key pair.
	KeyPairId() *string
	// The unique name of the key pair.
	KeyPairName() *string
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
	// The Systems Manager Parameter Store parameter with the pair's private key material.
	PrivateKey() awsssm.IStringParameter
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The type of the key pair.
	Type() KeyPairType
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

// The jsii proxy struct for KeyPair
type jsiiProxy_KeyPair struct {
	internal.Type__awscdkResource
	jsiiProxy_IKeyPair
}

func (j *jsiiProxy_KeyPair) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) Format() KeyPairFormat {
	var returns KeyPairFormat
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) HasImportedMaterial() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasImportedMaterial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) KeyPairFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) KeyPairId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) KeyPairName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) PrivateKey() awsssm.IStringParameter {
	var returns awsssm.IStringParameter
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyPair) Type() KeyPairType {
	var returns KeyPairType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewKeyPair(scope constructs.Construct, id *string, props *KeyPairProps) KeyPair {
	_init_.Initialize()

	if err := validateNewKeyPairParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyPair{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.KeyPair",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKeyPair_Override(k KeyPair, scope constructs.Construct, id *string, props *KeyPairProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.KeyPair",
		[]interface{}{scope, id, props},
		k,
	)
}

// Imports a key pair with a name and optional type.
func KeyPair_FromKeyPairAttributes(scope constructs.Construct, id *string, attrs *KeyPairAttributes) IKeyPair {
	_init_.Initialize()

	if err := validateKeyPair_FromKeyPairAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IKeyPair

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"fromKeyPairAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a key pair based on the name.
func KeyPair_FromKeyPairName(scope constructs.Construct, id *string, keyPairName *string) IKeyPair {
	_init_.Initialize()

	if err := validateKeyPair_FromKeyPairNameParameters(scope, id, keyPairName); err != nil {
		panic(err)
	}
	var returns IKeyPair

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"fromKeyPairName",
		[]interface{}{scope, id, keyPairName},
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
func KeyPair_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyPair_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func KeyPair_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKeyPair_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func KeyPair_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKeyPair_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func KeyPair_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.KeyPair",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyPair) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := k.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (k *jsiiProxy_KeyPair) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyPair) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (k *jsiiProxy_KeyPair) GetResourceNameAttribute(nameAttr *string) *string {
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

func (k *jsiiProxy_KeyPair) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

