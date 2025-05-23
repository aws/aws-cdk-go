package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// EndpointGroup construct.
//
// Example:
//   var listener listener
//
//   // Non-open ALB
//   var alb applicationLoadBalancer
//
//   // Remember that there is only one AGA security group per VPC.
//   var vpc vpc
//
//
//   endpointGroup := listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &ApplicationLoadBalancerEndpointOptions{
//   			PreserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//   agaSg := endpointGroup.ConnectionsPeer(jsii.String("GlobalAcceleratorSG"), vpc)
//
//   // Allow connections from the AGA to the ALB
//   alb.Connections.AllowFrom(agaSg, ec2.Port_Tcp(jsii.Number(443)))
//
type EndpointGroup interface {
	awscdk.Resource
	IEndpointGroup
	// EndpointGroup ARN.
	EndpointGroupArn() *string
	// The name of the endpoint group.
	EndpointGroupName() *string
	// The array of the endpoints in this endpoint group.
	Endpoints() *[]IEndpoint
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
	// Add an endpoint.
	AddEndpoint(endpoint IEndpoint)
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
	// Return an object that represents the Accelerator's Security Group.
	//
	// Uses a Custom Resource to look up the Security Group that Accelerator
	// creates at deploy time. Requires your VPC ID to perform the lookup.
	//
	// The Security Group will only be created if you enable **Client IP
	// Preservation** on any of the endpoints.
	//
	// You cannot manipulate the rules inside this security group, but you can
	// use this security group as a Peer in Connections rules on other
	// constructs.
	ConnectionsPeer(id *string, vpc awsec2.IVpc) awsec2.IPeer
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

// The jsii proxy struct for EndpointGroup
type jsiiProxy_EndpointGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IEndpointGroup
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Endpoints() *[]IEndpoint {
	var returns *[]IEndpoint
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewEndpointGroup(scope constructs.Construct, id *string, props *EndpointGroupProps) EndpointGroup {
	_init_.Initialize()

	if err := validateNewEndpointGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EndpointGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEndpointGroup_Override(e EndpointGroup, scope constructs.Construct, id *string, props *EndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from ARN.
func EndpointGroup_FromEndpointGroupArn(scope constructs.Construct, id *string, endpointGroupArn *string) IEndpointGroup {
	_init_.Initialize()

	if err := validateEndpointGroup_FromEndpointGroupArnParameters(scope, id, endpointGroupArn); err != nil {
		panic(err)
	}
	var returns IEndpointGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"fromEndpointGroupArn",
		[]interface{}{scope, id, endpointGroupArn},
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
func EndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEndpointGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func EndpointGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEndpointGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EndpointGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEndpointGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func EndpointGroup_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EndpointGroup) AddEndpoint(endpoint IEndpoint) {
	if err := e.validateAddEndpointParameters(endpoint); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addEndpoint",
		[]interface{}{endpoint},
	)
}

func (e *jsiiProxy_EndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_EndpointGroup) ConnectionsPeer(id *string, vpc awsec2.IVpc) awsec2.IPeer {
	if err := e.validateConnectionsPeerParameters(id, vpc); err != nil {
		panic(err)
	}
	var returns awsec2.IPeer

	_jsii_.Invoke(
		e,
		"connectionsPeer",
		[]interface{}{id, vpc},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := e.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := e.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

