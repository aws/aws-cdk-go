package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A record to delegate further lookups to a different set of name servers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cidrRoutingConfig cidrRoutingConfig
//   var geoLocation geoLocation
//   var healthCheck healthCheck
//   var hostedZone hostedZone
//
//   zoneDelegationRecord := awscdk.Aws_route53.NewZoneDelegationRecord(this, jsii.String("MyZoneDelegationRecord"), &ZoneDelegationRecordProps{
//   	NameServers: []*string{
//   		jsii.String("nameServers"),
//   	},
//   	Zone: hostedZone,
//
//   	// the properties below are optional
//   	CidrRoutingConfig: cidrRoutingConfig,
//   	Comment: jsii.String("comment"),
//   	DeleteExisting: jsii.Boolean(false),
//   	GeoLocation: geoLocation,
//   	HealthCheck: healthCheck,
//   	MultiValueAnswer: jsii.Boolean(false),
//   	RecordName: jsii.String("recordName"),
//   	Region: jsii.String("region"),
//   	SetIdentifier: jsii.String("setIdentifier"),
//   	Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//   	Weight: jsii.Number(123),
//   })
//
type ZoneDelegationRecord interface {
	RecordSet
	// The domain name of the record.
	DomainName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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

// The jsii proxy struct for ZoneDelegationRecord
type jsiiProxy_ZoneDelegationRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_ZoneDelegationRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewZoneDelegationRecord(scope constructs.Construct, id *string, props *ZoneDelegationRecordProps) ZoneDelegationRecord {
	_init_.Initialize()

	if err := validateNewZoneDelegationRecordParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneDelegationRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewZoneDelegationRecord_Override(z ZoneDelegationRecord, scope constructs.Construct, id *string, props *ZoneDelegationRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		[]interface{}{scope, id, props},
		z,
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
func ZoneDelegationRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneDelegationRecord_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ZoneDelegationRecord_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateZoneDelegationRecord_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ZoneDelegationRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateZoneDelegationRecord_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ZoneDelegationRecord_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZoneDelegationRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := z.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (z *jsiiProxy_ZoneDelegationRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDelegationRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := z.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDelegationRecord) GetResourceNameAttribute(nameAttr *string) *string {
	if err := z.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDelegationRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

