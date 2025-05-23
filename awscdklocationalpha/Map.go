package awscdklocationalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdklocationalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The Amazon Location Service Map.
//
// Example:
//   location.NewMap(this, jsii.String("Map"), &MapProps{
//   	MapName: jsii.String("my-map"),
//   	Style: location.Style_VECTOR_ESRI_NAVIGATION,
//   	CustomLayers: []pOI{
//   		location.CustomLayer_*pOI,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/location/latest/developerguide/map-concepts.html
//
// Experimental.
type Map interface {
	awscdk.Resource
	IMap
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the Map.
	// Experimental.
	MapArn() *string
	// The timestamp for when the map resource was created in ISO 8601 format.
	// Experimental.
	MapCreateTime() *string
	// The name of the map.
	// Experimental.
	MapName() *string
	// The timestamp for when the map resource was last updated in ISO 8601 format.
	// Experimental.
	MapUpdateTime() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given principal identity permissions to perform the actions on this map.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to rendering a map resource.
	// Experimental.
	GrantRendering(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Map
type jsiiProxy_Map struct {
	internal.Type__awscdkResource
	jsiiProxy_IMap
}

func (j *jsiiProxy_Map) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) MapArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) MapCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) MapName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) MapUpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewMap(scope constructs.Construct, id *string, props *MapProps) Map {
	_init_.Initialize()

	if err := validateNewMapParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Map{}

	_jsii_.Create(
		"@aws-cdk/aws-location-alpha.Map",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMap_Override(m Map, scope constructs.Construct, id *string, props *MapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-location-alpha.Map",
		[]interface{}{scope, id, props},
		m,
	)
}

// Use an existing map by ARN.
// Experimental.
func Map_FromMapArn(scope constructs.Construct, id *string, mapArn *string) IMap {
	_init_.Initialize()

	if err := validateMap_FromMapArnParameters(scope, id, mapArn); err != nil {
		panic(err)
	}
	var returns IMap

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-location-alpha.Map",
		"fromMapArn",
		[]interface{}{scope, id, mapArn},
		&returns,
	)

	return returns
}

// Use an existing map by name.
// Experimental.
func Map_FromMapName(scope constructs.Construct, id *string, mapName *string) IMap {
	_init_.Initialize()

	if err := validateMap_FromMapNameParameters(scope, id, mapName); err != nil {
		panic(err)
	}
	var returns IMap

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-location-alpha.Map",
		"fromMapName",
		[]interface{}{scope, id, mapName},
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
// Experimental.
func Map_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-location-alpha.Map",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Map_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMap_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-location-alpha.Map",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Map_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMap_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-location-alpha.Map",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Map_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-location-alpha.Map",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Map) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_Map) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := m.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) GetResourceNameAttribute(nameAttr *string) *string {
	if err := m.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := m.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		m,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) GrantRendering(grantee awsiam.IGrantable) awsiam.Grant {
	if err := m.validateGrantRenderingParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		m,
		"grantRendering",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

