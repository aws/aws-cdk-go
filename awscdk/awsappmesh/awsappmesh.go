package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppMesh::GatewayRoute`.
type CfnGatewayRoute interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrGatewayRouteName() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	AttrVirtualGatewayName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	GatewayRouteName() *string
	SetGatewayRouteName(val *string)
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualGatewayName() *string
	SetVirtualGatewayName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGatewayRoute
type jsiiProxy_CfnGatewayRoute struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGatewayRoute) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrGatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) AttrVirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) GatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoute) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::GatewayRoute`.
func NewCfnGatewayRoute(scope constructs.Construct, id *string, props *CfnGatewayRouteProps) CfnGatewayRoute {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayRoute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::GatewayRoute`.
func NewCfnGatewayRoute_Override(c CfnGatewayRoute, scope constructs.Construct, id *string, props *CfnGatewayRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGatewayRoute) SetGatewayRouteName(val *string) {
	_jsii_.Set(
		j,
		"gatewayRouteName",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayRoute) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayRoute) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayRoute) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayRoute) SetVirtualGatewayName(val *string) {
	_jsii_.Set(
		j,
		"virtualGatewayName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnGatewayRoute_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGatewayRoute_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnGatewayRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayRoute_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGatewayRoute) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnGatewayRoute_GatewayRouteSpecProperty struct {
	// `CfnGatewayRoute.GatewayRouteSpecProperty.GrpcRoute`.
	GrpcRoute interface{} `json:"grpcRoute"`
	// `CfnGatewayRoute.GatewayRouteSpecProperty.Http2Route`.
	Http2Route interface{} `json:"http2Route"`
	// `CfnGatewayRoute.GatewayRouteSpecProperty.HttpRoute`.
	HttpRoute interface{} `json:"httpRoute"`
}

type CfnGatewayRoute_GatewayRouteTargetProperty struct {
	// `CfnGatewayRoute.GatewayRouteTargetProperty.VirtualService`.
	VirtualService interface{} `json:"virtualService"`
}

type CfnGatewayRoute_GatewayRouteVirtualServiceProperty struct {
	// `CfnGatewayRoute.GatewayRouteVirtualServiceProperty.VirtualServiceName`.
	VirtualServiceName *string `json:"virtualServiceName"`
}

type CfnGatewayRoute_GrpcGatewayRouteActionProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteActionProperty.Target`.
	Target interface{} `json:"target"`
}

type CfnGatewayRoute_GrpcGatewayRouteMatchProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteMatchProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
}

type CfnGatewayRoute_GrpcGatewayRouteProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnGatewayRoute.GrpcGatewayRouteProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnGatewayRoute_HttpGatewayRouteActionProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteActionProperty.Target`.
	Target interface{} `json:"target"`
}

type CfnGatewayRoute_HttpGatewayRouteMatchProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
}

type CfnGatewayRoute_HttpGatewayRouteProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnGatewayRoute.HttpGatewayRouteProperty.Match`.
	Match interface{} `json:"match"`
}

// Properties for defining a `AWS::AppMesh::GatewayRoute`.
type CfnGatewayRouteProps struct {
	// `AWS::AppMesh::GatewayRoute.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::GatewayRoute.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::GatewayRoute.VirtualGatewayName`.
	VirtualGatewayName *string `json:"virtualGatewayName"`
	// `AWS::AppMesh::GatewayRoute.GatewayRouteName`.
	GatewayRouteName *string `json:"gatewayRouteName"`
	// `AWS::AppMesh::GatewayRoute.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::GatewayRoute.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::AppMesh::Mesh`.
type CfnMesh interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMesh
type jsiiProxy_CfnMesh struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMesh) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMesh) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::Mesh`.
func NewCfnMesh(scope constructs.Construct, id *string, props *CfnMeshProps) CfnMesh {
	_init_.Initialize()

	j := jsiiProxy_CfnMesh{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::Mesh`.
func NewCfnMesh_Override(c CfnMesh, scope constructs.Construct, id *string, props *CfnMeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMesh) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnMesh) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnMesh_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMesh_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnMesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMesh_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnMesh) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnMesh) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMesh) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnMesh) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnMesh) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnMesh) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnMesh) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnMesh) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMesh) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnMesh) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMesh) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMesh) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnMesh) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnMesh) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnMesh) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnMesh_EgressFilterProperty struct {
	// `CfnMesh.EgressFilterProperty.Type`.
	Type *string `json:"type"`
}

type CfnMesh_MeshSpecProperty struct {
	// `CfnMesh.MeshSpecProperty.EgressFilter`.
	EgressFilter interface{} `json:"egressFilter"`
}

// Properties for defining a `AWS::AppMesh::Mesh`.
type CfnMeshProps struct {
	// `AWS::AppMesh::Mesh.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::Mesh.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::Mesh.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::AppMesh::Route`.
type CfnRoute interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrRouteName() *string
	AttrUid() *string
	AttrVirtualRouterName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	RouteName() *string
	SetRouteName(val *string)
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualRouterName() *string
	SetVirtualRouterName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRoute
type jsiiProxy_CfnRoute struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRoute) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AttrVirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualRouterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::Route`.
func NewCfnRoute(scope constructs.Construct, id *string, props *CfnRouteProps) CfnRoute {
	_init_.Initialize()

	j := jsiiProxy_CfnRoute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::Route`.
func NewCfnRoute_Override(c CfnRoute, scope constructs.Construct, id *string, props *CfnRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRoute) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetRouteName(val *string) {
	_jsii_.Set(
		j,
		"routeName",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetVirtualRouterName(val *string) {
	_jsii_.Set(
		j,
		"virtualRouterName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnRoute_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRoute_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoute_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnRoute) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnRoute) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnRoute) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnRoute) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRoute) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRoute) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnRoute) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnRoute) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnRoute_DurationProperty struct {
	// `CfnRoute.DurationProperty.Unit`.
	Unit *string `json:"unit"`
	// `CfnRoute.DurationProperty.Value`.
	Value *float64 `json:"value"`
}

type CfnRoute_GrpcRetryPolicyProperty struct {
	// `CfnRoute.GrpcRetryPolicyProperty.MaxRetries`.
	MaxRetries *float64 `json:"maxRetries"`
	// `CfnRoute.GrpcRetryPolicyProperty.PerRetryTimeout`.
	PerRetryTimeout interface{} `json:"perRetryTimeout"`
	// `CfnRoute.GrpcRetryPolicyProperty.GrpcRetryEvents`.
	GrpcRetryEvents *[]*string `json:"grpcRetryEvents"`
	// `CfnRoute.GrpcRetryPolicyProperty.HttpRetryEvents`.
	HttpRetryEvents *[]*string `json:"httpRetryEvents"`
	// `CfnRoute.GrpcRetryPolicyProperty.TcpRetryEvents`.
	TcpRetryEvents *[]*string `json:"tcpRetryEvents"`
}

type CfnRoute_GrpcRouteActionProperty struct {
	// `CfnRoute.GrpcRouteActionProperty.WeightedTargets`.
	WeightedTargets interface{} `json:"weightedTargets"`
}

type CfnRoute_GrpcRouteMatchProperty struct {
	// `CfnRoute.GrpcRouteMatchProperty.Metadata`.
	Metadata interface{} `json:"metadata"`
	// `CfnRoute.GrpcRouteMatchProperty.MethodName`.
	MethodName *string `json:"methodName"`
	// `CfnRoute.GrpcRouteMatchProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
}

type CfnRoute_GrpcRouteMetadataMatchMethodProperty struct {
	// `CfnRoute.GrpcRouteMetadataMatchMethodProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnRoute.GrpcRouteMetadataMatchMethodProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnRoute.GrpcRouteMetadataMatchMethodProperty.Range`.
	Range interface{} `json:"range"`
	// `CfnRoute.GrpcRouteMetadataMatchMethodProperty.Regex`.
	Regex *string `json:"regex"`
	// `CfnRoute.GrpcRouteMetadataMatchMethodProperty.Suffix`.
	Suffix *string `json:"suffix"`
}

type CfnRoute_GrpcRouteMetadataProperty struct {
	// `CfnRoute.GrpcRouteMetadataProperty.Name`.
	Name *string `json:"name"`
	// `CfnRoute.GrpcRouteMetadataProperty.Invert`.
	Invert interface{} `json:"invert"`
	// `CfnRoute.GrpcRouteMetadataProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnRoute_GrpcRouteProperty struct {
	// `CfnRoute.GrpcRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnRoute.GrpcRouteProperty.Match`.
	Match interface{} `json:"match"`
	// `CfnRoute.GrpcRouteProperty.RetryPolicy`.
	RetryPolicy interface{} `json:"retryPolicy"`
	// `CfnRoute.GrpcRouteProperty.Timeout`.
	Timeout interface{} `json:"timeout"`
}

type CfnRoute_GrpcTimeoutProperty struct {
	// `CfnRoute.GrpcTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
	// `CfnRoute.GrpcTimeoutProperty.PerRequest`.
	PerRequest interface{} `json:"perRequest"`
}

type CfnRoute_HeaderMatchMethodProperty struct {
	// `CfnRoute.HeaderMatchMethodProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnRoute.HeaderMatchMethodProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnRoute.HeaderMatchMethodProperty.Range`.
	Range interface{} `json:"range"`
	// `CfnRoute.HeaderMatchMethodProperty.Regex`.
	Regex *string `json:"regex"`
	// `CfnRoute.HeaderMatchMethodProperty.Suffix`.
	Suffix *string `json:"suffix"`
}

type CfnRoute_HttpRetryPolicyProperty struct {
	// `CfnRoute.HttpRetryPolicyProperty.MaxRetries`.
	MaxRetries *float64 `json:"maxRetries"`
	// `CfnRoute.HttpRetryPolicyProperty.PerRetryTimeout`.
	PerRetryTimeout interface{} `json:"perRetryTimeout"`
	// `CfnRoute.HttpRetryPolicyProperty.HttpRetryEvents`.
	HttpRetryEvents *[]*string `json:"httpRetryEvents"`
	// `CfnRoute.HttpRetryPolicyProperty.TcpRetryEvents`.
	TcpRetryEvents *[]*string `json:"tcpRetryEvents"`
}

type CfnRoute_HttpRouteActionProperty struct {
	// `CfnRoute.HttpRouteActionProperty.WeightedTargets`.
	WeightedTargets interface{} `json:"weightedTargets"`
}

type CfnRoute_HttpRouteHeaderProperty struct {
	// `CfnRoute.HttpRouteHeaderProperty.Name`.
	Name *string `json:"name"`
	// `CfnRoute.HttpRouteHeaderProperty.Invert`.
	Invert interface{} `json:"invert"`
	// `CfnRoute.HttpRouteHeaderProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnRoute_HttpRouteMatchProperty struct {
	// `CfnRoute.HttpRouteMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnRoute.HttpRouteMatchProperty.Headers`.
	Headers interface{} `json:"headers"`
	// `CfnRoute.HttpRouteMatchProperty.Method`.
	Method *string `json:"method"`
	// `CfnRoute.HttpRouteMatchProperty.Scheme`.
	Scheme *string `json:"scheme"`
}

type CfnRoute_HttpRouteProperty struct {
	// `CfnRoute.HttpRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnRoute.HttpRouteProperty.Match`.
	Match interface{} `json:"match"`
	// `CfnRoute.HttpRouteProperty.RetryPolicy`.
	RetryPolicy interface{} `json:"retryPolicy"`
	// `CfnRoute.HttpRouteProperty.Timeout`.
	Timeout interface{} `json:"timeout"`
}

type CfnRoute_HttpTimeoutProperty struct {
	// `CfnRoute.HttpTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
	// `CfnRoute.HttpTimeoutProperty.PerRequest`.
	PerRequest interface{} `json:"perRequest"`
}

type CfnRoute_MatchRangeProperty struct {
	// `CfnRoute.MatchRangeProperty.End`.
	End *float64 `json:"end"`
	// `CfnRoute.MatchRangeProperty.Start`.
	Start *float64 `json:"start"`
}

type CfnRoute_RouteSpecProperty struct {
	// `CfnRoute.RouteSpecProperty.GrpcRoute`.
	GrpcRoute interface{} `json:"grpcRoute"`
	// `CfnRoute.RouteSpecProperty.Http2Route`.
	Http2Route interface{} `json:"http2Route"`
	// `CfnRoute.RouteSpecProperty.HttpRoute`.
	HttpRoute interface{} `json:"httpRoute"`
	// `CfnRoute.RouteSpecProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnRoute.RouteSpecProperty.TcpRoute`.
	TcpRoute interface{} `json:"tcpRoute"`
}

type CfnRoute_TcpRouteActionProperty struct {
	// `CfnRoute.TcpRouteActionProperty.WeightedTargets`.
	WeightedTargets interface{} `json:"weightedTargets"`
}

type CfnRoute_TcpRouteProperty struct {
	// `CfnRoute.TcpRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnRoute.TcpRouteProperty.Timeout`.
	Timeout interface{} `json:"timeout"`
}

type CfnRoute_TcpTimeoutProperty struct {
	// `CfnRoute.TcpTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
}

type CfnRoute_WeightedTargetProperty struct {
	// `CfnRoute.WeightedTargetProperty.VirtualNode`.
	VirtualNode *string `json:"virtualNode"`
	// `CfnRoute.WeightedTargetProperty.Weight`.
	Weight *float64 `json:"weight"`
}

// Properties for defining a `AWS::AppMesh::Route`.
type CfnRouteProps struct {
	// `AWS::AppMesh::Route.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::Route.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::Route.VirtualRouterName`.
	VirtualRouterName *string `json:"virtualRouterName"`
	// `AWS::AppMesh::Route.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::Route.RouteName`.
	RouteName *string `json:"routeName"`
	// `AWS::AppMesh::Route.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::AppMesh::VirtualGateway`.
type CfnVirtualGateway interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	AttrVirtualGatewayName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualGatewayName() *string
	SetVirtualGatewayName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVirtualGateway
type jsiiProxy_CfnVirtualGateway struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVirtualGateway) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) AttrVirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGateway) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::VirtualGateway`.
func NewCfnVirtualGateway(scope constructs.Construct, id *string, props *CfnVirtualGatewayProps) CfnVirtualGateway {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualGateway`.
func NewCfnVirtualGateway_Override(c CfnVirtualGateway, scope constructs.Construct, id *string, props *CfnVirtualGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualGateway) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualGateway) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualGateway) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualGateway) SetVirtualGatewayName(val *string) {
	_jsii_.Set(
		j,
		"virtualGatewayName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnVirtualGateway_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVirtualGateway_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnVirtualGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualGateway_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualGateway) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnVirtualGateway_SubjectAlternativeNameMatchersProperty struct {
	// `CfnVirtualGateway.SubjectAlternativeNameMatchersProperty.Exact`.
	Exact *[]*string `json:"exact"`
}

type CfnVirtualGateway_SubjectAlternativeNamesProperty struct {
	// `CfnVirtualGateway.SubjectAlternativeNamesProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnVirtualGateway_VirtualGatewayAccessLogProperty struct {
	// `CfnVirtualGateway.VirtualGatewayAccessLogProperty.File`.
	File interface{} `json:"file"`
}

type CfnVirtualGateway_VirtualGatewayBackendDefaultsProperty struct {
	// `CfnVirtualGateway.VirtualGatewayBackendDefaultsProperty.ClientPolicy`.
	ClientPolicy interface{} `json:"clientPolicy"`
}

type CfnVirtualGateway_VirtualGatewayClientPolicyProperty struct {
	// `CfnVirtualGateway.VirtualGatewayClientPolicyProperty.TLS`.
	Tls interface{} `json:"tls"`
}

type CfnVirtualGateway_VirtualGatewayClientPolicyTlsProperty struct {
	// `CfnVirtualGateway.VirtualGatewayClientPolicyTlsProperty.Validation`.
	Validation interface{} `json:"validation"`
	// `CfnVirtualGateway.VirtualGatewayClientPolicyTlsProperty.Certificate`.
	Certificate interface{} `json:"certificate"`
	// `CfnVirtualGateway.VirtualGatewayClientPolicyTlsProperty.Enforce`.
	Enforce interface{} `json:"enforce"`
	// `CfnVirtualGateway.VirtualGatewayClientPolicyTlsProperty.Ports`.
	Ports interface{} `json:"ports"`
}

type CfnVirtualGateway_VirtualGatewayClientTlsCertificateProperty struct {
	// `CfnVirtualGateway.VirtualGatewayClientTlsCertificateProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualGateway.VirtualGatewayClientTlsCertificateProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualGateway_VirtualGatewayConnectionPoolProperty struct {
	// `CfnVirtualGateway.VirtualGatewayConnectionPoolProperty.GRPC`.
	Grpc interface{} `json:"grpc"`
	// `CfnVirtualGateway.VirtualGatewayConnectionPoolProperty.HTTP`.
	Http interface{} `json:"http"`
	// `CfnVirtualGateway.VirtualGatewayConnectionPoolProperty.HTTP2`.
	Http2 interface{} `json:"http2"`
}

type CfnVirtualGateway_VirtualGatewayFileAccessLogProperty struct {
	// `CfnVirtualGateway.VirtualGatewayFileAccessLogProperty.Path`.
	Path *string `json:"path"`
}

type CfnVirtualGateway_VirtualGatewayGrpcConnectionPoolProperty struct {
	// `CfnVirtualGateway.VirtualGatewayGrpcConnectionPoolProperty.MaxRequests`.
	MaxRequests *float64 `json:"maxRequests"`
}

type CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty struct {
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.HealthyThreshold`.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.IntervalMillis`.
	IntervalMillis *float64 `json:"intervalMillis"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.TimeoutMillis`.
	TimeoutMillis *float64 `json:"timeoutMillis"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.UnhealthyThreshold`.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.Path`.
	Path *string `json:"path"`
	// `CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty.Port`.
	Port *float64 `json:"port"`
}

type CfnVirtualGateway_VirtualGatewayHttp2ConnectionPoolProperty struct {
	// `CfnVirtualGateway.VirtualGatewayHttp2ConnectionPoolProperty.MaxRequests`.
	MaxRequests *float64 `json:"maxRequests"`
}

type CfnVirtualGateway_VirtualGatewayHttpConnectionPoolProperty struct {
	// `CfnVirtualGateway.VirtualGatewayHttpConnectionPoolProperty.MaxConnections`.
	MaxConnections *float64 `json:"maxConnections"`
	// `CfnVirtualGateway.VirtualGatewayHttpConnectionPoolProperty.MaxPendingRequests`.
	MaxPendingRequests *float64 `json:"maxPendingRequests"`
}

type CfnVirtualGateway_VirtualGatewayListenerProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerProperty.PortMapping`.
	PortMapping interface{} `json:"portMapping"`
	// `CfnVirtualGateway.VirtualGatewayListenerProperty.ConnectionPool`.
	ConnectionPool interface{} `json:"connectionPool"`
	// `CfnVirtualGateway.VirtualGatewayListenerProperty.HealthCheck`.
	HealthCheck interface{} `json:"healthCheck"`
	// `CfnVirtualGateway.VirtualGatewayListenerProperty.TLS`.
	Tls interface{} `json:"tls"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsAcmCertificateProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsAcmCertificateProperty.CertificateArn`.
	CertificateArn *string `json:"certificateArn"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsCertificateProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsCertificateProperty.ACM`.
	Acm interface{} `json:"acm"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsCertificateProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsCertificateProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsFileCertificateProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsFileCertificateProperty.CertificateChain`.
	CertificateChain *string `json:"certificateChain"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsFileCertificateProperty.PrivateKey`.
	PrivateKey *string `json:"privateKey"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsProperty.Certificate`.
	Certificate interface{} `json:"certificate"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsProperty.Mode`.
	Mode *string `json:"mode"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsProperty.Validation`.
	Validation interface{} `json:"validation"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsSdsCertificateProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsSdsCertificateProperty.SecretName`.
	SecretName *string `json:"secretName"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextProperty.Trust`.
	Trust interface{} `json:"trust"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextProperty.SubjectAlternativeNames`.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames"`
}

type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextTrustProperty struct {
	// `CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextTrustProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextTrustProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualGateway_VirtualGatewayLoggingProperty struct {
	// `CfnVirtualGateway.VirtualGatewayLoggingProperty.AccessLog`.
	AccessLog interface{} `json:"accessLog"`
}

type CfnVirtualGateway_VirtualGatewayPortMappingProperty struct {
	// `CfnVirtualGateway.VirtualGatewayPortMappingProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnVirtualGateway.VirtualGatewayPortMappingProperty.Protocol`.
	Protocol *string `json:"protocol"`
}

type CfnVirtualGateway_VirtualGatewaySpecProperty struct {
	// `CfnVirtualGateway.VirtualGatewaySpecProperty.Listeners`.
	Listeners interface{} `json:"listeners"`
	// `CfnVirtualGateway.VirtualGatewaySpecProperty.BackendDefaults`.
	BackendDefaults interface{} `json:"backendDefaults"`
	// `CfnVirtualGateway.VirtualGatewaySpecProperty.Logging`.
	Logging interface{} `json:"logging"`
}

type CfnVirtualGateway_VirtualGatewayTlsValidationContextAcmTrustProperty struct {
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextAcmTrustProperty.CertificateAuthorityArns`.
	CertificateAuthorityArns *[]*string `json:"certificateAuthorityArns"`
}

type CfnVirtualGateway_VirtualGatewayTlsValidationContextFileTrustProperty struct {
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextFileTrustProperty.CertificateChain`.
	CertificateChain *string `json:"certificateChain"`
}

type CfnVirtualGateway_VirtualGatewayTlsValidationContextProperty struct {
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextProperty.Trust`.
	Trust interface{} `json:"trust"`
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextProperty.SubjectAlternativeNames`.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames"`
}

type CfnVirtualGateway_VirtualGatewayTlsValidationContextSdsTrustProperty struct {
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextSdsTrustProperty.SecretName`.
	SecretName *string `json:"secretName"`
}

type CfnVirtualGateway_VirtualGatewayTlsValidationContextTrustProperty struct {
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextTrustProperty.ACM`.
	Acm interface{} `json:"acm"`
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextTrustProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualGateway.VirtualGatewayTlsValidationContextTrustProperty.SDS`.
	Sds interface{} `json:"sds"`
}

// Properties for defining a `AWS::AppMesh::VirtualGateway`.
type CfnVirtualGatewayProps struct {
	// `AWS::AppMesh::VirtualGateway.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::VirtualGateway.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::VirtualGateway.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::VirtualGateway.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::AppMesh::VirtualGateway.VirtualGatewayName`.
	VirtualGatewayName *string `json:"virtualGatewayName"`
}

// A CloudFormation `AWS::AppMesh::VirtualNode`.
type CfnVirtualNode interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	AttrVirtualNodeName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualNodeName() *string
	SetVirtualNodeName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVirtualNode
type jsiiProxy_CfnVirtualNode struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVirtualNode) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrVirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualNodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) VirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNode(scope constructs.Construct, id *string, props *CfnVirtualNodeProps) CfnVirtualNode {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualNode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNode_Override(c CfnVirtualNode, scope constructs.Construct, id *string, props *CfnVirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetVirtualNodeName(val *string) {
	_jsii_.Set(
		j,
		"virtualNodeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnVirtualNode_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVirtualNode_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnVirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualNode_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualNode) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnVirtualNode) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnVirtualNode_AccessLogProperty struct {
	// `CfnVirtualNode.AccessLogProperty.File`.
	File interface{} `json:"file"`
}

type CfnVirtualNode_AwsCloudMapInstanceAttributeProperty struct {
	// `CfnVirtualNode.AwsCloudMapInstanceAttributeProperty.Key`.
	Key *string `json:"key"`
	// `CfnVirtualNode.AwsCloudMapInstanceAttributeProperty.Value`.
	Value *string `json:"value"`
}

type CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty struct {
	// `CfnVirtualNode.AwsCloudMapServiceDiscoveryProperty.NamespaceName`.
	NamespaceName *string `json:"namespaceName"`
	// `CfnVirtualNode.AwsCloudMapServiceDiscoveryProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
	// `CfnVirtualNode.AwsCloudMapServiceDiscoveryProperty.Attributes`.
	Attributes interface{} `json:"attributes"`
}

type CfnVirtualNode_BackendDefaultsProperty struct {
	// `CfnVirtualNode.BackendDefaultsProperty.ClientPolicy`.
	ClientPolicy interface{} `json:"clientPolicy"`
}

type CfnVirtualNode_BackendProperty struct {
	// `CfnVirtualNode.BackendProperty.VirtualService`.
	VirtualService interface{} `json:"virtualService"`
}

type CfnVirtualNode_ClientPolicyProperty struct {
	// `CfnVirtualNode.ClientPolicyProperty.TLS`.
	Tls interface{} `json:"tls"`
}

type CfnVirtualNode_ClientPolicyTlsProperty struct {
	// `CfnVirtualNode.ClientPolicyTlsProperty.Validation`.
	Validation interface{} `json:"validation"`
	// `CfnVirtualNode.ClientPolicyTlsProperty.Certificate`.
	Certificate interface{} `json:"certificate"`
	// `CfnVirtualNode.ClientPolicyTlsProperty.Enforce`.
	Enforce interface{} `json:"enforce"`
	// `CfnVirtualNode.ClientPolicyTlsProperty.Ports`.
	Ports interface{} `json:"ports"`
}

type CfnVirtualNode_ClientTlsCertificateProperty struct {
	// `CfnVirtualNode.ClientTlsCertificateProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualNode.ClientTlsCertificateProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualNode_DnsServiceDiscoveryProperty struct {
	// `CfnVirtualNode.DnsServiceDiscoveryProperty.Hostname`.
	Hostname *string `json:"hostname"`
}

type CfnVirtualNode_DurationProperty struct {
	// `CfnVirtualNode.DurationProperty.Unit`.
	Unit *string `json:"unit"`
	// `CfnVirtualNode.DurationProperty.Value`.
	Value *float64 `json:"value"`
}

type CfnVirtualNode_FileAccessLogProperty struct {
	// `CfnVirtualNode.FileAccessLogProperty.Path`.
	Path *string `json:"path"`
}

type CfnVirtualNode_GrpcTimeoutProperty struct {
	// `CfnVirtualNode.GrpcTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
	// `CfnVirtualNode.GrpcTimeoutProperty.PerRequest`.
	PerRequest interface{} `json:"perRequest"`
}

type CfnVirtualNode_HealthCheckProperty struct {
	// `CfnVirtualNode.HealthCheckProperty.HealthyThreshold`.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// `CfnVirtualNode.HealthCheckProperty.IntervalMillis`.
	IntervalMillis *float64 `json:"intervalMillis"`
	// `CfnVirtualNode.HealthCheckProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnVirtualNode.HealthCheckProperty.TimeoutMillis`.
	TimeoutMillis *float64 `json:"timeoutMillis"`
	// `CfnVirtualNode.HealthCheckProperty.UnhealthyThreshold`.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
	// `CfnVirtualNode.HealthCheckProperty.Path`.
	Path *string `json:"path"`
	// `CfnVirtualNode.HealthCheckProperty.Port`.
	Port *float64 `json:"port"`
}

type CfnVirtualNode_HttpTimeoutProperty struct {
	// `CfnVirtualNode.HttpTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
	// `CfnVirtualNode.HttpTimeoutProperty.PerRequest`.
	PerRequest interface{} `json:"perRequest"`
}

type CfnVirtualNode_ListenerProperty struct {
	// `CfnVirtualNode.ListenerProperty.PortMapping`.
	PortMapping interface{} `json:"portMapping"`
	// `CfnVirtualNode.ListenerProperty.ConnectionPool`.
	ConnectionPool interface{} `json:"connectionPool"`
	// `CfnVirtualNode.ListenerProperty.HealthCheck`.
	HealthCheck interface{} `json:"healthCheck"`
	// `CfnVirtualNode.ListenerProperty.OutlierDetection`.
	OutlierDetection interface{} `json:"outlierDetection"`
	// `CfnVirtualNode.ListenerProperty.Timeout`.
	Timeout interface{} `json:"timeout"`
	// `CfnVirtualNode.ListenerProperty.TLS`.
	Tls interface{} `json:"tls"`
}

type CfnVirtualNode_ListenerTimeoutProperty struct {
	// `CfnVirtualNode.ListenerTimeoutProperty.GRPC`.
	Grpc interface{} `json:"grpc"`
	// `CfnVirtualNode.ListenerTimeoutProperty.HTTP`.
	Http interface{} `json:"http"`
	// `CfnVirtualNode.ListenerTimeoutProperty.HTTP2`.
	Http2 interface{} `json:"http2"`
	// `CfnVirtualNode.ListenerTimeoutProperty.TCP`.
	Tcp interface{} `json:"tcp"`
}

type CfnVirtualNode_ListenerTlsAcmCertificateProperty struct {
	// `CfnVirtualNode.ListenerTlsAcmCertificateProperty.CertificateArn`.
	CertificateArn *string `json:"certificateArn"`
}

type CfnVirtualNode_ListenerTlsCertificateProperty struct {
	// `CfnVirtualNode.ListenerTlsCertificateProperty.ACM`.
	Acm interface{} `json:"acm"`
	// `CfnVirtualNode.ListenerTlsCertificateProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualNode.ListenerTlsCertificateProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualNode_ListenerTlsFileCertificateProperty struct {
	// `CfnVirtualNode.ListenerTlsFileCertificateProperty.CertificateChain`.
	CertificateChain *string `json:"certificateChain"`
	// `CfnVirtualNode.ListenerTlsFileCertificateProperty.PrivateKey`.
	PrivateKey *string `json:"privateKey"`
}

type CfnVirtualNode_ListenerTlsProperty struct {
	// `CfnVirtualNode.ListenerTlsProperty.Certificate`.
	Certificate interface{} `json:"certificate"`
	// `CfnVirtualNode.ListenerTlsProperty.Mode`.
	Mode *string `json:"mode"`
	// `CfnVirtualNode.ListenerTlsProperty.Validation`.
	Validation interface{} `json:"validation"`
}

type CfnVirtualNode_ListenerTlsSdsCertificateProperty struct {
	// `CfnVirtualNode.ListenerTlsSdsCertificateProperty.SecretName`.
	SecretName *string `json:"secretName"`
}

type CfnVirtualNode_ListenerTlsValidationContextProperty struct {
	// `CfnVirtualNode.ListenerTlsValidationContextProperty.Trust`.
	Trust interface{} `json:"trust"`
	// `CfnVirtualNode.ListenerTlsValidationContextProperty.SubjectAlternativeNames`.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames"`
}

type CfnVirtualNode_ListenerTlsValidationContextTrustProperty struct {
	// `CfnVirtualNode.ListenerTlsValidationContextTrustProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualNode.ListenerTlsValidationContextTrustProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualNode_LoggingProperty struct {
	// `CfnVirtualNode.LoggingProperty.AccessLog`.
	AccessLog interface{} `json:"accessLog"`
}

type CfnVirtualNode_OutlierDetectionProperty struct {
	// `CfnVirtualNode.OutlierDetectionProperty.BaseEjectionDuration`.
	BaseEjectionDuration interface{} `json:"baseEjectionDuration"`
	// `CfnVirtualNode.OutlierDetectionProperty.Interval`.
	Interval interface{} `json:"interval"`
	// `CfnVirtualNode.OutlierDetectionProperty.MaxEjectionPercent`.
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	// `CfnVirtualNode.OutlierDetectionProperty.MaxServerErrors`.
	MaxServerErrors *float64 `json:"maxServerErrors"`
}

type CfnVirtualNode_PortMappingProperty struct {
	// `CfnVirtualNode.PortMappingProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnVirtualNode.PortMappingProperty.Protocol`.
	Protocol *string `json:"protocol"`
}

type CfnVirtualNode_ServiceDiscoveryProperty struct {
	// `CfnVirtualNode.ServiceDiscoveryProperty.AWSCloudMap`.
	AwsCloudMap interface{} `json:"awsCloudMap"`
	// `CfnVirtualNode.ServiceDiscoveryProperty.DNS`.
	Dns interface{} `json:"dns"`
}

type CfnVirtualNode_SubjectAlternativeNameMatchersProperty struct {
	// `CfnVirtualNode.SubjectAlternativeNameMatchersProperty.Exact`.
	Exact *[]*string `json:"exact"`
}

type CfnVirtualNode_SubjectAlternativeNamesProperty struct {
	// `CfnVirtualNode.SubjectAlternativeNamesProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnVirtualNode_TcpTimeoutProperty struct {
	// `CfnVirtualNode.TcpTimeoutProperty.Idle`.
	Idle interface{} `json:"idle"`
}

type CfnVirtualNode_TlsValidationContextAcmTrustProperty struct {
	// `CfnVirtualNode.TlsValidationContextAcmTrustProperty.CertificateAuthorityArns`.
	CertificateAuthorityArns *[]*string `json:"certificateAuthorityArns"`
}

type CfnVirtualNode_TlsValidationContextFileTrustProperty struct {
	// `CfnVirtualNode.TlsValidationContextFileTrustProperty.CertificateChain`.
	CertificateChain *string `json:"certificateChain"`
}

type CfnVirtualNode_TlsValidationContextProperty struct {
	// `CfnVirtualNode.TlsValidationContextProperty.Trust`.
	Trust interface{} `json:"trust"`
	// `CfnVirtualNode.TlsValidationContextProperty.SubjectAlternativeNames`.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames"`
}

type CfnVirtualNode_TlsValidationContextSdsTrustProperty struct {
	// `CfnVirtualNode.TlsValidationContextSdsTrustProperty.SecretName`.
	SecretName *string `json:"secretName"`
}

type CfnVirtualNode_TlsValidationContextTrustProperty struct {
	// `CfnVirtualNode.TlsValidationContextTrustProperty.ACM`.
	Acm interface{} `json:"acm"`
	// `CfnVirtualNode.TlsValidationContextTrustProperty.File`.
	File interface{} `json:"file"`
	// `CfnVirtualNode.TlsValidationContextTrustProperty.SDS`.
	Sds interface{} `json:"sds"`
}

type CfnVirtualNode_VirtualNodeConnectionPoolProperty struct {
	// `CfnVirtualNode.VirtualNodeConnectionPoolProperty.GRPC`.
	Grpc interface{} `json:"grpc"`
	// `CfnVirtualNode.VirtualNodeConnectionPoolProperty.HTTP`.
	Http interface{} `json:"http"`
	// `CfnVirtualNode.VirtualNodeConnectionPoolProperty.HTTP2`.
	Http2 interface{} `json:"http2"`
	// `CfnVirtualNode.VirtualNodeConnectionPoolProperty.TCP`.
	Tcp interface{} `json:"tcp"`
}

type CfnVirtualNode_VirtualNodeGrpcConnectionPoolProperty struct {
	// `CfnVirtualNode.VirtualNodeGrpcConnectionPoolProperty.MaxRequests`.
	MaxRequests *float64 `json:"maxRequests"`
}

type CfnVirtualNode_VirtualNodeHttp2ConnectionPoolProperty struct {
	// `CfnVirtualNode.VirtualNodeHttp2ConnectionPoolProperty.MaxRequests`.
	MaxRequests *float64 `json:"maxRequests"`
}

type CfnVirtualNode_VirtualNodeHttpConnectionPoolProperty struct {
	// `CfnVirtualNode.VirtualNodeHttpConnectionPoolProperty.MaxConnections`.
	MaxConnections *float64 `json:"maxConnections"`
	// `CfnVirtualNode.VirtualNodeHttpConnectionPoolProperty.MaxPendingRequests`.
	MaxPendingRequests *float64 `json:"maxPendingRequests"`
}

type CfnVirtualNode_VirtualNodeSpecProperty struct {
	// `CfnVirtualNode.VirtualNodeSpecProperty.BackendDefaults`.
	BackendDefaults interface{} `json:"backendDefaults"`
	// `CfnVirtualNode.VirtualNodeSpecProperty.Backends`.
	Backends interface{} `json:"backends"`
	// `CfnVirtualNode.VirtualNodeSpecProperty.Listeners`.
	Listeners interface{} `json:"listeners"`
	// `CfnVirtualNode.VirtualNodeSpecProperty.Logging`.
	Logging interface{} `json:"logging"`
	// `CfnVirtualNode.VirtualNodeSpecProperty.ServiceDiscovery`.
	ServiceDiscovery interface{} `json:"serviceDiscovery"`
}

type CfnVirtualNode_VirtualNodeTcpConnectionPoolProperty struct {
	// `CfnVirtualNode.VirtualNodeTcpConnectionPoolProperty.MaxConnections`.
	MaxConnections *float64 `json:"maxConnections"`
}

type CfnVirtualNode_VirtualServiceBackendProperty struct {
	// `CfnVirtualNode.VirtualServiceBackendProperty.VirtualServiceName`.
	VirtualServiceName *string `json:"virtualServiceName"`
	// `CfnVirtualNode.VirtualServiceBackendProperty.ClientPolicy`.
	ClientPolicy interface{} `json:"clientPolicy"`
}

// Properties for defining a `AWS::AppMesh::VirtualNode`.
type CfnVirtualNodeProps struct {
	// `AWS::AppMesh::VirtualNode.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::VirtualNode.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::VirtualNode.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::VirtualNode.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::AppMesh::VirtualNode.VirtualNodeName`.
	VirtualNodeName *string `json:"virtualNodeName"`
}

// A CloudFormation `AWS::AppMesh::VirtualRouter`.
type CfnVirtualRouter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	AttrVirtualRouterName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualRouterName() *string
	SetVirtualRouterName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVirtualRouter
type jsiiProxy_CfnVirtualRouter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVirtualRouter) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) AttrVirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualRouterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualRouter) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::VirtualRouter`.
func NewCfnVirtualRouter(scope constructs.Construct, id *string, props *CfnVirtualRouterProps) CfnVirtualRouter {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualRouter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualRouter`.
func NewCfnVirtualRouter_Override(c CfnVirtualRouter, scope constructs.Construct, id *string, props *CfnVirtualRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualRouter) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualRouter) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualRouter) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualRouter) SetVirtualRouterName(val *string) {
	_jsii_.Set(
		j,
		"virtualRouterName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnVirtualRouter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVirtualRouter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnVirtualRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualRouter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualRouter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnVirtualRouter_PortMappingProperty struct {
	// `CfnVirtualRouter.PortMappingProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnVirtualRouter.PortMappingProperty.Protocol`.
	Protocol *string `json:"protocol"`
}

type CfnVirtualRouter_VirtualRouterListenerProperty struct {
	// `CfnVirtualRouter.VirtualRouterListenerProperty.PortMapping`.
	PortMapping interface{} `json:"portMapping"`
}

type CfnVirtualRouter_VirtualRouterSpecProperty struct {
	// `CfnVirtualRouter.VirtualRouterSpecProperty.Listeners`.
	Listeners interface{} `json:"listeners"`
}

// Properties for defining a `AWS::AppMesh::VirtualRouter`.
type CfnVirtualRouterProps struct {
	// `AWS::AppMesh::VirtualRouter.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::VirtualRouter.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::VirtualRouter.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::VirtualRouter.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::AppMesh::VirtualRouter.VirtualRouterName`.
	VirtualRouterName *string `json:"virtualRouterName"`
}

// A CloudFormation `AWS::AppMesh::VirtualService`.
type CfnVirtualService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrMeshName() *string
	AttrMeshOwner() *string
	AttrResourceOwner() *string
	AttrUid() *string
	AttrVirtualServiceName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MeshName() *string
	SetMeshName(val *string)
	MeshOwner() *string
	SetMeshOwner(val *string)
	Node() constructs.Node
	Ref() *string
	Spec() interface{}
	SetSpec(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VirtualServiceName() *string
	SetVirtualServiceName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVirtualService
type jsiiProxy_CfnVirtualService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVirtualService) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) AttrVirtualServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualService) VirtualServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::VirtualService`.
func NewCfnVirtualService(scope constructs.Construct, id *string, props *CfnVirtualServiceProps) CfnVirtualService {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualService`.
func NewCfnVirtualService_Override(c CfnVirtualService, scope constructs.Construct, id *string, props *CfnVirtualServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualService) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualService) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualService) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualService) SetVirtualServiceName(val *string) {
	_jsii_.Set(
		j,
		"virtualServiceName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnVirtualService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVirtualService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnVirtualService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnVirtualService) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnVirtualService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnVirtualService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnVirtualService_VirtualNodeServiceProviderProperty struct {
	// `CfnVirtualService.VirtualNodeServiceProviderProperty.VirtualNodeName`.
	VirtualNodeName *string `json:"virtualNodeName"`
}

type CfnVirtualService_VirtualRouterServiceProviderProperty struct {
	// `CfnVirtualService.VirtualRouterServiceProviderProperty.VirtualRouterName`.
	VirtualRouterName *string `json:"virtualRouterName"`
}

type CfnVirtualService_VirtualServiceProviderProperty struct {
	// `CfnVirtualService.VirtualServiceProviderProperty.VirtualNode`.
	VirtualNode interface{} `json:"virtualNode"`
	// `CfnVirtualService.VirtualServiceProviderProperty.VirtualRouter`.
	VirtualRouter interface{} `json:"virtualRouter"`
}

type CfnVirtualService_VirtualServiceSpecProperty struct {
	// `CfnVirtualService.VirtualServiceSpecProperty.Provider`.
	Provider interface{} `json:"provider"`
}

// Properties for defining a `AWS::AppMesh::VirtualService`.
type CfnVirtualServiceProps struct {
	// `AWS::AppMesh::VirtualService.MeshName`.
	MeshName *string `json:"meshName"`
	// `AWS::AppMesh::VirtualService.Spec`.
	Spec interface{} `json:"spec"`
	// `AWS::AppMesh::VirtualService.VirtualServiceName`.
	VirtualServiceName *string `json:"virtualServiceName"`
	// `AWS::AppMesh::VirtualService.MeshOwner`.
	MeshOwner *string `json:"meshOwner"`
	// `AWS::AppMesh::VirtualService.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

