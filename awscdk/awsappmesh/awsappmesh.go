package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v3"
)

// Configuration for Envoy Access logs for mesh endpoints.
//
// TODO: EXAMPLE
//
// Experimental.
type AccessLog interface {
	Bind(scope awscdk.Construct) *AccessLogConfig
}

// The jsii proxy struct for AccessLog
type jsiiProxy_AccessLog struct {
	_ byte // padding
}

// Experimental.
func NewAccessLog_Override(a AccessLog) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.AccessLog",
		nil, // no parameters
		a,
	)
}

// Path to a file to write access logs to.
// Experimental.
func AccessLog_FromFilePath(filePath *string) AccessLog {
	_init_.Initialize()

	var returns AccessLog

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.AccessLog",
		"fromFilePath",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Called when the AccessLog type is initialized.
//
// Can be used to enforce
// mutual exclusivity with future properties
// Experimental.
func (a *jsiiProxy_AccessLog) Bind(scope awscdk.Construct) *AccessLogConfig {
	var returns *AccessLogConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// All Properties for Envoy Access logs for mesh endpoints.
// Experimental.
type AccessLogConfig struct {
	// VirtualGateway CFN configuration for Access Logging.
	// Experimental.
	VirtualGatewayAccessLog *CfnVirtualGateway_VirtualGatewayAccessLogProperty `json:"virtualGatewayAccessLog"`
	// VirtualNode CFN configuration for Access Logging.
	// Experimental.
	VirtualNodeAccessLog *CfnVirtualNode_AccessLogProperty `json:"virtualNodeAccessLog"`
}

// Contains static factory methods to create backends.
//
// TODO: EXAMPLE
//
// Experimental.
type Backend interface {
	Bind(_scope awscdk.Construct) *BackendConfig
}

// The jsii proxy struct for Backend
type jsiiProxy_Backend struct {
	_ byte // padding
}

// Experimental.
func NewBackend_Override(b Backend) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.Backend",
		nil, // no parameters
		b,
	)
}

// Construct a Virtual Service backend.
// Experimental.
func Backend_VirtualService(virtualService IVirtualService, props *VirtualServiceBackendOptions) Backend {
	_init_.Initialize()

	var returns Backend

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Backend",
		"virtualService",
		[]interface{}{virtualService, props},
		&returns,
	)

	return returns
}

// Return backend config.
// Experimental.
func (b *jsiiProxy_Backend) Bind(_scope awscdk.Construct) *BackendConfig {
	var returns *BackendConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Properties for a backend.
// Experimental.
type BackendConfig struct {
	// Config for a Virtual Service backend.
	// Experimental.
	VirtualServiceBackend *CfnVirtualNode_BackendProperty `json:"virtualServiceBackend"`
}

// Represents the properties needed to define backend defaults.
//
// TODO: EXAMPLE
//
// Experimental.
type BackendDefaults struct {
	// TLS properties for Client policy for backend defaults.
	// Experimental.
	TlsClientPolicy *TlsClientPolicy `json:"tlsClientPolicy"`
}

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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnGatewayRoute) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnGatewayRoute(scope awscdk.Construct, id *string, props *CfnGatewayRouteProps) CfnGatewayRoute {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayRoute{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnGatewayRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::GatewayRoute`.
func NewCfnGatewayRoute_Override(c CfnGatewayRoute, scope awscdk.Construct, id *string, props *CfnGatewayRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnGatewayRoute",
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
		"monocdk.aws_appmesh.CfnGatewayRoute",
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
		"monocdk.aws_appmesh.CfnGatewayRoute",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGatewayRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnGatewayRoute",
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
		"monocdk.aws_appmesh.CfnGatewayRoute",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnGatewayRoute) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGatewayRoute) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

type CfnGatewayRoute_GatewayRouteHostnameMatchProperty struct {
	// `CfnGatewayRoute.GatewayRouteHostnameMatchProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnGatewayRoute.GatewayRouteHostnameMatchProperty.Suffix`.
	Suffix *string `json:"suffix"`
}

type CfnGatewayRoute_GatewayRouteHostnameRewriteProperty struct {
	// `CfnGatewayRoute.GatewayRouteHostnameRewriteProperty.DefaultTargetHostname`.
	DefaultTargetHostname *string `json:"defaultTargetHostname"`
}

type CfnGatewayRoute_GatewayRouteMetadataMatchProperty struct {
	// `CfnGatewayRoute.GatewayRouteMetadataMatchProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnGatewayRoute.GatewayRouteMetadataMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnGatewayRoute.GatewayRouteMetadataMatchProperty.Range`.
	Range interface{} `json:"range"`
	// `CfnGatewayRoute.GatewayRouteMetadataMatchProperty.Regex`.
	Regex *string `json:"regex"`
	// `CfnGatewayRoute.GatewayRouteMetadataMatchProperty.Suffix`.
	Suffix *string `json:"suffix"`
}

type CfnGatewayRoute_GatewayRouteRangeMatchProperty struct {
	// `CfnGatewayRoute.GatewayRouteRangeMatchProperty.End`.
	End *float64 `json:"end"`
	// `CfnGatewayRoute.GatewayRouteRangeMatchProperty.Start`.
	Start *float64 `json:"start"`
}

type CfnGatewayRoute_GatewayRouteSpecProperty struct {
	// `CfnGatewayRoute.GatewayRouteSpecProperty.GrpcRoute`.
	GrpcRoute interface{} `json:"grpcRoute"`
	// `CfnGatewayRoute.GatewayRouteSpecProperty.Http2Route`.
	Http2Route interface{} `json:"http2Route"`
	// `CfnGatewayRoute.GatewayRouteSpecProperty.HttpRoute`.
	HttpRoute interface{} `json:"httpRoute"`
	// `CfnGatewayRoute.GatewayRouteSpecProperty.Priority`.
	Priority *float64 `json:"priority"`
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
	// `CfnGatewayRoute.GrpcGatewayRouteActionProperty.Rewrite`.
	Rewrite interface{} `json:"rewrite"`
}

type CfnGatewayRoute_GrpcGatewayRouteMatchProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteMatchProperty.Hostname`.
	Hostname interface{} `json:"hostname"`
	// `CfnGatewayRoute.GrpcGatewayRouteMatchProperty.Metadata`.
	Metadata interface{} `json:"metadata"`
	// `CfnGatewayRoute.GrpcGatewayRouteMatchProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
}

type CfnGatewayRoute_GrpcGatewayRouteMetadataProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteMetadataProperty.Name`.
	Name *string `json:"name"`
	// `CfnGatewayRoute.GrpcGatewayRouteMetadataProperty.Invert`.
	Invert interface{} `json:"invert"`
	// `CfnGatewayRoute.GrpcGatewayRouteMetadataProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnGatewayRoute_GrpcGatewayRouteProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnGatewayRoute.GrpcGatewayRouteProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnGatewayRoute_GrpcGatewayRouteRewriteProperty struct {
	// `CfnGatewayRoute.GrpcGatewayRouteRewriteProperty.Hostname`.
	Hostname interface{} `json:"hostname"`
}

type CfnGatewayRoute_HttpGatewayRouteActionProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteActionProperty.Target`.
	Target interface{} `json:"target"`
	// `CfnGatewayRoute.HttpGatewayRouteActionProperty.Rewrite`.
	Rewrite interface{} `json:"rewrite"`
}

type CfnGatewayRoute_HttpGatewayRouteHeaderMatchProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty.Range`.
	Range interface{} `json:"range"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty.Regex`.
	Regex *string `json:"regex"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty.Suffix`.
	Suffix *string `json:"suffix"`
}

type CfnGatewayRoute_HttpGatewayRouteHeaderProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteHeaderProperty.Name`.
	Name *string `json:"name"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderProperty.Invert`.
	Invert interface{} `json:"invert"`
	// `CfnGatewayRoute.HttpGatewayRouteHeaderProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnGatewayRoute_HttpGatewayRouteMatchProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Headers`.
	Headers interface{} `json:"headers"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Hostname`.
	Hostname interface{} `json:"hostname"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Method`.
	Method *string `json:"method"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Path`.
	Path interface{} `json:"path"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.QueryParameters`.
	QueryParameters interface{} `json:"queryParameters"`
}

type CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty struct {
	// `CfnGatewayRoute.HttpGatewayRoutePathRewriteProperty.Exact`.
	Exact *string `json:"exact"`
}

type CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty struct {
	// `CfnGatewayRoute.HttpGatewayRoutePrefixRewriteProperty.DefaultPrefix`.
	DefaultPrefix *string `json:"defaultPrefix"`
	// `CfnGatewayRoute.HttpGatewayRoutePrefixRewriteProperty.Value`.
	Value *string `json:"value"`
}

type CfnGatewayRoute_HttpGatewayRouteProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnGatewayRoute.HttpGatewayRouteProperty.Match`.
	Match interface{} `json:"match"`
}

type CfnGatewayRoute_HttpGatewayRouteRewriteProperty struct {
	// `CfnGatewayRoute.HttpGatewayRouteRewriteProperty.Hostname`.
	Hostname interface{} `json:"hostname"`
	// `CfnGatewayRoute.HttpGatewayRouteRewriteProperty.Path`.
	Path interface{} `json:"path"`
	// `CfnGatewayRoute.HttpGatewayRouteRewriteProperty.Prefix`.
	Prefix interface{} `json:"prefix"`
}

type CfnGatewayRoute_HttpPathMatchProperty struct {
	// `CfnGatewayRoute.HttpPathMatchProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnGatewayRoute.HttpPathMatchProperty.Regex`.
	Regex *string `json:"regex"`
}

type CfnGatewayRoute_HttpQueryParameterMatchProperty struct {
	// `CfnGatewayRoute.HttpQueryParameterMatchProperty.Exact`.
	Exact *string `json:"exact"`
}

type CfnGatewayRoute_QueryParameterProperty struct {
	// `CfnGatewayRoute.QueryParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnGatewayRoute.QueryParameterProperty.Match`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnMesh) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnMesh(scope awscdk.Construct, id *string, props *CfnMeshProps) CfnMesh {
	_init_.Initialize()

	j := jsiiProxy_CfnMesh{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnMesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::Mesh`.
func NewCfnMesh_Override(c CfnMesh, scope awscdk.Construct, id *string, props *CfnMeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnMesh",
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
		"monocdk.aws_appmesh.CfnMesh",
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
		"monocdk.aws_appmesh.CfnMesh",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnMesh",
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
		"monocdk.aws_appmesh.CfnMesh",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnMesh) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMesh) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMesh) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMesh) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMesh) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMesh) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMesh) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnRoute) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnRoute(scope awscdk.Construct, id *string, props *CfnRouteProps) CfnRoute {
	_init_.Initialize()

	j := jsiiProxy_CfnRoute{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::Route`.
func NewCfnRoute_Override(c CfnRoute, scope awscdk.Construct, id *string, props *CfnRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnRoute",
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
		"monocdk.aws_appmesh.CfnRoute",
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
		"monocdk.aws_appmesh.CfnRoute",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnRoute",
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
		"monocdk.aws_appmesh.CfnRoute",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnRoute) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnRoute) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnRoute) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnRoute) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnRoute) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnRoute) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnRoute) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

type CfnRoute_HttpPathMatchProperty struct {
	// `CfnRoute.HttpPathMatchProperty.Exact`.
	Exact *string `json:"exact"`
	// `CfnRoute.HttpPathMatchProperty.Regex`.
	Regex *string `json:"regex"`
}

type CfnRoute_HttpQueryParameterMatchProperty struct {
	// `CfnRoute.HttpQueryParameterMatchProperty.Exact`.
	Exact *string `json:"exact"`
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
	// `CfnRoute.HttpRouteMatchProperty.Headers`.
	Headers interface{} `json:"headers"`
	// `CfnRoute.HttpRouteMatchProperty.Method`.
	Method *string `json:"method"`
	// `CfnRoute.HttpRouteMatchProperty.Path`.
	Path interface{} `json:"path"`
	// `CfnRoute.HttpRouteMatchProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnRoute.HttpRouteMatchProperty.QueryParameters`.
	QueryParameters interface{} `json:"queryParameters"`
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

type CfnRoute_QueryParameterProperty struct {
	// `CfnRoute.QueryParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnRoute.QueryParameterProperty.Match`.
	Match interface{} `json:"match"`
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnVirtualGateway) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnVirtualGateway(scope awscdk.Construct, id *string, props *CfnVirtualGatewayProps) CfnVirtualGateway {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualGateway{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualGateway`.
func NewCfnVirtualGateway_Override(c CfnVirtualGateway, scope awscdk.Construct, id *string, props *CfnVirtualGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualGateway",
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
		"monocdk.aws_appmesh.CfnVirtualGateway",
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
		"monocdk.aws_appmesh.CfnVirtualGateway",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVirtualGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnVirtualGateway",
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
		"monocdk.aws_appmesh.CfnVirtualGateway",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnVirtualGateway) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualGateway) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnVirtualNode) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnVirtualNode(scope awscdk.Construct, id *string, props *CfnVirtualNodeProps) CfnVirtualNode {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualNode{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNode_Override(c CfnVirtualNode, scope awscdk.Construct, id *string, props *CfnVirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualNode",
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
		"monocdk.aws_appmesh.CfnVirtualNode",
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
		"monocdk.aws_appmesh.CfnVirtualNode",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnVirtualNode",
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
		"monocdk.aws_appmesh.CfnVirtualNode",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnVirtualNode) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualNode) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
	// `CfnVirtualNode.DnsServiceDiscoveryProperty.ResponseType`.
	ResponseType *string `json:"responseType"`
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnVirtualRouter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnVirtualRouter(scope awscdk.Construct, id *string, props *CfnVirtualRouterProps) CfnVirtualRouter {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualRouter{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualRouter`.
func NewCfnVirtualRouter_Override(c CfnVirtualRouter, scope awscdk.Construct, id *string, props *CfnVirtualRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualRouter",
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
		"monocdk.aws_appmesh.CfnVirtualRouter",
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
		"monocdk.aws_appmesh.CfnVirtualRouter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVirtualRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnVirtualRouter",
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
		"monocdk.aws_appmesh.CfnVirtualRouter",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnVirtualRouter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualRouter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnVirtualService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnVirtualService(scope awscdk.Construct, id *string, props *CfnVirtualServiceProps) CfnVirtualService {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualService{}

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualService`.
func NewCfnVirtualService_Override(c CfnVirtualService, scope awscdk.Construct, id *string, props *CfnVirtualServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.CfnVirtualService",
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
		"monocdk.aws_appmesh.CfnVirtualService",
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
		"monocdk.aws_appmesh.CfnVirtualService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVirtualService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.CfnVirtualService",
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
		"monocdk.aws_appmesh.CfnVirtualService",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
func (c *jsiiProxy_CfnVirtualService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnVirtualService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

// Enum of DNS service discovery response type.
//
// TODO: EXAMPLE
//
// Experimental.
type DnsResponseType string

const (
	DnsResponseType_LOAD_BALANCER DnsResponseType = "LOAD_BALANCER"
	DnsResponseType_ENDPOINTS DnsResponseType = "ENDPOINTS"
)

// GatewayRoute represents a new or existing gateway route attached to a VirtualGateway and Mesh.
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html
//
// Experimental.
type GatewayRoute interface {
	awscdk.Resource
	IGatewayRoute
	Env() *awscdk.ResourceEnvironment
	GatewayRouteArn() *string
	GatewayRouteName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualGateway() IVirtualGateway
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for GatewayRoute
type jsiiProxy_GatewayRoute struct {
	internal.Type__awscdkResource
	jsiiProxy_IGatewayRoute
}

func (j *jsiiProxy_GatewayRoute) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) GatewayRouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) GatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayRoute) VirtualGateway() IVirtualGateway {
	var returns IVirtualGateway
	_jsii_.Get(
		j,
		"virtualGateway",
		&returns,
	)
	return returns
}


// Experimental.
func NewGatewayRoute(scope constructs.Construct, id *string, props *GatewayRouteProps) GatewayRoute {
	_init_.Initialize()

	j := jsiiProxy_GatewayRoute{}

	_jsii_.Create(
		"monocdk.aws_appmesh.GatewayRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGatewayRoute_Override(g GatewayRoute, scope constructs.Construct, id *string, props *GatewayRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.GatewayRoute",
		[]interface{}{scope, id, props},
		g,
	)
}

// Import an existing GatewayRoute given an ARN.
// Experimental.
func GatewayRoute_FromGatewayRouteArn(scope constructs.Construct, id *string, gatewayRouteArn *string) IGatewayRoute {
	_init_.Initialize()

	var returns IGatewayRoute

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRoute",
		"fromGatewayRouteArn",
		[]interface{}{scope, id, gatewayRouteArn},
		&returns,
	)

	return returns
}

// Import an existing GatewayRoute given attributes.
// Experimental.
func GatewayRoute_FromGatewayRouteAttributes(scope constructs.Construct, id *string, attrs *GatewayRouteAttributes) IGatewayRoute {
	_init_.Initialize()

	var returns IGatewayRoute

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRoute",
		"fromGatewayRouteAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func GatewayRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GatewayRoute_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRoute",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (g *jsiiProxy_GatewayRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (g *jsiiProxy_GatewayRoute) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (g *jsiiProxy_GatewayRoute) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (g *jsiiProxy_GatewayRoute) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GatewayRoute) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GatewayRoute) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GatewayRoute) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GatewayRoute) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GatewayRoute) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GatewayRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GatewayRoute) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties necessary to import a reusable GatewayRoute.
// Experimental.
type GatewayRouteAttributes struct {
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `json:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	// Experimental.
	VirtualGateway IVirtualGateway `json:"virtualGateway"`
}

// Basic configuration properties for a GatewayRoute.
//
// TODO: EXAMPLE
//
// Experimental.
type GatewayRouteBaseProps struct {
	// What protocol the route uses.
	// Experimental.
	RouteSpec GatewayRouteSpec `json:"routeSpec"`
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `json:"gatewayRouteName"`
}

// Used to generate host name matching methods.
//
// TODO: EXAMPLE
//
// Experimental.
type GatewayRouteHostnameMatch interface {
	Bind(scope awscdk.Construct) *GatewayRouteHostnameMatchConfig
}

// The jsii proxy struct for GatewayRouteHostnameMatch
type jsiiProxy_GatewayRouteHostnameMatch struct {
	_ byte // padding
}

// Experimental.
func NewGatewayRouteHostnameMatch_Override(g GatewayRouteHostnameMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		nil, // no parameters
		g,
	)
}

// The value of the host name with the given name must end with the specified characters.
// Experimental.
func GatewayRouteHostnameMatch_EndsWith(suffix *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		"endsWith",
		[]interface{}{suffix},
		&returns,
	)

	return returns
}

// The value of the host name must match the specified value exactly.
// Experimental.
func GatewayRouteHostnameMatch_Exactly(name *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		"exactly",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the gateway route host name match configuration.
// Experimental.
func (g *jsiiProxy_GatewayRouteHostnameMatch) Bind(scope awscdk.Construct) *GatewayRouteHostnameMatchConfig {
	var returns *GatewayRouteHostnameMatchConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Configuration for gateway route host name match.
// Experimental.
type GatewayRouteHostnameMatchConfig struct {
	// GatewayRoute CFN configuration for host name match.
	// Experimental.
	HostnameMatch *CfnGatewayRoute_GatewayRouteHostnameMatchProperty `json:"hostnameMatch"`
}

// Properties to define a new GatewayRoute.
// Experimental.
type GatewayRouteProps struct {
	// What protocol the route uses.
	// Experimental.
	RouteSpec GatewayRouteSpec `json:"routeSpec"`
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `json:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	// Experimental.
	VirtualGateway IVirtualGateway `json:"virtualGateway"`
}

// Used to generate specs with different protocols for a GatewayRoute.
//
// TODO: EXAMPLE
//
// Experimental.
type GatewayRouteSpec interface {
	Bind(scope awscdk.Construct) *GatewayRouteSpecConfig
}

// The jsii proxy struct for GatewayRouteSpec
type jsiiProxy_GatewayRouteSpec struct {
	_ byte // padding
}

// Experimental.
func NewGatewayRouteSpec_Override(g GatewayRouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.GatewayRouteSpec",
		nil, // no parameters
		g,
	)
}

// Creates an gRPC Based GatewayRoute.
// Experimental.
func GatewayRouteSpec_Grpc(options *GrpcGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based GatewayRoute.
// Experimental.
func GatewayRouteSpec_Http(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based GatewayRoute.
// Experimental.
func GatewayRouteSpec_Http2(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteSpec",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Called when the GatewayRouteSpec type is initialized.
//
// Can be used to enforce
// mutual exclusivity with future properties
// Experimental.
func (g *jsiiProxy_GatewayRouteSpec) Bind(scope awscdk.Construct) *GatewayRouteSpecConfig {
	var returns *GatewayRouteSpecConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// All Properties for GatewayRoute Specs.
// Experimental.
type GatewayRouteSpecConfig struct {
	// The spec for a grpc gateway route.
	// Experimental.
	GrpcSpecConfig *CfnGatewayRoute_GrpcGatewayRouteProperty `json:"grpcSpecConfig"`
	// The spec for an http2 gateway route.
	// Experimental.
	Http2SpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `json:"http2SpecConfig"`
	// The spec for an http gateway route.
	// Experimental.
	HttpSpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `json:"httpSpecConfig"`
}

// Connection pool properties for gRPC listeners.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcConnectionPool struct {
	// The maximum requests in the pool.
	// Experimental.
	MaxRequests *float64 `json:"maxRequests"`
}

// Represents the properties needed to define GRPC Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcGatewayListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *GrpcConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// The criterion for determining a request match for this GatewayRoute.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcGatewayRouteMatch struct {
	// Create host name based gRPC gateway route match.
	// Experimental.
	Hostname GatewayRouteHostnameMatch `json:"hostname"`
	// Create metadata based gRPC gateway route match.
	//
	// All specified metadata must match for the route to match.
	// Experimental.
	Metadata *[]HeaderMatch `json:"metadata"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	// Experimental.
	RewriteRequestHostname *bool `json:"rewriteRequestHostname"`
	// Create service name based gRPC gateway route match.
	// Experimental.
	ServiceName *string `json:"serviceName"`
}

// Properties specific for a gRPC GatewayRoute.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcGatewayRouteSpecOptions struct {
	// The criterion for determining a request match for this GatewayRoute.
	// Experimental.
	Match *GrpcGatewayRouteMatch `json:"match"`
	// The VirtualService this GatewayRoute directs traffic to.
	// Experimental.
	RouteTarget IVirtualService `json:"routeTarget"`
}

// Properties used to define GRPC Based healthchecks.
// Experimental.
type GrpcHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// Experimental.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// The time period between each health check execution.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

// gRPC events.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcRetryEvent string

const (
	GrpcRetryEvent_CANCELLED GrpcRetryEvent = "CANCELLED"
	GrpcRetryEvent_DEADLINE_EXCEEDED GrpcRetryEvent = "DEADLINE_EXCEEDED"
	GrpcRetryEvent_INTERNAL_ERROR GrpcRetryEvent = "INTERNAL_ERROR"
	GrpcRetryEvent_RESOURCE_EXHAUSTED GrpcRetryEvent = "RESOURCE_EXHAUSTED"
	GrpcRetryEvent_UNAVAILABLE GrpcRetryEvent = "UNAVAILABLE"
)

// gRPC retry policy.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcRetryPolicy struct {
	// The maximum number of retry attempts.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// The timeout for each retry attempt.
	// Experimental.
	RetryTimeout awscdk.Duration `json:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Experimental.
	HttpRetryEvents *[]HttpRetryEvent `json:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	// Experimental.
	TcpRetryEvents *[]TcpRetryEvent `json:"tcpRetryEvents"`
	// gRPC events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Experimental.
	GrpcRetryEvents *[]GrpcRetryEvent `json:"grpcRetryEvents"`
}

// The criterion for determining a request match for this Route.
//
// At least one match type must be selected.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcRouteMatch struct {
	// Create metadata based gRPC route match.
	//
	// All specified metadata must match for the route to match.
	// Experimental.
	Metadata *[]HeaderMatch `json:"metadata"`
	// The method name to match from the request.
	//
	// If the method name is specified, service name must be also provided.
	// Experimental.
	MethodName *string `json:"methodName"`
	// Create service name based gRPC route match.
	// Experimental.
	ServiceName *string `json:"serviceName"`
}

// Properties specific for a GRPC Based Routes.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcRouteSpecOptions struct {
	// The priority for the route.
	//
	// Routes are matched based on the specified
	// value, where 0 is the highest priority.
	// Experimental.
	Priority *float64 `json:"priority"`
	// The criterion for determining a request match for this Route.
	// Experimental.
	Match *GrpcRouteMatch `json:"match"`
	// List of targets that traffic is routed to when a request matches the route.
	// Experimental.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets"`
	// The retry policy.
	// Experimental.
	RetryPolicy *GrpcRetryPolicy `json:"retryPolicy"`
	// An object that represents a grpc timeout.
	// Experimental.
	Timeout *GrpcTimeout `json:"timeout"`
}

// Represents timeouts for GRPC protocols.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Experimental.
	Idle awscdk.Duration `json:"idle"`
	// Represents per request timeout.
	// Experimental.
	PerRequest awscdk.Duration `json:"perRequest"`
}

// Represent the GRPC Node Listener prorperty.
//
// TODO: EXAMPLE
//
// Experimental.
type GrpcVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *GrpcConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `json:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Timeout for GRPC protocol.
	// Experimental.
	Timeout *GrpcTimeout `json:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// Used to generate header matching methods.
//
// TODO: EXAMPLE
//
// Experimental.
type HeaderMatch interface {
	Bind(scope awscdk.Construct) *HeaderMatchConfig
}

// The jsii proxy struct for HeaderMatch
type jsiiProxy_HeaderMatch struct {
	_ byte // padding
}

// Experimental.
func NewHeaderMatch_Override(h HeaderMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HeaderMatch",
		nil, // no parameters
		h,
	)
}

// The value of the header with the given name in the request must not end with the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotEndWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotEndWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not include the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotMatchRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotMatchRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not start with the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotStartWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotStartWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must end with the specified characters.
// Experimental.
func HeaderMatch_ValueEndsWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueEndsWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must match the specified value exactly.
// Experimental.
func HeaderMatch_ValueIs(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueIs",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not match the specified value exactly.
// Experimental.
func HeaderMatch_ValueIsNot(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueIsNot",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must include the specified characters.
// Experimental.
func HeaderMatch_ValueMatchesRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueMatchesRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must be in a range of values.
// Experimental.
func HeaderMatch_ValuesIsInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valuesIsInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not be in a range of values.
// Experimental.
func HeaderMatch_ValuesIsNotInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valuesIsNotInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must start with the specified characters.
// Experimental.
func HeaderMatch_ValueStartsWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueStartsWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// Returns the header match configuration.
// Experimental.
func (h *jsiiProxy_HeaderMatch) Bind(scope awscdk.Construct) *HeaderMatchConfig {
	var returns *HeaderMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Configuration for `HeaderMatch`.
// Experimental.
type HeaderMatchConfig struct {
	// Route CFN configuration for the route header match.
	// Experimental.
	HeaderMatch *CfnRoute_HttpRouteHeaderProperty `json:"headerMatch"`
}

// Contains static factory methods for creating health checks for different protocols.
//
// TODO: EXAMPLE
//
// Experimental.
type HealthCheck interface {
	Bind(scope awscdk.Construct, options *HealthCheckBindOptions) *HealthCheckConfig
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

// Experimental.
func NewHealthCheck_Override(h HealthCheck) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HealthCheck",
		nil, // no parameters
		h,
	)
}

// Construct a GRPC health check.
// Experimental.
func HealthCheck_Grpc(options *GrpcHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP health check.
// Experimental.
func HealthCheck_Http(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP2 health check.
// Experimental.
func HealthCheck_Http2(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a TCP health check.
// Experimental.
func HealthCheck_Tcp(options *TcpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"tcp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Called when the AccessLog type is initialized.
//
// Can be used to enforce
// mutual exclusivity with future properties
// Experimental.
func (h *jsiiProxy_HealthCheck) Bind(scope awscdk.Construct, options *HealthCheckBindOptions) *HealthCheckConfig {
	var returns *HealthCheckConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Options used for creating the Health Check object.
// Experimental.
type HealthCheckBindOptions struct {
	// Port for Health Check interface.
	// Experimental.
	DefaultPort *float64 `json:"defaultPort"`
}

// All Properties for Health Checks for mesh endpoints.
// Experimental.
type HealthCheckConfig struct {
	// VirtualGateway CFN configuration for Health Checks.
	// Experimental.
	VirtualGatewayHealthCheck *CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty `json:"virtualGatewayHealthCheck"`
	// VirtualNode CFN configuration for Health Checks.
	// Experimental.
	VirtualNodeHealthCheck *CfnVirtualNode_HealthCheckProperty `json:"virtualNodeHealthCheck"`
}

// Connection pool properties for HTTP2 listeners.
// Experimental.
type Http2ConnectionPool struct {
	// The maximum requests in the pool.
	// Experimental.
	MaxRequests *float64 `json:"maxRequests"`
}

// Represents the properties needed to define HTTP2 Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type Http2GatewayListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *Http2ConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// Represent the HTTP2 Node Listener prorperty.
// Experimental.
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	// Experimental.
	ConnectionPool *Http2ConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `json:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Timeout for HTTP protocol.
	// Experimental.
	Timeout *HttpTimeout `json:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// Connection pool properties for HTTP listeners.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpConnectionPool struct {
	// The maximum connections in the pool.
	// Experimental.
	MaxConnections *float64 `json:"maxConnections"`
	// The maximum pending requests in the pool.
	// Experimental.
	MaxPendingRequests *float64 `json:"maxPendingRequests"`
}

// Represents the properties needed to define HTTP Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpGatewayListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *HttpConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// The criterion for determining a request match for this GatewayRoute.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpGatewayRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the gateway route to match.
	// Experimental.
	Headers *[]HeaderMatch `json:"headers"`
	// The gateway route host name to be matched on.
	// Experimental.
	Hostname GatewayRouteHostnameMatch `json:"hostname"`
	// The method to match on.
	// Experimental.
	Method HttpRouteMethod `json:"method"`
	// Specify how to match requests based on the 'path' part of their URL.
	// Experimental.
	Path HttpGatewayRoutePathMatch `json:"path"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	// Experimental.
	QueryParameters *[]QueryParameterMatch `json:"queryParameters"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	// Experimental.
	RewriteRequestHostname *bool `json:"rewriteRequestHostname"`
}

// Defines HTTP gateway route matching based on the URL path of the request.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpGatewayRoutePathMatch interface {
	Bind(scope awscdk.Construct) *HttpGatewayRoutePathMatchConfig
}

// The jsii proxy struct for HttpGatewayRoutePathMatch
type jsiiProxy_HttpGatewayRoutePathMatch struct {
	_ byte // padding
}

// Experimental.
func NewHttpGatewayRoutePathMatch_Override(h HttpGatewayRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
// Experimental.
func HttpGatewayRoutePathMatch_Exactly(path *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"exactly",
		[]interface{}{path, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
// Experimental.
func HttpGatewayRoutePathMatch_Regex(regex *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"regex",
		[]interface{}{regex, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
// Experimental.
func HttpGatewayRoutePathMatch_StartsWith(prefix *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"startsWith",
		[]interface{}{prefix, rewriteTo},
		&returns,
	)

	return returns
}

// Returns the gateway route path match configuration.
// Experimental.
func (h *jsiiProxy_HttpGatewayRoutePathMatch) Bind(scope awscdk.Construct) *HttpGatewayRoutePathMatchConfig {
	var returns *HttpGatewayRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// The type returned from the `bind()` method in {@link HttpGatewayRoutePathMatch}.
// Experimental.
type HttpGatewayRoutePathMatchConfig struct {
	// Gateway route configuration for matching on the prefix of the URL path of the request.
	// Experimental.
	PrefixPathMatch *string `json:"prefixPathMatch"`
	// Gateway route configuration for rewriting the prefix of the URL path of the request.
	// Experimental.
	PrefixPathRewrite *CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty `json:"prefixPathRewrite"`
	// Gateway route configuration for matching on the complete URL path of the request.
	// Experimental.
	WholePathMatch *CfnGatewayRoute_HttpPathMatchProperty `json:"wholePathMatch"`
	// Gateway route configuration for rewriting the complete URL path of the request..
	// Experimental.
	WholePathRewrite *CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty `json:"wholePathRewrite"`
}

// Properties specific for HTTP Based GatewayRoutes.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpGatewayRouteSpecOptions struct {
	// The VirtualService this GatewayRoute directs traffic to.
	// Experimental.
	RouteTarget IVirtualService `json:"routeTarget"`
	// The criterion for determining a request match for this GatewayRoute.
	//
	// When path match is defined, this may optionally determine the path rewrite configuration.
	// Experimental.
	Match *HttpGatewayRouteMatch `json:"match"`
}

// Properties used to define HTTP Based healthchecks.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// Experimental.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// The time period between each health check execution.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// The destination path for the health check request.
	// Experimental.
	Path *string `json:"path"`
	// The amount of time to wait when receiving a response from the health check.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

// HTTP events on which to retry.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRetryEvent string

const (
	HttpRetryEvent_SERVER_ERROR HttpRetryEvent = "SERVER_ERROR"
	HttpRetryEvent_GATEWAY_ERROR HttpRetryEvent = "GATEWAY_ERROR"
	HttpRetryEvent_CLIENT_ERROR HttpRetryEvent = "CLIENT_ERROR"
	HttpRetryEvent_STREAM_ERROR HttpRetryEvent = "STREAM_ERROR"
)

// HTTP retry policy.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRetryPolicy struct {
	// The maximum number of retry attempts.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// The timeout for each retry attempt.
	// Experimental.
	RetryTimeout awscdk.Duration `json:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Experimental.
	HttpRetryEvents *[]HttpRetryEvent `json:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	// Experimental.
	TcpRetryEvents *[]TcpRetryEvent `json:"tcpRetryEvents"`
}

// The criterion for determining a request match for this Route.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the route to match.
	// Experimental.
	Headers *[]HeaderMatch `json:"headers"`
	// The HTTP client request method to match on.
	// Experimental.
	Method HttpRouteMethod `json:"method"`
	// Specifies how is the request matched based on the path part of its URL.
	// Experimental.
	Path HttpRoutePathMatch `json:"path"`
	// The client request protocol to match on.
	//
	// Applicable only for HTTP2 routes.
	// Experimental.
	Protocol HttpRouteProtocol `json:"protocol"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	// Experimental.
	QueryParameters *[]QueryParameterMatch `json:"queryParameters"`
}

// Supported values for matching routes based on the HTTP request method.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRouteMethod string

const (
	HttpRouteMethod_GET HttpRouteMethod = "GET"
	HttpRouteMethod_HEAD HttpRouteMethod = "HEAD"
	HttpRouteMethod_POST HttpRouteMethod = "POST"
	HttpRouteMethod_PUT HttpRouteMethod = "PUT"
	HttpRouteMethod_DELETE HttpRouteMethod = "DELETE"
	HttpRouteMethod_CONNECT HttpRouteMethod = "CONNECT"
	HttpRouteMethod_OPTIONS HttpRouteMethod = "OPTIONS"
	HttpRouteMethod_TRACE HttpRouteMethod = "TRACE"
	HttpRouteMethod_PATCH HttpRouteMethod = "PATCH"
)

// Defines HTTP route matching based on the URL path of the request.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRoutePathMatch interface {
	Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig
}

// The jsii proxy struct for HttpRoutePathMatch
type jsiiProxy_HttpRoutePathMatch struct {
	_ byte // padding
}

// Experimental.
func NewHttpRoutePathMatch_Override(h HttpRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
// Experimental.
func HttpRoutePathMatch_Exactly(path *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"exactly",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
// Experimental.
func HttpRoutePathMatch_Regex(regex *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"regex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
// Experimental.
func HttpRoutePathMatch_StartsWith(prefix *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"startsWith",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the route path match configuration.
// Experimental.
func (h *jsiiProxy_HttpRoutePathMatch) Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig {
	var returns *HttpRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// The type returned from the `bind()` method in {@link HttpRoutePathMatch}.
// Experimental.
type HttpRoutePathMatchConfig struct {
	// Route configuration for matching on the prefix of the URL path of the request.
	// Experimental.
	PrefixPathMatch *string `json:"prefixPathMatch"`
	// Route configuration for matching on the complete URL path of the request.
	// Experimental.
	WholePathMatch *CfnRoute_HttpPathMatchProperty `json:"wholePathMatch"`
}

// Supported :scheme options for HTTP2.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRouteProtocol string

const (
	HttpRouteProtocol_HTTP HttpRouteProtocol = "HTTP"
	HttpRouteProtocol_HTTPS HttpRouteProtocol = "HTTPS"
)

// Properties specific for HTTP Based Routes.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpRouteSpecOptions struct {
	// The priority for the route.
	//
	// Routes are matched based on the specified
	// value, where 0 is the highest priority.
	// Experimental.
	Priority *float64 `json:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	// Experimental.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets"`
	// The criterion for determining a request match for this Route.
	// Experimental.
	Match *HttpRouteMatch `json:"match"`
	// The retry policy.
	// Experimental.
	RetryPolicy *HttpRetryPolicy `json:"retryPolicy"`
	// An object that represents a http timeout.
	// Experimental.
	Timeout *HttpTimeout `json:"timeout"`
}

// Represents timeouts for HTTP protocols.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Experimental.
	Idle awscdk.Duration `json:"idle"`
	// Represents per request timeout.
	// Experimental.
	PerRequest awscdk.Duration `json:"perRequest"`
}

// Represent the HTTP Node Listener prorperty.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *HttpConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `json:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Timeout for HTTP protocol.
	// Experimental.
	Timeout *HttpTimeout `json:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// Interface for which all GatewayRoute based classes MUST implement.
// Experimental.
type IGatewayRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the GatewayRoute.
	// Experimental.
	GatewayRouteArn() *string
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName() *string
	// The VirtualGateway the GatewayRoute belongs to.
	// Experimental.
	VirtualGateway() IVirtualGateway
}

// The jsii proxy for IGatewayRoute
type jsiiProxy_IGatewayRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) VirtualGateway() IVirtualGateway {
	var returns IVirtualGateway
	_jsii_.Get(
		j,
		"virtualGateway",
		&returns,
	)
	return returns
}

// Interface which all Mesh based classes MUST implement.
// Experimental.
type IMesh interface {
	awscdk.IResource
	// Creates a new VirtualGateway in this Mesh.
	//
	// Note that the Gateway is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	// Experimental.
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	// Creates a new VirtualNode in this Mesh.
	//
	// Note that the Node is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	// Experimental.
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	// Creates a new VirtualRouter in this Mesh.
	//
	// Note that the Router is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	// Experimental.
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
	// The Amazon Resource Name (ARN) of the AppMesh mesh.
	// Experimental.
	MeshArn() *string
	// The name of the AppMesh mesh.
	// Experimental.
	MeshName() *string
}

// The jsii proxy for IMesh
type jsiiProxy_IMesh struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IMesh) AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway {
	var returns VirtualGateway

	_jsii_.Invoke(
		i,
		"addVirtualGateway",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMesh) AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode {
	var returns VirtualNode

	_jsii_.Invoke(
		i,
		"addVirtualNode",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMesh) AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter {
	var returns VirtualRouter

	_jsii_.Invoke(
		i,
		"addVirtualRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMesh) MeshArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMesh) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

// Interface for which all Route based classes MUST implement.
// Experimental.
type IRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the route.
	// Experimental.
	RouteArn() *string
	// The name of the route.
	// Experimental.
	RouteName() *string
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter() IVirtualRouter
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRoute) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) VirtualRouter() IVirtualRouter {
	var returns IVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}

// Interface which all Virtual Gateway based classes must implement.
// Experimental.
type IVirtualGateway interface {
	awscdk.IResource
	// Utility method to add a new GatewayRoute to the VirtualGateway.
	// Experimental.
	AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	// Experimental.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualGateway belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualGateway.
	// Experimental.
	VirtualGatewayArn() *string
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName() *string
}

// The jsii proxy for IVirtualGateway
type jsiiProxy_IVirtualGateway struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVirtualGateway) AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute {
	var returns GatewayRoute

	_jsii_.Invoke(
		i,
		"addGatewayRoute",
		[]interface{}{id, route},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVirtualGateway) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualGateway) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}

// Interface which all VirtualNode based classes must implement.
// Experimental.
type IVirtualNode interface {
	awscdk.IResource
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	// Experimental.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualNode belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name belonging to the VirtualNode.
	//
	// Set this value as the APPMESH_VIRTUAL_NODE_NAME environment variable for
	// your task group's Envoy proxy container in your task definition or pod
	// spec.
	// Experimental.
	VirtualNodeArn() *string
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName() *string
}

// The jsii proxy for IVirtualNode
type jsiiProxy_IVirtualNode struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVirtualNode) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualNode) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) VirtualNodeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) VirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeName",
		&returns,
	)
	return returns
}

// Interface which all VirtualRouter based classes MUST implement.
// Experimental.
type IVirtualRouter interface {
	awscdk.IResource
	// Add a single route to the router.
	// Experimental.
	AddRoute(id *string, props *RouteBaseProps) Route
	// The Mesh which the VirtualRouter belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualRouter.
	// Experimental.
	VirtualRouterArn() *string
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName() *string
}

// The jsii proxy for IVirtualRouter
type jsiiProxy_IVirtualRouter struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVirtualRouter) AddRoute(id *string, props *RouteBaseProps) Route {
	var returns Route

	_jsii_.Invoke(
		i,
		"addRoute",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualRouter) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) VirtualRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}

// Represents the interface which all VirtualService based classes MUST implement.
// Experimental.
type IVirtualService interface {
	awscdk.IResource
	// The Mesh which the VirtualService belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the virtual service.
	// Experimental.
	VirtualServiceArn() *string
	// The name of the VirtualService.
	// Experimental.
	VirtualServiceName() *string
}

// The jsii proxy for IVirtualService
type jsiiProxy_IVirtualService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVirtualService) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualService) VirtualServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualService) VirtualServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceName",
		&returns,
	)
	return returns
}

// Represents TLS properties for listener.
//
// TODO: EXAMPLE
//
// Experimental.
type ListenerTlsOptions struct {
	// Represents TLS certificate.
	// Experimental.
	Certificate TlsCertificate `json:"certificate"`
	// The TLS mode.
	// Experimental.
	Mode TlsMode `json:"mode"`
	// Represents a listener's TLS validation context.
	//
	// The client certificate will only be validated if the client provides it, enabling mutual TLS.
	// Experimental.
	MutualTlsValidation *MutualTlsValidation `json:"mutualTlsValidation"`
}

// Define a new AppMesh mesh.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html
//
// Experimental.
type Mesh interface {
	awscdk.Resource
	IMesh
	Env() *awscdk.ResourceEnvironment
	MeshArn() *string
	MeshName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Mesh
type jsiiProxy_Mesh struct {
	internal.Type__awscdkResource
	jsiiProxy_IMesh
}

func (j *jsiiProxy_Mesh) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) MeshArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewMesh(scope constructs.Construct, id *string, props *MeshProps) Mesh {
	_init_.Initialize()

	j := jsiiProxy_Mesh{}

	_jsii_.Create(
		"monocdk.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMesh_Override(m Mesh, scope constructs.Construct, id *string, props *MeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import an existing mesh by arn.
// Experimental.
func Mesh_FromMeshArn(scope constructs.Construct, id *string, meshArn *string) IMesh {
	_init_.Initialize()

	var returns IMesh

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Mesh",
		"fromMeshArn",
		[]interface{}{scope, id, meshArn},
		&returns,
	)

	return returns
}

// Import an existing mesh by name.
// Experimental.
func Mesh_FromMeshName(scope constructs.Construct, id *string, meshName *string) IMesh {
	_init_.Initialize()

	var returns IMesh

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Mesh",
		"fromMeshName",
		[]interface{}{scope, id, meshName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Mesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Mesh",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Mesh_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Mesh",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a VirtualGateway to the Mesh.
// Experimental.
func (m *jsiiProxy_Mesh) AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway {
	var returns VirtualGateway

	_jsii_.Invoke(
		m,
		"addVirtualGateway",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds a VirtualNode to the Mesh.
// Experimental.
func (m *jsiiProxy_Mesh) AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode {
	var returns VirtualNode

	_jsii_.Invoke(
		m,
		"addVirtualNode",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds a VirtualRouter to the Mesh with the given id and props.
// Experimental.
func (m *jsiiProxy_Mesh) AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter {
	var returns VirtualRouter

	_jsii_.Invoke(
		m,
		"addVirtualRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

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
func (m *jsiiProxy_Mesh) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (m *jsiiProxy_Mesh) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (m *jsiiProxy_Mesh) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (m *jsiiProxy_Mesh) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (m *jsiiProxy_Mesh) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (m *jsiiProxy_Mesh) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (m *jsiiProxy_Mesh) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (m *jsiiProxy_Mesh) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (m *jsiiProxy_Mesh) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (m *jsiiProxy_Mesh) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (m *jsiiProxy_Mesh) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A utility enum defined for the egressFilter type property, the default of DROP_ALL, allows traffic only to other resources inside the mesh, or API calls to amazon resources.
//
// TODO: EXAMPLE
//
// Experimental.
type MeshFilterType string

const (
	MeshFilterType_ALLOW_ALL MeshFilterType = "ALLOW_ALL"
	MeshFilterType_DROP_ALL MeshFilterType = "DROP_ALL"
)

// The set of properties used when creating a Mesh.
//
// TODO: EXAMPLE
//
// Experimental.
type MeshProps struct {
	// Egress filter to be applied to the Mesh.
	// Experimental.
	EgressFilter MeshFilterType `json:"egressFilter"`
	// The name of the Mesh being defined.
	// Experimental.
	MeshName *string `json:"meshName"`
}

// Represents a TLS certificate that is supported for mutual TLS authentication.
//
// TODO: EXAMPLE
//
// Experimental.
type MutualTlsCertificate interface {
	TlsCertificate
	Differentiator() *bool
	Bind(_scope awscdk.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for MutualTlsCertificate
type jsiiProxy_MutualTlsCertificate struct {
	jsiiProxy_TlsCertificate
}

func (j *jsiiProxy_MutualTlsCertificate) Differentiator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"differentiator",
		&returns,
	)
	return returns
}


// Experimental.
func NewMutualTlsCertificate_Override(m MutualTlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.MutualTlsCertificate",
		nil, // no parameters
		m,
	)
}

// Returns an ACM TLS Certificate.
// Experimental.
func MutualTlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
// Experimental.
func MutualTlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
// Experimental.
func MutualTlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns TLS certificate based provider.
// Experimental.
func (m *jsiiProxy_MutualTlsCertificate) Bind(_scope awscdk.Construct) *TlsCertificateConfig {
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the properties needed to define TLS Validation context that is supported for mutual TLS authentication.
//
// TODO: EXAMPLE
//
// Experimental.
type MutualTlsValidation struct {
	// Reference to where to retrieve the trust chain.
	// Experimental.
	Trust MutualTlsValidationTrust `json:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	// Experimental.
	SubjectAlternativeNames SubjectAlternativeNames `json:"subjectAlternativeNames"`
}

// Represents a TLS Validation Context Trust that is supported for mutual TLS authentication.
//
// TODO: EXAMPLE
//
// Experimental.
type MutualTlsValidationTrust interface {
	TlsValidationTrust
	Differentiator() *bool
	Bind(scope awscdk.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for MutualTlsValidationTrust
type jsiiProxy_MutualTlsValidationTrust struct {
	jsiiProxy_TlsValidationTrust
}

func (j *jsiiProxy_MutualTlsValidationTrust) Differentiator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"differentiator",
		&returns,
	)
	return returns
}


// Experimental.
func NewMutualTlsValidationTrust_Override(m MutualTlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.MutualTlsValidationTrust",
		nil, // no parameters
		m,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
// Experimental.
func MutualTlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
// Experimental.
func MutualTlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
// Experimental.
func MutualTlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.MutualTlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns Trust context based on trust type.
// Experimental.
func (m *jsiiProxy_MutualTlsValidationTrust) Bind(scope awscdk.Construct) *TlsValidationTrustConfig {
	var returns *TlsValidationTrustConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Represents the outlier detection for a listener.
//
// TODO: EXAMPLE
//
// Experimental.
type OutlierDetection struct {
	// The base amount of time for which a host is ejected.
	// Experimental.
	BaseEjectionDuration awscdk.Duration `json:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at
	// least one host regardless of the value.
	// Experimental.
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	// Number of consecutive 5xx errors required for ejection.
	// Experimental.
	MaxServerErrors *float64 `json:"maxServerErrors"`
}

// Enum of supported AppMesh protocols.
// Deprecated: not for use outside package
type Protocol string

const (
	Protocol_HTTP Protocol = "HTTP"
	Protocol_TCP Protocol = "TCP"
	Protocol_HTTP2 Protocol = "HTTP2"
	Protocol_GRPC Protocol = "GRPC"
)

// Used to generate query parameter matching methods.
//
// TODO: EXAMPLE
//
// Experimental.
type QueryParameterMatch interface {
	Bind(scope awscdk.Construct) *QueryParameterMatchConfig
}

// The jsii proxy struct for QueryParameterMatch
type jsiiProxy_QueryParameterMatch struct {
	_ byte // padding
}

// Experimental.
func NewQueryParameterMatch_Override(q QueryParameterMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.QueryParameterMatch",
		nil, // no parameters
		q,
	)
}

// The value of the query parameter with the given name in the request must match the specified value exactly.
// Experimental.
func QueryParameterMatch_ValueIs(queryParameterName *string, queryParameterValue *string) QueryParameterMatch {
	_init_.Initialize()

	var returns QueryParameterMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.QueryParameterMatch",
		"valueIs",
		[]interface{}{queryParameterName, queryParameterValue},
		&returns,
	)

	return returns
}

// Returns the query parameter match configuration.
// Experimental.
func (q *jsiiProxy_QueryParameterMatch) Bind(scope awscdk.Construct) *QueryParameterMatchConfig {
	var returns *QueryParameterMatchConfig

	_jsii_.Invoke(
		q,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Configuration for `QueryParameterMatch`.
// Experimental.
type QueryParameterMatchConfig struct {
	// Route CFN configuration for route query parameter match.
	// Experimental.
	QueryParameterMatch *CfnRoute_QueryParameterProperty `json:"queryParameterMatch"`
}

// Route represents a new or existing route attached to a VirtualRouter and Mesh.
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html
//
// Experimental.
type Route interface {
	awscdk.Resource
	IRoute
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RouteArn() *string
	RouteName() *string
	Stack() awscdk.Stack
	VirtualRouter() IVirtualRouter
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Route
type jsiiProxy_Route struct {
	internal.Type__awscdkResource
	jsiiProxy_IRoute
}

func (j *jsiiProxy_Route) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) VirtualRouter() IVirtualRouter {
	var returns IVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}


// Experimental.
func NewRoute(scope constructs.Construct, id *string, props *RouteProps) Route {
	_init_.Initialize()

	j := jsiiProxy_Route{}

	_jsii_.Create(
		"monocdk.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRoute_Override(r Route, scope constructs.Construct, id *string, props *RouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing Route given an ARN.
// Experimental.
func Route_FromRouteArn(scope constructs.Construct, id *string, routeArn *string) IRoute {
	_init_.Initialize()

	var returns IRoute

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Route",
		"fromRouteArn",
		[]interface{}{scope, id, routeArn},
		&returns,
	)

	return returns
}

// Import an existing Route given attributes.
// Experimental.
func Route_FromRouteAttributes(scope constructs.Construct, id *string, attrs *RouteAttributes) IRoute {
	_init_.Initialize()

	var returns IRoute

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Route",
		"fromRouteAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Route_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Route",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Route_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Route",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (r *jsiiProxy_Route) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (r *jsiiProxy_Route) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (r *jsiiProxy_Route) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (r *jsiiProxy_Route) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (r *jsiiProxy_Route) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_Route) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (r *jsiiProxy_Route) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (r *jsiiProxy_Route) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_Route) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (r *jsiiProxy_Route) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (r *jsiiProxy_Route) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties ncecessary to import a reusable Route.
// Experimental.
type RouteAttributes struct {
	// The name of the Route.
	// Experimental.
	RouteName *string `json:"routeName"`
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter IVirtualRouter `json:"virtualRouter"`
}

// Base interface properties for all Routes.
//
// TODO: EXAMPLE
//
// Experimental.
type RouteBaseProps struct {
	// Protocol specific spec.
	// Experimental.
	RouteSpec RouteSpec `json:"routeSpec"`
	// The name of the route.
	// Experimental.
	RouteName *string `json:"routeName"`
}

// Properties to define new Routes.
// Experimental.
type RouteProps struct {
	// Protocol specific spec.
	// Experimental.
	RouteSpec RouteSpec `json:"routeSpec"`
	// The name of the route.
	// Experimental.
	RouteName *string `json:"routeName"`
	// The service mesh to define the route in.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter IVirtualRouter `json:"virtualRouter"`
}

// Used to generate specs with different protocols for a RouteSpec.
//
// TODO: EXAMPLE
//
// Experimental.
type RouteSpec interface {
	Bind(scope awscdk.Construct) *RouteSpecConfig
}

// The jsii proxy struct for RouteSpec
type jsiiProxy_RouteSpec struct {
	_ byte // padding
}

// Experimental.
func NewRouteSpec_Override(r RouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.RouteSpec",
		nil, // no parameters
		r,
	)
}

// Creates a GRPC Based RouteSpec.
// Experimental.
func RouteSpec_Grpc(options *GrpcRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based RouteSpec.
// Experimental.
func RouteSpec_Http(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based RouteSpec.
// Experimental.
func RouteSpec_Http2(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a TCP Based RouteSpec.
// Experimental.
func RouteSpec_Tcp(options *TcpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"tcp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Called when the RouteSpec type is initialized.
//
// Can be used to enforce
// mutual exclusivity with future properties
// Experimental.
func (r *jsiiProxy_RouteSpec) Bind(scope awscdk.Construct) *RouteSpecConfig {
	var returns *RouteSpecConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// All Properties for Route Specs.
// Experimental.
type RouteSpecConfig struct {
	// The spec for a grpc route.
	// Experimental.
	GrpcRouteSpec *CfnRoute_GrpcRouteProperty `json:"grpcRouteSpec"`
	// The spec for an http2 route.
	// Experimental.
	Http2RouteSpec *CfnRoute_HttpRouteProperty `json:"http2RouteSpec"`
	// The spec for an http route.
	// Experimental.
	HttpRouteSpec *CfnRoute_HttpRouteProperty `json:"httpRouteSpec"`
	// The priority for the route.
	//
	// Routes are matched based on the specified
	// value, where 0 is the highest priority.
	// Experimental.
	Priority *float64 `json:"priority"`
	// The spec for a tcp route.
	// Experimental.
	TcpRouteSpec *CfnRoute_TcpRouteProperty `json:"tcpRouteSpec"`
}

// Base options for all route specs.
// Experimental.
type RouteSpecOptionsBase struct {
	// The priority for the route.
	//
	// Routes are matched based on the specified
	// value, where 0 is the highest priority.
	// Experimental.
	Priority *float64 `json:"priority"`
}

// Provides the Service Discovery method a VirtualNode uses.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceDiscovery interface {
	Bind(scope awscdk.Construct) *ServiceDiscoveryConfig
}

// The jsii proxy struct for ServiceDiscovery
type jsiiProxy_ServiceDiscovery struct {
	_ byte // padding
}

// Experimental.
func NewServiceDiscovery_Override(s ServiceDiscovery) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.ServiceDiscovery",
		nil, // no parameters
		s,
	)
}

// Returns Cloud Map based service discovery.
// Experimental.
func ServiceDiscovery_CloudMap(service awsservicediscovery.IService, instanceAttributes *map[string]*string) ServiceDiscovery {
	_init_.Initialize()

	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.ServiceDiscovery",
		"cloudMap",
		[]interface{}{service, instanceAttributes},
		&returns,
	)

	return returns
}

// Returns DNS based service discovery.
// Experimental.
func ServiceDiscovery_Dns(hostname *string, responseType DnsResponseType) ServiceDiscovery {
	_init_.Initialize()

	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.ServiceDiscovery",
		"dns",
		[]interface{}{hostname, responseType},
		&returns,
	)

	return returns
}

// Binds the current object when adding Service Discovery to a VirtualNode.
// Experimental.
func (s *jsiiProxy_ServiceDiscovery) Bind(scope awscdk.Construct) *ServiceDiscoveryConfig {
	var returns *ServiceDiscoveryConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for VirtualNode Service Discovery.
// Experimental.
type ServiceDiscoveryConfig struct {
	// Cloud Map based Service Discovery.
	// Experimental.
	Cloudmap *CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty `json:"cloudmap"`
	// DNS based Service Discovery.
	// Experimental.
	Dns *CfnVirtualNode_DnsServiceDiscoveryProperty `json:"dns"`
}

// Used to generate Subject Alternative Names Matchers.
//
// TODO: EXAMPLE
//
// Experimental.
type SubjectAlternativeNames interface {
	Bind(scope awscdk.Construct) *SubjectAlternativeNamesMatcherConfig
}

// The jsii proxy struct for SubjectAlternativeNames
type jsiiProxy_SubjectAlternativeNames struct {
	_ byte // padding
}

// Experimental.
func NewSubjectAlternativeNames_Override(s SubjectAlternativeNames) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.SubjectAlternativeNames",
		nil, // no parameters
		s,
	)
}

// The values of the SAN must match the specified values exactly.
// Experimental.
func SubjectAlternativeNames_MatchingExactly(names ...*string) SubjectAlternativeNames {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range names {
		args = append(args, a)
	}

	var returns SubjectAlternativeNames

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.SubjectAlternativeNames",
		"matchingExactly",
		args,
		&returns,
	)

	return returns
}

// Returns Subject Alternative Names Matcher based on method type.
// Experimental.
func (s *jsiiProxy_SubjectAlternativeNames) Bind(scope awscdk.Construct) *SubjectAlternativeNamesMatcherConfig {
	var returns *SubjectAlternativeNamesMatcherConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// All Properties for Subject Alternative Names Matcher for both Client Policy and Listener.
// Experimental.
type SubjectAlternativeNamesMatcherConfig struct {
	// VirtualNode CFN configuration for subject alternative names secured by the certificate.
	// Experimental.
	SubjectAlternativeNamesMatch *CfnVirtualNode_SubjectAlternativeNameMatchersProperty `json:"subjectAlternativeNamesMatch"`
}

// Connection pool properties for TCP listeners.
// Experimental.
type TcpConnectionPool struct {
	// The maximum connections in the pool.
	// Experimental.
	MaxConnections *float64 `json:"maxConnections"`
}

// Properties used to define TCP Based healthchecks.
// Experimental.
type TcpHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// Experimental.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// The time period between each health check execution.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

// TCP events on which you may retry.
// Experimental.
type TcpRetryEvent string

const (
	TcpRetryEvent_CONNECTION_ERROR TcpRetryEvent = "CONNECTION_ERROR"
)

// Properties specific for a TCP Based Routes.
// Experimental.
type TcpRouteSpecOptions struct {
	// The priority for the route.
	//
	// Routes are matched based on the specified
	// value, where 0 is the highest priority.
	// Experimental.
	Priority *float64 `json:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	// Experimental.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets"`
	// An object that represents a tcp timeout.
	// Experimental.
	Timeout *TcpTimeout `json:"timeout"`
}

// Represents timeouts for TCP protocols.
// Experimental.
type TcpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Experimental.
	Idle awscdk.Duration `json:"idle"`
}

// Represent the TCP Node Listener prorperty.
// Experimental.
type TcpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *TcpConnectionPool `json:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `json:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `json:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `json:"port"`
	// Timeout for TCP protocol.
	// Experimental.
	Timeout *TcpTimeout `json:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `json:"tls"`
}

// Represents a TLS certificate.
//
// TODO: EXAMPLE
//
// Experimental.
type TlsCertificate interface {
	Bind(_scope awscdk.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for TlsCertificate
type jsiiProxy_TlsCertificate struct {
	_ byte // padding
}

// Experimental.
func NewTlsCertificate_Override(t TlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.TlsCertificate",
		nil, // no parameters
		t,
	)
}

// Returns an ACM TLS Certificate.
// Experimental.
func TlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
// Experimental.
func TlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
// Experimental.
func TlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns TLS certificate based provider.
// Experimental.
func (t *jsiiProxy_TlsCertificate) Bind(_scope awscdk.Construct) *TlsCertificateConfig {
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// A wrapper for the tls config returned by {@link TlsCertificate.bind}.
// Experimental.
type TlsCertificateConfig struct {
	// The CFN shape for a TLS certificate.
	// Experimental.
	TlsCertificate *CfnVirtualNode_ListenerTlsCertificateProperty `json:"tlsCertificate"`
}

// Represents the properties needed to define client policy.
//
// TODO: EXAMPLE
//
// Experimental.
type TlsClientPolicy struct {
	// Represents the object for TLS validation context.
	// Experimental.
	Validation *TlsValidation `json:"validation"`
	// Whether the policy is enforced.
	// Experimental.
	Enforce *bool `json:"enforce"`
	// Represents a client TLS certificate.
	//
	// The certificate will be sent only if the server requests it, enabling mutual TLS.
	// Experimental.
	MutualTlsCertificate MutualTlsCertificate `json:"mutualTlsCertificate"`
	// TLS is enforced on the ports specified here.
	//
	// If no ports are specified, TLS will be enforced on all the ports.
	// Experimental.
	Ports *[]*float64 `json:"ports"`
}

// Enum of supported TLS modes.
//
// TODO: EXAMPLE
//
// Experimental.
type TlsMode string

const (
	TlsMode_STRICT TlsMode = "STRICT"
	TlsMode_PERMISSIVE TlsMode = "PERMISSIVE"
	TlsMode_DISABLED TlsMode = "DISABLED"
)

// Represents the properties needed to define TLS Validation context.
//
// TODO: EXAMPLE
//
// Experimental.
type TlsValidation struct {
	// Reference to where to retrieve the trust chain.
	// Experimental.
	Trust TlsValidationTrust `json:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	// Experimental.
	SubjectAlternativeNames SubjectAlternativeNames `json:"subjectAlternativeNames"`
}

// Defines the TLS Validation Context Trust.
//
// TODO: EXAMPLE
//
// Experimental.
type TlsValidationTrust interface {
	Bind(scope awscdk.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for TlsValidationTrust
type jsiiProxy_TlsValidationTrust struct {
	_ byte // padding
}

// Experimental.
func NewTlsValidationTrust_Override(t TlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.TlsValidationTrust",
		nil, // no parameters
		t,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
// Experimental.
func TlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
// Experimental.
func TlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
// Experimental.
func TlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns Trust context based on trust type.
// Experimental.
func (t *jsiiProxy_TlsValidationTrust) Bind(scope awscdk.Construct) *TlsValidationTrustConfig {
	var returns *TlsValidationTrustConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// All Properties for TLS Validation Trusts for both Client Policy and Listener.
// Experimental.
type TlsValidationTrustConfig struct {
	// VirtualNode CFN configuration for client policy's TLS Validation Trust.
	// Experimental.
	TlsValidationTrust *CfnVirtualNode_TlsValidationContextTrustProperty `json:"tlsValidationTrust"`
}

// VirtualGateway represents a newly defined App Mesh Virtual Gateway.
//
// A virtual gateway allows resources that are outside of your mesh to communicate to resources that
// are inside of your mesh.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html
//
// Experimental.
type VirtualGateway interface {
	awscdk.Resource
	IVirtualGateway
	Env() *awscdk.ResourceEnvironment
	Listeners() *[]*VirtualGatewayListenerConfig
	Mesh() IMesh
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualGatewayArn() *string
	VirtualGatewayName() *string
	AddGatewayRoute(id *string, props *GatewayRouteBaseProps) GatewayRoute
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for VirtualGateway
type jsiiProxy_VirtualGateway struct {
	internal.Type__awscdkResource
	jsiiProxy_IVirtualGateway
}

func (j *jsiiProxy_VirtualGateway) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) Listeners() *[]*VirtualGatewayListenerConfig {
	var returns *[]*VirtualGatewayListenerConfig
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) VirtualGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualGateway) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVirtualGateway(scope constructs.Construct, id *string, props *VirtualGatewayProps) VirtualGateway {
	_init_.Initialize()

	j := jsiiProxy_VirtualGateway{}

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVirtualGateway_Override(v VirtualGateway, scope constructs.Construct, id *string, props *VirtualGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualGateway",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualGateway given an ARN.
// Experimental.
func VirtualGateway_FromVirtualGatewayArn(scope constructs.Construct, id *string, virtualGatewayArn *string) IVirtualGateway {
	_init_.Initialize()

	var returns IVirtualGateway

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGateway",
		"fromVirtualGatewayArn",
		[]interface{}{scope, id, virtualGatewayArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualGateway given its attributes.
// Experimental.
func VirtualGateway_FromVirtualGatewayAttributes(scope constructs.Construct, id *string, attrs *VirtualGatewayAttributes) IVirtualGateway {
	_init_.Initialize()

	var returns IVirtualGateway

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGateway",
		"fromVirtualGatewayAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func VirtualGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VirtualGateway_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGateway",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Utility method to add a new GatewayRoute to the VirtualGateway.
// Experimental.
func (v *jsiiProxy_VirtualGateway) AddGatewayRoute(id *string, props *GatewayRouteBaseProps) GatewayRoute {
	var returns GatewayRoute

	_jsii_.Invoke(
		v,
		"addGatewayRoute",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

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
func (v *jsiiProxy_VirtualGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (v *jsiiProxy_VirtualGateway) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (v *jsiiProxy_VirtualGateway) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (v *jsiiProxy_VirtualGateway) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants the given entity `appmesh:StreamAggregatedResources`.
// Experimental.
func (v *jsiiProxy_VirtualGateway) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualGateway) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualGateway) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualGateway) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualGateway) Prepare() {
	_jsii_.InvokeVoid(
		v,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualGateway) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (v *jsiiProxy_VirtualGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualGateway) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Unterface with properties necessary to import a reusable VirtualGateway.
// Experimental.
type VirtualGatewayAttributes struct {
	// The Mesh that the VirtualGateway belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// The name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName *string `json:"virtualGatewayName"`
}

// Basic configuration properties for a VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualGatewayBaseProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Experimental.
	AccessLog AccessLog `json:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `json:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Experimental.
	Listeners *[]VirtualGatewayListener `json:"listeners"`
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName *string `json:"virtualGatewayName"`
}

// Represents the properties needed to define listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualGatewayListener interface {
	Bind(scope awscdk.Construct) *VirtualGatewayListenerConfig
}

// The jsii proxy struct for VirtualGatewayListener
type jsiiProxy_VirtualGatewayListener struct {
	_ byte // padding
}

// Experimental.
func NewVirtualGatewayListener_Override(v VirtualGatewayListener) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualGatewayListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualGateway.
// Experimental.
func VirtualGatewayListener_Grpc(options *GrpcGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGatewayListener",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualGateway.
// Experimental.
func VirtualGatewayListener_Http(options *HttpGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGatewayListener",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualGateway.
// Experimental.
func VirtualGatewayListener_Http2(options *Http2GatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualGatewayListener",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Called when the GatewayListener type is initialized.
//
// Can be used to enforce
// mutual exclusivity
// Experimental.
func (v *jsiiProxy_VirtualGatewayListener) Bind(scope awscdk.Construct) *VirtualGatewayListenerConfig {
	var returns *VirtualGatewayListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for a VirtualGateway listener.
// Experimental.
type VirtualGatewayListenerConfig struct {
	// Single listener config for a VirtualGateway.
	// Experimental.
	Listener *CfnVirtualGateway_VirtualGatewayListenerProperty `json:"listener"`
}

// Properties used when creating a new VirtualGateway.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualGatewayProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Experimental.
	AccessLog AccessLog `json:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `json:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Experimental.
	Listeners *[]VirtualGatewayListener `json:"listeners"`
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName *string `json:"virtualGatewayName"`
	// The Mesh which the VirtualGateway belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
}

// VirtualNode represents a newly defined AppMesh VirtualNode.
//
// Any inbound traffic that your virtual node expects should be specified as a
// listener. Any outbound traffic that your virtual node expects to reach
// should be specified as a backend.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html
//
// Experimental.
type VirtualNode interface {
	awscdk.Resource
	IVirtualNode
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualNodeArn() *string
	VirtualNodeName() *string
	AddBackend(backend Backend)
	AddListener(listener VirtualNodeListener)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for VirtualNode
type jsiiProxy_VirtualNode struct {
	internal.Type__awscdkResource
	jsiiProxy_IVirtualNode
}

func (j *jsiiProxy_VirtualNode) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) VirtualNodeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNode) VirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVirtualNode(scope constructs.Construct, id *string, props *VirtualNodeProps) VirtualNode {
	_init_.Initialize()

	j := jsiiProxy_VirtualNode{}

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVirtualNode_Override(v VirtualNode, scope constructs.Construct, id *string, props *VirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualNode",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualNode given an ARN.
// Experimental.
func VirtualNode_FromVirtualNodeArn(scope constructs.Construct, id *string, virtualNodeArn *string) IVirtualNode {
	_init_.Initialize()

	var returns IVirtualNode

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNode",
		"fromVirtualNodeArn",
		[]interface{}{scope, id, virtualNodeArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualNode given its name.
// Experimental.
func VirtualNode_FromVirtualNodeAttributes(scope constructs.Construct, id *string, attrs *VirtualNodeAttributes) IVirtualNode {
	_init_.Initialize()

	var returns IVirtualNode

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNode",
		"fromVirtualNodeAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func VirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VirtualNode_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNode",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a Virtual Services that this node is expected to send outbound traffic to.
// Experimental.
func (v *jsiiProxy_VirtualNode) AddBackend(backend Backend) {
	_jsii_.InvokeVoid(
		v,
		"addBackend",
		[]interface{}{backend},
	)
}

// Utility method to add an inbound listener for this VirtualNode.
//
// Note: At this time, Virtual Nodes support at most one listener. Adding
// more than one will result in a failure to deploy the CloudFormation stack.
// However, the App Mesh team has plans to add support for multiple listeners
// on Virtual Nodes and Virtual Routers.
// See: https://github.com/aws/aws-app-mesh-roadmap/issues/120
//
// Experimental.
func (v *jsiiProxy_VirtualNode) AddListener(listener VirtualNodeListener) {
	_jsii_.InvokeVoid(
		v,
		"addListener",
		[]interface{}{listener},
	)
}

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
func (v *jsiiProxy_VirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (v *jsiiProxy_VirtualNode) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (v *jsiiProxy_VirtualNode) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (v *jsiiProxy_VirtualNode) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants the given entity `appmesh:StreamAggregatedResources`.
// Experimental.
func (v *jsiiProxy_VirtualNode) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualNode) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualNode) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualNode) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualNode) Prepare() {
	_jsii_.InvokeVoid(
		v,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualNode) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (v *jsiiProxy_VirtualNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualNode) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties necessary to import a reusable VirtualNode.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualNodeAttributes struct {
	// The Mesh that the VirtualNode belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName *string `json:"virtualNodeName"`
}

// Basic configuration properties for a VirtualNode.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualNodeBaseProps struct {
	// Access Logging Configuration for the virtual node.
	// Experimental.
	AccessLog AccessLog `json:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `json:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	// Experimental.
	Backends *[]Backend `json:"backends"`
	// Initial listener for the virtual node.
	// Experimental.
	Listeners *[]VirtualNodeListener `json:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	// Experimental.
	ServiceDiscovery ServiceDiscovery `json:"serviceDiscovery"`
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName *string `json:"virtualNodeName"`
}

// Defines listener for a VirtualNode.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualNodeListener interface {
	Bind(scope awscdk.Construct) *VirtualNodeListenerConfig
}

// The jsii proxy struct for VirtualNodeListener
type jsiiProxy_VirtualNodeListener struct {
	_ byte // padding
}

// Experimental.
func NewVirtualNodeListener_Override(v VirtualNodeListener) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualNodeListener",
		nil, // no parameters
		v,
	)
}

// Returns an GRPC Listener for a VirtualNode.
// Experimental.
func VirtualNodeListener_Grpc(props *GrpcVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNodeListener",
		"grpc",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualNode.
// Experimental.
func VirtualNodeListener_Http(props *HttpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNodeListener",
		"http",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualNode.
// Experimental.
func VirtualNodeListener_Http2(props *Http2VirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNodeListener",
		"http2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an TCP Listener for a VirtualNode.
// Experimental.
func VirtualNodeListener_Tcp(props *TcpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualNodeListener",
		"tcp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Binds the current object when adding Listener to a VirtualNode.
// Experimental.
func (v *jsiiProxy_VirtualNodeListener) Bind(scope awscdk.Construct) *VirtualNodeListenerConfig {
	var returns *VirtualNodeListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for a VirtualNode listener.
// Experimental.
type VirtualNodeListenerConfig struct {
	// Single listener config for a VirtualNode.
	// Experimental.
	Listener *CfnVirtualNode_ListenerProperty `json:"listener"`
}

// The properties used when creating a new VirtualNode.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualNodeProps struct {
	// Access Logging Configuration for the virtual node.
	// Experimental.
	AccessLog AccessLog `json:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `json:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	// Experimental.
	Backends *[]Backend `json:"backends"`
	// Initial listener for the virtual node.
	// Experimental.
	Listeners *[]VirtualNodeListener `json:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	// Experimental.
	ServiceDiscovery ServiceDiscovery `json:"serviceDiscovery"`
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName *string `json:"virtualNodeName"`
	// The Mesh which the VirtualNode belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
}

// TODO: EXAMPLE
//
// Experimental.
type VirtualRouter interface {
	awscdk.Resource
	IVirtualRouter
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualRouterArn() *string
	VirtualRouterName() *string
	AddRoute(id *string, props *RouteBaseProps) Route
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for VirtualRouter
type jsiiProxy_VirtualRouter struct {
	internal.Type__awscdkResource
	jsiiProxy_IVirtualRouter
}

func (j *jsiiProxy_VirtualRouter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) VirtualRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVirtualRouter(scope constructs.Construct, id *string, props *VirtualRouterProps) VirtualRouter {
	_init_.Initialize()

	j := jsiiProxy_VirtualRouter{}

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVirtualRouter_Override(v VirtualRouter, scope constructs.Construct, id *string, props *VirtualRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualRouter given an ARN.
// Experimental.
func VirtualRouter_FromVirtualRouterArn(scope constructs.Construct, id *string, virtualRouterArn *string) IVirtualRouter {
	_init_.Initialize()

	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouter",
		"fromVirtualRouterArn",
		[]interface{}{scope, id, virtualRouterArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualRouter given attributes.
// Experimental.
func VirtualRouter_FromVirtualRouterAttributes(scope constructs.Construct, id *string, attrs *VirtualRouterAttributes) IVirtualRouter {
	_init_.Initialize()

	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouter",
		"fromVirtualRouterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func VirtualRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VirtualRouter_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a single route to the router.
// Experimental.
func (v *jsiiProxy_VirtualRouter) AddRoute(id *string, props *RouteBaseProps) Route {
	var returns Route

	_jsii_.Invoke(
		v,
		"addRoute",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

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
func (v *jsiiProxy_VirtualRouter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (v *jsiiProxy_VirtualRouter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (v *jsiiProxy_VirtualRouter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (v *jsiiProxy_VirtualRouter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualRouter) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualRouter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualRouter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualRouter) Prepare() {
	_jsii_.InvokeVoid(
		v,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualRouter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (v *jsiiProxy_VirtualRouter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualRouter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties ncecessary to import a reusable VirtualRouter.
// Experimental.
type VirtualRouterAttributes struct {
	// The Mesh which the VirtualRouter belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName *string `json:"virtualRouterName"`
}

// Interface with base properties all routers willl inherit.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	// Experimental.
	Listeners *[]VirtualRouterListener `json:"listeners"`
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName *string `json:"virtualRouterName"`
}

// Represents the properties needed to define listeners for a VirtualRouter.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualRouterListener interface {
	Bind(scope awscdk.Construct) *VirtualRouterListenerConfig
}

// The jsii proxy struct for VirtualRouterListener
type jsiiProxy_VirtualRouterListener struct {
	_ byte // padding
}

// Experimental.
func NewVirtualRouterListener_Override(v VirtualRouterListener) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualRouterListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualRouter.
// Experimental.
func VirtualRouterListener_Grpc(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouterListener",
		"grpc",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualRouter.
// Experimental.
func VirtualRouterListener_Http(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouterListener",
		"http",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualRouter.
// Experimental.
func VirtualRouterListener_Http2(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouterListener",
		"http2",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns a TCP Listener for a VirtualRouter.
// Experimental.
func VirtualRouterListener_Tcp(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualRouterListener",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Called when the VirtualRouterListener type is initialized.
//
// Can be used to enforce
// mutual exclusivity
// Experimental.
func (v *jsiiProxy_VirtualRouterListener) Bind(scope awscdk.Construct) *VirtualRouterListenerConfig {
	var returns *VirtualRouterListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for a VirtualRouter listener.
// Experimental.
type VirtualRouterListenerConfig struct {
	// Single listener config for a VirtualRouter.
	// Experimental.
	Listener *CfnVirtualRouter_VirtualRouterListenerProperty `json:"listener"`
}

// The properties used when creating a new VirtualRouter.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualRouterProps struct {
	// Listener specification for the VirtualRouter.
	// Experimental.
	Listeners *[]VirtualRouterListener `json:"listeners"`
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName *string `json:"virtualRouterName"`
	// The Mesh which the VirtualRouter belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
}

// VirtualService represents a service inside an AppMesh.
//
// It routes traffic either to a Virtual Node or to a Virtual Router.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html
//
// Experimental.
type VirtualService interface {
	awscdk.Resource
	IVirtualService
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualServiceArn() *string
	VirtualServiceName() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for VirtualService
type jsiiProxy_VirtualService struct {
	internal.Type__awscdkResource
	jsiiProxy_IVirtualService
}

func (j *jsiiProxy_VirtualService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) VirtualServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) VirtualServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVirtualService(scope constructs.Construct, id *string, props *VirtualServiceProps) VirtualService {
	_init_.Initialize()

	j := jsiiProxy_VirtualService{}

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVirtualService_Override(v VirtualService, scope constructs.Construct, id *string, props *VirtualServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualService",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualService given an ARN.
// Experimental.
func VirtualService_FromVirtualServiceArn(scope constructs.Construct, id *string, virtualServiceArn *string) IVirtualService {
	_init_.Initialize()

	var returns IVirtualService

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualService",
		"fromVirtualServiceArn",
		[]interface{}{scope, id, virtualServiceArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualService given its attributes.
// Experimental.
func VirtualService_FromVirtualServiceAttributes(scope constructs.Construct, id *string, attrs *VirtualServiceAttributes) IVirtualService {
	_init_.Initialize()

	var returns IVirtualService

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualService",
		"fromVirtualServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func VirtualService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VirtualService_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualService",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (v *jsiiProxy_VirtualService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (v *jsiiProxy_VirtualService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (v *jsiiProxy_VirtualService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (v *jsiiProxy_VirtualService) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualService) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (v *jsiiProxy_VirtualService) Prepare() {
	_jsii_.InvokeVoid(
		v,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (v *jsiiProxy_VirtualService) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (v *jsiiProxy_VirtualService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (v *jsiiProxy_VirtualService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties ncecessary to import a reusable VirtualService.
// Experimental.
type VirtualServiceAttributes struct {
	// The Mesh which the VirtualService belongs to.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// The name of the VirtualService, it is recommended this follows the fully-qualified domain name format.
	// Experimental.
	VirtualServiceName *string `json:"virtualServiceName"`
}

// Represents the properties needed to define a Virtual Service backend.
// Experimental.
type VirtualServiceBackendOptions struct {
	// TLS properties for  Client policy for the backend.
	// Experimental.
	TlsClientPolicy *TlsClientPolicy `json:"tlsClientPolicy"`
}

// The properties applied to the VirtualService being defined.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualServiceProps struct {
	// The VirtualNode or VirtualRouter which the VirtualService uses as its provider.
	// Experimental.
	VirtualServiceProvider VirtualServiceProvider `json:"virtualServiceProvider"`
	// The name of the VirtualService.
	//
	// It is recommended this follows the fully-qualified domain name format,
	// such as "my-service.default.svc.cluster.local".
	//
	// Example value: `service.domain.local`
	// Experimental.
	VirtualServiceName *string `json:"virtualServiceName"`
}

// Represents the properties needed to define the provider for a VirtualService.
//
// TODO: EXAMPLE
//
// Experimental.
type VirtualServiceProvider interface {
	Bind(_construct constructs.Construct) *VirtualServiceProviderConfig
}

// The jsii proxy struct for VirtualServiceProvider
type jsiiProxy_VirtualServiceProvider struct {
	_ byte // padding
}

// Experimental.
func NewVirtualServiceProvider_Override(v VirtualServiceProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		nil, // no parameters
		v,
	)
}

// Returns an Empty Provider for a VirtualService.
//
// This provides no routing capabilities
// and should only be used as a placeholder
// Experimental.
func VirtualServiceProvider_None(mesh IMesh) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"none",
		[]interface{}{mesh},
		&returns,
	)

	return returns
}

// Returns a VirtualNode based Provider for a VirtualService.
// Experimental.
func VirtualServiceProvider_VirtualNode(virtualNode IVirtualNode) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"virtualNode",
		[]interface{}{virtualNode},
		&returns,
	)

	return returns
}

// Returns a VirtualRouter based Provider for a VirtualService.
// Experimental.
func VirtualServiceProvider_VirtualRouter(virtualRouter IVirtualRouter) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"virtualRouter",
		[]interface{}{virtualRouter},
		&returns,
	)

	return returns
}

// Enforces mutual exclusivity for VirtualService provider types.
// Experimental.
func (v *jsiiProxy_VirtualServiceProvider) Bind(_construct constructs.Construct) *VirtualServiceProviderConfig {
	var returns *VirtualServiceProviderConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{_construct},
		&returns,
	)

	return returns
}

// Properties for a VirtualService provider.
// Experimental.
type VirtualServiceProviderConfig struct {
	// Mesh the Provider is using.
	// Experimental.
	Mesh IMesh `json:"mesh"`
	// Virtual Node based provider.
	// Experimental.
	VirtualNodeProvider *CfnVirtualService_VirtualNodeServiceProviderProperty `json:"virtualNodeProvider"`
	// Virtual Router based provider.
	// Experimental.
	VirtualRouterProvider *CfnVirtualService_VirtualRouterServiceProviderProperty `json:"virtualRouterProvider"`
}

// Properties for the Weighted Targets in the route.
// Experimental.
type WeightedTarget struct {
	// The VirtualNode the route points to.
	// Experimental.
	VirtualNode IVirtualNode `json:"virtualNode"`
	// The weight for the target.
	// Experimental.
	Weight *float64 `json:"weight"`
}

