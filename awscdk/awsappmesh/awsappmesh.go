package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for Envoy Access logs for mesh endpoints.
//
// TODO: EXAMPLE
//
type AccessLog interface {
	Bind(scope constructs.Construct) *AccessLogConfig
}

// The jsii proxy struct for AccessLog
type jsiiProxy_AccessLog struct {
	_ byte // padding
}

func NewAccessLog_Override(a AccessLog) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		nil, // no parameters
		a,
	)
}

// Path to a file to write access logs to.
func AccessLog_FromFilePath(filePath *string) AccessLog {
	_init_.Initialize()

	var returns AccessLog

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.AccessLog",
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
func (a *jsiiProxy_AccessLog) Bind(scope constructs.Construct) *AccessLogConfig {
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
//
// TODO: EXAMPLE
//
type AccessLogConfig struct {
	// VirtualGateway CFN configuration for Access Logging.
	VirtualGatewayAccessLog *CfnVirtualGateway_VirtualGatewayAccessLogProperty `json:"virtualGatewayAccessLog" yaml:"virtualGatewayAccessLog"`
	// VirtualNode CFN configuration for Access Logging.
	VirtualNodeAccessLog *CfnVirtualNode_AccessLogProperty `json:"virtualNodeAccessLog" yaml:"virtualNodeAccessLog"`
}

// Contains static factory methods to create backends.
//
// TODO: EXAMPLE
//
type Backend interface {
	Bind(_scope constructs.Construct) *BackendConfig
}

// The jsii proxy struct for Backend
type jsiiProxy_Backend struct {
	_ byte // padding
}

func NewBackend_Override(b Backend) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Backend",
		nil, // no parameters
		b,
	)
}

// Construct a Virtual Service backend.
func Backend_VirtualService(virtualService IVirtualService, props *VirtualServiceBackendOptions) Backend {
	_init_.Initialize()

	var returns Backend

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Backend",
		"virtualService",
		[]interface{}{virtualService, props},
		&returns,
	)

	return returns
}

// Return backend config.
func (b *jsiiProxy_Backend) Bind(_scope constructs.Construct) *BackendConfig {
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
//
// TODO: EXAMPLE
//
type BackendConfig struct {
	// Config for a Virtual Service backend.
	VirtualServiceBackend *CfnVirtualNode_BackendProperty `json:"virtualServiceBackend" yaml:"virtualServiceBackend"`
}

// Represents the properties needed to define backend defaults.
//
// TODO: EXAMPLE
//
type BackendDefaults struct {
	// TLS properties for Client policy for backend defaults.
	TlsClientPolicy *TlsClientPolicy `json:"tlsClientPolicy" yaml:"tlsClientPolicy"`
}

// A CloudFormation `AWS::AppMesh::GatewayRoute`.
//
// Creates a gateway route.
//
// A gateway route is attached to a virtual gateway and routes traffic to an existing virtual service. If a route matches a request, it can distribute traffic to a target virtual service.
//
// For more information about gateway routes, see [Gateway routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnGatewayRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnGatewayRoute) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object representing the gateway route host name to match.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteHostnameMatchProperty struct {
	// The exact host name to match on.
	Exact *string `json:"exact" yaml:"exact"`
	// The specified ending characters of the host name to match on.
	Suffix *string `json:"suffix" yaml:"suffix"`
}

// An object representing the gateway route host name to rewrite.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteHostnameRewriteProperty struct {
	// The default target host name to write to.
	DefaultTargetHostname *string `json:"defaultTargetHostname" yaml:"defaultTargetHostname"`
}

// An object representing the method header to be matched.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteMetadataMatchProperty struct {
	// The exact method header to be matched on.
	Exact *string `json:"exact" yaml:"exact"`
	// The specified beginning characters of the method header to be matched on.
	Prefix *string `json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `json:"range" yaml:"range"`
	// The regex used to match the method header.
	Regex *string `json:"regex" yaml:"regex"`
	// The specified ending characters of the method header to match on.
	Suffix *string `json:"suffix" yaml:"suffix"`
}

// An object that represents the range of values to match on.
//
// The first character of the range is included in the range, though the last character is not. For example, if the range specified were 1-100, only values 1-99 would be matched.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteRangeMatchProperty struct {
	// The end of the range.
	End *float64 `json:"end" yaml:"end"`
	// The start of the range.
	Start *float64 `json:"start" yaml:"start"`
}

// An object that represents a gateway route specification.
//
// Specify one gateway route type.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteSpecProperty struct {
	// An object that represents the specification of a gRPC gateway route.
	GrpcRoute interface{} `json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 gateway route.
	Http2Route interface{} `json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP gateway route.
	HttpRoute interface{} `json:"httpRoute" yaml:"httpRoute"`
	// The ordering of the gateway routes spec.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// An object that represents a gateway route target.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteTargetProperty struct {
	// An object that represents a virtual service gateway route target.
	VirtualService interface{} `json:"virtualService" yaml:"virtualService"`
}

// An object that represents the virtual service that traffic is routed to.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GatewayRouteVirtualServiceProperty struct {
	// The name of the virtual service that traffic is routed to.
	VirtualServiceName *string `json:"virtualServiceName" yaml:"virtualServiceName"`
}

// An object that represents the action to take if a match is determined.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GrpcGatewayRouteActionProperty struct {
	// An object that represents the target that traffic is routed to when a request matches the gateway route.
	Target interface{} `json:"target" yaml:"target"`
	// The gateway route action to rewrite.
	Rewrite interface{} `json:"rewrite" yaml:"rewrite"`
}

// An object that represents the criteria for determining a request match.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GrpcGatewayRouteMatchProperty struct {
	// The gateway route host name to be matched on.
	Hostname interface{} `json:"hostname" yaml:"hostname"`
	// The gateway route metadata to be matched on.
	Metadata interface{} `json:"metadata" yaml:"metadata"`
	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// An object representing the metadata of the gateway route.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GrpcGatewayRouteMetadataProperty struct {
	// A name for the gateway route metadata.
	Name *string `json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `json:"invert" yaml:"invert"`
	// The criteria for determining a metadata match.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents a gRPC gateway route.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GrpcGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents the gateway route to rewrite.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_GrpcGatewayRouteRewriteProperty struct {
	// The host name of the gateway route to rewrite.
	Hostname interface{} `json:"hostname" yaml:"hostname"`
}

// An object that represents the action to take if a match is determined.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteActionProperty struct {
	// An object that represents the target that traffic is routed to when a request matches the gateway route.
	Target interface{} `json:"target" yaml:"target"`
	// The gateway route action to rewrite.
	Rewrite interface{} `json:"rewrite" yaml:"rewrite"`
}

// An object that represents the method and value to match with the header value sent in a request.
//
// Specify one match method.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteHeaderMatchProperty struct {
	// The value sent by the client must match the specified value exactly.
	Exact *string `json:"exact" yaml:"exact"`
	// The value sent by the client must begin with the specified characters.
	Prefix *string `json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `json:"range" yaml:"range"`
	// The value sent by the client must include the specified characters.
	Regex *string `json:"regex" yaml:"regex"`
	// The value sent by the client must end with the specified characters.
	Suffix *string `json:"suffix" yaml:"suffix"`
}

// An object that represents the HTTP header in the gateway route.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteHeaderProperty struct {
	// A name for the HTTP header in the gateway route that will be matched on.
	Name *string `json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `json:"invert" yaml:"invert"`
	// An object that represents the method and value to match with the header value sent in a request.
	//
	// Specify one match method.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents the criteria for determining a request match.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteMatchProperty struct {
	// The client request headers to match on.
	Headers interface{} `json:"headers" yaml:"headers"`
	// The host name to match on.
	Hostname interface{} `json:"hostname" yaml:"hostname"`
	// The method to match on.
	Method *string `json:"method" yaml:"method"`
	// The path to match on.
	Path interface{} `json:"path" yaml:"path"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	Prefix *string `json:"prefix" yaml:"prefix"`
	// The query parameter to match on.
	QueryParameters interface{} `json:"queryParameters" yaml:"queryParameters"`
}

// An object that represents the path to rewrite.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty struct {
	// The exact path to rewrite.
	Exact *string `json:"exact" yaml:"exact"`
}

// An object representing the beginning characters of the route to rewrite.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty struct {
	// The default prefix used to replace the incoming route prefix when rewritten.
	DefaultPrefix *string `json:"defaultPrefix" yaml:"defaultPrefix"`
	// The value used to replace the incoming route prefix when rewritten.
	Value *string `json:"value" yaml:"value"`
}

// An object that represents an HTTP gateway route.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `json:"match" yaml:"match"`
}

// An object representing the gateway route to rewrite.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpGatewayRouteRewriteProperty struct {
	// The host name to rewrite.
	Hostname interface{} `json:"hostname" yaml:"hostname"`
	// The path to rewrite.
	Path interface{} `json:"path" yaml:"path"`
	// The specified beginning characters to rewrite.
	Prefix interface{} `json:"prefix" yaml:"prefix"`
}

// An object representing the path to match in the request.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpPathMatchProperty struct {
	// The exact path to match on.
	Exact *string `json:"exact" yaml:"exact"`
	// The regex used to match the path.
	Regex *string `json:"regex" yaml:"regex"`
}

// An object representing the query parameter to match.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_HttpQueryParameterMatchProperty struct {
	// The exact query parameter to match on.
	Exact *string `json:"exact" yaml:"exact"`
}

// An object that represents the query parameter in the request.
//
// TODO: EXAMPLE
//
type CfnGatewayRoute_QueryParameterProperty struct {
	// A name for the query parameter that will be matched on.
	Name *string `json:"name" yaml:"name"`
	// The query parameter to match on.
	Match interface{} `json:"match" yaml:"match"`
}

// Properties for defining a `CfnGatewayRoute`.
//
// TODO: EXAMPLE
//
type CfnGatewayRouteProps struct {
	// The name of the service mesh that the resource resides in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The specifications of the gateway route.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The virtual gateway that the gateway route is associated with.
	VirtualGatewayName *string `json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The name of the gateway route.
	GatewayRouteName *string `json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the gateway route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AppMesh::Mesh`.
//
// Creates a service mesh.
//
// A service mesh is a logical boundary for network traffic between services that are represented by resources within the mesh. After you create your service mesh, you can create virtual services, virtual nodes, virtual routers, and routes to distribute traffic between the applications in your mesh.
//
// For more information about service meshes, see [Service meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnMesh) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnMesh) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents the egress filter rules for a service mesh.
//
// TODO: EXAMPLE
//
type CfnMesh_EgressFilterProperty struct {
	// The egress filter type.
	//
	// By default, the type is `DROP_ALL` , which allows egress only from virtual nodes to other defined resources in the service mesh (and any traffic to `*.amazonaws.com` for AWS API calls). You can set the egress filter type to `ALLOW_ALL` to allow egress to any endpoint inside or outside of the service mesh.
	Type *string `json:"type" yaml:"type"`
}

// TODO: EXAMPLE
//
type CfnMesh_MeshServiceDiscoveryProperty struct {
}

// An object that represents the specification of a service mesh.
//
// TODO: EXAMPLE
//
type CfnMesh_MeshSpecProperty struct {
	// The egress filter rules for the service mesh.
	EgressFilter interface{} `json:"egressFilter" yaml:"egressFilter"`
}

// Properties for defining a `CfnMesh`.
//
// TODO: EXAMPLE
//
type CfnMeshProps struct {
	// The name to use for the service mesh.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The service mesh specification to apply.
	Spec interface{} `json:"spec" yaml:"spec"`
	// Optional metadata that you can apply to the service mesh to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AppMesh::Route`.
//
// Creates a route that is associated with a virtual router.
//
// You can route several different protocols and define a retry policy for a route. Traffic can be routed to one or more virtual nodes.
//
// For more information about routes, see [Routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnRoute) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents a duration of time.
//
// TODO: EXAMPLE
//
type CfnRoute_DurationProperty struct {
	// A unit of time.
	Unit *string `json:"unit" yaml:"unit"`
	// A number of time units.
	Value *float64 `json:"value" yaml:"value"`
}

// An object that represents a retry policy.
//
// Specify at least one value for at least one of the types of `RetryEvents` , a value for `maxRetries` , and a value for `perRetryTimeout` . Both `server-error` and `gateway-error` under `httpRetryEvents` include the Envoy `reset` policy. For more information on the `reset` policy, see the [Envoy documentation](https://docs.aws.amazon.com/https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on) .
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRetryPolicyProperty struct {
	// The maximum number of retry attempts.
	MaxRetries *float64 `json:"maxRetries" yaml:"maxRetries"`
	// The timeout for each retry attempt.
	PerRetryTimeout interface{} `json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Specify at least one of the valid values.
	GrpcRetryEvents *[]*string `json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
	// Specify at least one of the following values.
	//
	// - *server-error* – HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511
	// - *gateway-error* – HTTP status codes 502, 503, and 504
	// - *client-error* – HTTP status code 409
	// - *stream-error* – Retry on refused stream
	HttpRetryEvents *[]*string `json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Specify a valid value.
	//
	// The event occurs before any processing of a request has started and is encountered when the upstream is temporarily or permanently unavailable.
	TcpRetryEvents *[]*string `json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

// An object that represents the action to take if a match is determined.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	WeightedTargets interface{} `json:"weightedTargets" yaml:"weightedTargets"`
}

// An object that represents the criteria for determining a request match.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRouteMatchProperty struct {
	// An object that represents the data to match from the request.
	Metadata interface{} `json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If you specify a name, you must also specify a `serviceName` .
	MethodName *string `json:"methodName" yaml:"methodName"`
	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// An object that represents the match method.
//
// Specify one of the match values.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRouteMetadataMatchMethodProperty struct {
	// The value sent by the client must match the specified value exactly.
	Exact *string `json:"exact" yaml:"exact"`
	// The value sent by the client must begin with the specified characters.
	Prefix *string `json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `json:"range" yaml:"range"`
	// The value sent by the client must include the specified characters.
	Regex *string `json:"regex" yaml:"regex"`
	// The value sent by the client must end with the specified characters.
	Suffix *string `json:"suffix" yaml:"suffix"`
}

// An object that represents the match metadata for the route.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRouteMetadataProperty struct {
	// The name of the route.
	Name *string `json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `json:"invert" yaml:"invert"`
	// An object that represents the data to match from the request.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents a gRPC route type.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `json:"match" yaml:"match"`
	// An object that represents a retry policy.
	RetryPolicy interface{} `json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	Timeout interface{} `json:"timeout" yaml:"timeout"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnRoute_GrpcTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
	// An object that represents a per request timeout.
	//
	// The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15 seconds for the source and destination virtual node and the route.
	PerRequest interface{} `json:"perRequest" yaml:"perRequest"`
}

// An object that represents the method and value to match with the header value sent in a request.
//
// Specify one match method.
//
// TODO: EXAMPLE
//
type CfnRoute_HeaderMatchMethodProperty struct {
	// The value sent by the client must match the specified value exactly.
	Exact *string `json:"exact" yaml:"exact"`
	// The value sent by the client must begin with the specified characters.
	Prefix *string `json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `json:"range" yaml:"range"`
	// The value sent by the client must include the specified characters.
	Regex *string `json:"regex" yaml:"regex"`
	// The value sent by the client must end with the specified characters.
	Suffix *string `json:"suffix" yaml:"suffix"`
}

// An object representing the path to match in the request.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpPathMatchProperty struct {
	// The exact path to match on.
	Exact *string `json:"exact" yaml:"exact"`
	// The regex used to match the path.
	Regex *string `json:"regex" yaml:"regex"`
}

// An object representing the query parameter to match.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpQueryParameterMatchProperty struct {
	// The exact query parameter to match on.
	Exact *string `json:"exact" yaml:"exact"`
}

// An object that represents a retry policy.
//
// Specify at least one value for at least one of the types of `RetryEvents` , a value for `maxRetries` , and a value for `perRetryTimeout` . Both `server-error` and `gateway-error` under `httpRetryEvents` include the Envoy `reset` policy. For more information on the `reset` policy, see the [Envoy documentation](https://docs.aws.amazon.com/https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on) .
//
// TODO: EXAMPLE
//
type CfnRoute_HttpRetryPolicyProperty struct {
	// The maximum number of retry attempts.
	MaxRetries *float64 `json:"maxRetries" yaml:"maxRetries"`
	// The timeout for each retry attempt.
	PerRetryTimeout interface{} `json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Specify at least one of the following values.
	//
	// - *server-error* – HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511
	// - *gateway-error* – HTTP status codes 502, 503, and 504
	// - *client-error* – HTTP status code 409
	// - *stream-error* – Retry on refused stream
	HttpRetryEvents *[]*string `json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Specify a valid value.
	//
	// The event occurs before any processing of a request has started and is encountered when the upstream is temporarily or permanently unavailable.
	TcpRetryEvents *[]*string `json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

// An object that represents the action to take if a match is determined.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	WeightedTargets interface{} `json:"weightedTargets" yaml:"weightedTargets"`
}

// An object that represents the HTTP header in the request.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpRouteHeaderProperty struct {
	// A name for the HTTP header in the client request that will be matched on.
	Name *string `json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `json:"invert" yaml:"invert"`
	// The `HeaderMatchMethod` object.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents the requirements for a route to match HTTP requests for a virtual router.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpRouteMatchProperty struct {
	// The client request headers to match on.
	Headers interface{} `json:"headers" yaml:"headers"`
	// The client request method to match on.
	//
	// Specify only one.
	Method *string `json:"method" yaml:"method"`
	// The client request path to match on.
	Path interface{} `json:"path" yaml:"path"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	Prefix *string `json:"prefix" yaml:"prefix"`
	// The client request query parameters to match on.
	QueryParameters interface{} `json:"queryParameters" yaml:"queryParameters"`
	// The client request scheme to match on.
	//
	// Specify only one. Applicable only for HTTP2 routes.
	Scheme *string `json:"scheme" yaml:"scheme"`
}

// An object that represents an HTTP or HTTP/2 route type.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `json:"match" yaml:"match"`
	// An object that represents a retry policy.
	RetryPolicy interface{} `json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	Timeout interface{} `json:"timeout" yaml:"timeout"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnRoute_HttpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
	// An object that represents a per request timeout.
	//
	// The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15 seconds for the source and destination virtual node and the route.
	PerRequest interface{} `json:"perRequest" yaml:"perRequest"`
}

// An object that represents the range of values to match on.
//
// The first character of the range is included in the range, though the last character is not. For example, if the range specified were 1-100, only values 1-99 would be matched.
//
// TODO: EXAMPLE
//
type CfnRoute_MatchRangeProperty struct {
	// The end of the range.
	End *float64 `json:"end" yaml:"end"`
	// The start of the range.
	Start *float64 `json:"start" yaml:"start"`
}

// An object that represents the query parameter in the request.
//
// TODO: EXAMPLE
//
type CfnRoute_QueryParameterProperty struct {
	// A name for the query parameter that will be matched on.
	Name *string `json:"name" yaml:"name"`
	// The query parameter to match on.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents a route specification.
//
// Specify one route type.
//
// TODO: EXAMPLE
//
type CfnRoute_RouteSpecProperty struct {
	// An object that represents the specification of a gRPC route.
	GrpcRoute interface{} `json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 route.
	Http2Route interface{} `json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP route.
	HttpRoute interface{} `json:"httpRoute" yaml:"httpRoute"`
	// The priority for the route.
	//
	// Routes are matched based on the specified value, where 0 is the highest priority.
	Priority *float64 `json:"priority" yaml:"priority"`
	// An object that represents the specification of a TCP route.
	TcpRoute interface{} `json:"tcpRoute" yaml:"tcpRoute"`
}

// An object that represents the action to take if a match is determined.
//
// TODO: EXAMPLE
//
type CfnRoute_TcpRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	WeightedTargets interface{} `json:"weightedTargets" yaml:"weightedTargets"`
}

// An object that represents a TCP route type.
//
// TODO: EXAMPLE
//
type CfnRoute_TcpRouteProperty struct {
	// The action to take if a match is determined.
	Action interface{} `json:"action" yaml:"action"`
	// An object that represents types of timeouts.
	Timeout interface{} `json:"timeout" yaml:"timeout"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnRoute_TcpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
}

// An object that represents a target and its relative weight.
//
// Traffic is distributed across targets according to their relative weight. For example, a weighted target with a relative weight of 50 receives five times as much traffic as one with a relative weight of 10. The total weight for all targets combined must be less than or equal to 100.
//
// TODO: EXAMPLE
//
type CfnRoute_WeightedTargetProperty struct {
	// The virtual node to associate with the weighted target.
	VirtualNode *string `json:"virtualNode" yaml:"virtualNode"`
	// The relative weight of the weighted target.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Properties for defining a `CfnRoute`.
//
// TODO: EXAMPLE
//
type CfnRouteProps struct {
	// The name of the service mesh to create the route in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The route specification to apply.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The name of the virtual router in which to create the route.
	//
	// If the virtual router is in a shared mesh, then you must be the owner of the virtual router resource.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// The name to use for the route.
	RouteName *string `json:"routeName" yaml:"routeName"`
	// Optional metadata that you can apply to the route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AppMesh::VirtualGateway`.
//
// Creates a virtual gateway.
//
// A virtual gateway allows resources outside your mesh to communicate to resources that are inside your mesh. The virtual gateway represents an Envoy proxy running in an Amazon ECS task, in a Kubernetes service, or on an Amazon EC2 instance. Unlike a virtual node, which represents an Envoy running with an application, a virtual gateway represents Envoy deployed by itself.
//
// For more information about virtual gateways, see [Virtual gateways](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnVirtualGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnVirtualGateway) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents the methods by which a subject alternative name on a peer Transport Layer Security (TLS) certificate can be matched.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_SubjectAlternativeNameMatchersProperty struct {
	// The values sent must match the specified values exactly.
	Exact *[]*string `json:"exact" yaml:"exact"`
}

// An object that represents the subject alternative names secured by the certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_SubjectAlternativeNamesProperty struct {
	// An object that represents the criteria for determining a SANs match.
	Match interface{} `json:"match" yaml:"match"`
}

// The access log configuration for a virtual gateway.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayAccessLogProperty struct {
	// The file object to send virtual gateway access logs to.
	File interface{} `json:"file" yaml:"file"`
}

// An object that represents the default properties for a backend.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayBackendDefaultsProperty struct {
	// A reference to an object that represents a client policy.
	ClientPolicy interface{} `json:"clientPolicy" yaml:"clientPolicy"`
}

// An object that represents a client policy.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayClientPolicyProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) client policy.
	Tls interface{} `json:"tls" yaml:"tls"`
}

// An object that represents a Transport Layer Security (TLS) client policy.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayClientPolicyTlsProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context.
	Validation interface{} `json:"validation" yaml:"validation"`
	// A reference to an object that represents a virtual gateway's client's Transport Layer Security (TLS) certificate.
	Certificate interface{} `json:"certificate" yaml:"certificate"`
	// Whether the policy is enforced.
	//
	// The default is `True` , if a value isn't specified.
	Enforce interface{} `json:"enforce" yaml:"enforce"`
	// One or more ports that the policy is enforced for.
	Ports interface{} `json:"ports" yaml:"ports"`
}

// An object that represents the virtual gateway's client's Transport Layer Security (TLS) certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayClientTlsCertificateProperty struct {
	// An object that represents a local file certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) .
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's client's Secret Discovery Service certificate.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents the type of virtual gateway connection pool.
//
// Only one protocol is used at a time and should be the same protocol as the one chosen under port mapping.
//
// If not present the default value for `maxPendingRequests` is `2147483647` .
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayConnectionPoolProperty struct {
	// An object that represents a type of connection pool.
	Grpc interface{} `json:"grpc" yaml:"grpc"`
	// An object that represents a type of connection pool.
	Http interface{} `json:"http" yaml:"http"`
	// An object that represents a type of connection pool.
	Http2 interface{} `json:"http2" yaml:"http2"`
}

// An object that represents an access log file.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayFileAccessLogProperty struct {
	// The file path to write access logs to.
	//
	// You can use `/dev/stdout` to send access logs to standard out and configure your Envoy container to use a log driver, such as `awslogs` , to export the access logs to a log storage service such as Amazon CloudWatch Logs. You can also specify a path in the Envoy container's file system to write the files to disk.
	Path *string `json:"path" yaml:"path"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayGrpcConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// An object that represents the health check policy for a virtual gateway's listener.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty struct {
	// The number of consecutive successful health checks that must occur before declaring the listener healthy.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period in milliseconds between each health check execution.
	IntervalMillis *float64 `json:"intervalMillis" yaml:"intervalMillis"`
	// The protocol for the health check request.
	//
	// If you specify `grpc` , then your service must conform to the [GRPC Health Checking Protocol](https://docs.aws.amazon.com/https://github.com/grpc/grpc/blob/master/doc/health-checking.md) .
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The amount of time to wait when receiving a response from the health check, in milliseconds.
	TimeoutMillis *float64 `json:"timeoutMillis" yaml:"timeoutMillis"`
	// The number of consecutive failed health checks that must occur before declaring a virtual gateway unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// The destination path for the health check request.
	//
	// This value is only used if the specified protocol is HTTP or HTTP/2. For any other protocol, this value is ignored.
	Path *string `json:"path" yaml:"path"`
	// The destination port for the health check request.
	//
	// This port must match the port defined in the `PortMapping` for the listener.
	Port *float64 `json:"port" yaml:"port"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayHttp2ConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayHttpConnectionPoolProperty struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster.
	MaxConnections *float64 `json:"maxConnections" yaml:"maxConnections"`
	// Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster.
	MaxPendingRequests *float64 `json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

// An object that represents a listener for a virtual gateway.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerProperty struct {
	// The port mapping information for the listener.
	PortMapping interface{} `json:"portMapping" yaml:"portMapping"`
	// The connection pool information for the listener.
	ConnectionPool interface{} `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck interface{} `json:"healthCheck" yaml:"healthCheck"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for the listener.
	Tls interface{} `json:"tls" yaml:"tls"`
}

// An object that represents an AWS Certificate Manager certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsAcmCertificateProperty struct {
	// The Amazon Resource Name (ARN) for the certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
}

// An object that represents a listener's Transport Layer Security (TLS) certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsCertificateProperty struct {
	// A reference to an object that represents an AWS Certificate Manager certificate.
	Acm interface{} `json:"acm" yaml:"acm"`
	// A reference to an object that represents a local file certificate.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Secret Discovery Service certificate.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents a local file certificate.
//
// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsFileCertificateProperty struct {
	// The certificate chain for the certificate.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
	// The private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.
	PrivateKey *string `json:"privateKey" yaml:"privateKey"`
}

// An object that represents the Transport Layer Security (TLS) properties for a listener.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsProperty struct {
	// An object that represents a Transport Layer Security (TLS) certificate.
	Certificate interface{} `json:"certificate" yaml:"certificate"`
	// Specify one of the following modes.
	//
	// - ** STRICT – Listener only accepts connections with TLS enabled.
	// - ** PERMISSIVE – Listener accepts connections with or without TLS enabled.
	// - ** DISABLED – Listener only accepts connections without TLS.
	Mode *string `json:"mode" yaml:"mode"`
	// A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	Validation interface{} `json:"validation" yaml:"validation"`
}

// An object that represents the virtual gateway's listener's Secret Discovery Service certificate.The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsSdsCertificateProperty struct {
	// A reference to an object that represents the name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a virtual gateway listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context trust.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextTrustProperty struct {
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents logging information.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayLoggingProperty struct {
	// The access log configuration.
	AccessLog interface{} `json:"accessLog" yaml:"accessLog"`
}

// An object that represents a port mapping.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayPortMappingProperty struct {
	// The port used for the port mapping.
	//
	// Specify one protocol.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	Protocol *string `json:"protocol" yaml:"protocol"`
}

// An object that represents the specification of a service mesh resource.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewaySpecProperty struct {
	// The listeners that the mesh endpoint is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `json:"listeners" yaml:"listeners"`
	// A reference to an object that represents the defaults for backends.
	BackendDefaults interface{} `json:"backendDefaults" yaml:"backendDefaults"`
	// An object that represents logging information.
	Logging interface{} `json:"logging" yaml:"logging"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextAcmTrustProperty struct {
	// One or more ACM Amazon Resource Name (ARN)s.
	CertificateAuthorityArns *[]*string `json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextFileTrustProperty struct {
	// The certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
}

// An object that represents a Transport Layer Security (TLS) validation context.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
//
// The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextSdsTrustProperty struct {
	// A reference to an object that represents the name of the secret for a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust.
//
// TODO: EXAMPLE
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextTrustProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
	Acm interface{} `json:"acm" yaml:"acm"`
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// Properties for defining a `CfnVirtualGateway`.
//
// TODO: EXAMPLE
//
type CfnVirtualGatewayProps struct {
	// The name of the service mesh that the virtual gateway resides in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The specifications of the virtual gateway.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual gateway to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The name of the virtual gateway.
	VirtualGatewayName *string `json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

// A CloudFormation `AWS::AppMesh::VirtualNode`.
//
// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment. When you create a virtual node, you can specify the service discovery information for your task group, and whether the proxy running in a task group will communicate with other proxies using Transport Layer Security (TLS).
//
// You define a `listener` for any inbound traffic that your virtual node expects. Any virtual service that your virtual node expects to communicate to is specified as a `backend` .
//
// The response metadata for your new virtual node contains the `arn` that is associated with the virtual node. Set this value to the full ARN; for example, `arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp` ) as the `APPMESH_RESOURCE_ARN` environment variable for your task group's Envoy proxy container in your task definition or pod spec. This is then mapped to the `node.id` and `node.cluster` Envoy parameters.
//
// > By default, App Mesh uses the name of the resource you specified in `APPMESH_RESOURCE_ARN` when Envoy is referring to itself in metrics and traces. You can override this behavior by setting the `APPMESH_RESOURCE_CLUSTER` environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html) . You must be using `1.15.0` or later of the Envoy image when setting these variables. For more information about App Mesh Envoy variables, see [Envoy image](https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html) in the AWS App Mesh User Guide.
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnVirtualNode) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnVirtualNode) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents the access logging information for a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_AccessLogProperty struct {
	// The file object to send virtual node access logs to.
	File interface{} `json:"file" yaml:"file"`
}

// An object that represents the AWS Cloud Map attribute information for your virtual node.
//
// > AWS Cloud Map is not available in the eu-south-1 Region.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_AwsCloudMapInstanceAttributeProperty struct {
	// The name of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	Key *string `json:"key" yaml:"key"`
	// The value of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	Value *string `json:"value" yaml:"value"`
}

// An object that represents the AWS Cloud Map service discovery information for your virtual node.
//
// > AWS Cloud Map is not available in the eu-south-1 Region.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty struct {
	// The name of the AWS Cloud Map namespace to use.
	NamespaceName *string `json:"namespaceName" yaml:"namespaceName"`
	// The name of the AWS Cloud Map service to use.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// A string map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance.
	//
	// Only instances that match all of the specified key/value pairs will be returned.
	Attributes interface{} `json:"attributes" yaml:"attributes"`
}

// An object that represents the default properties for a backend.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_BackendDefaultsProperty struct {
	// A reference to an object that represents a client policy.
	ClientPolicy interface{} `json:"clientPolicy" yaml:"clientPolicy"`
}

// An object that represents the backends that a virtual node is expected to send outbound traffic to.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_BackendProperty struct {
	// Specifies a virtual service to use as a backend.
	VirtualService interface{} `json:"virtualService" yaml:"virtualService"`
}

// An object that represents a client policy.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ClientPolicyProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) client policy.
	Tls interface{} `json:"tls" yaml:"tls"`
}

// A reference to an object that represents a Transport Layer Security (TLS) client policy.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ClientPolicyTlsProperty struct {
	// A reference to an object that represents a TLS validation context.
	Validation interface{} `json:"validation" yaml:"validation"`
	// A reference to an object that represents a client's TLS certificate.
	Certificate interface{} `json:"certificate" yaml:"certificate"`
	// Whether the policy is enforced.
	//
	// The default is `True` , if a value isn't specified.
	Enforce interface{} `json:"enforce" yaml:"enforce"`
	// One or more ports that the policy is enforced for.
	Ports interface{} `json:"ports" yaml:"ports"`
}

// An object that represents the client's certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ClientTlsCertificateProperty struct {
	// An object that represents a local file certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) .
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a client's TLS Secret Discovery Service certificate.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents the DNS service discovery information for your virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_DnsServiceDiscoveryProperty struct {
	// Specifies the DNS service discovery hostname for the virtual node.
	Hostname *string `json:"hostname" yaml:"hostname"`
	// Specifies the DNS response type for the virtual node.
	ResponseType *string `json:"responseType" yaml:"responseType"`
}

// An object that represents a duration of time.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_DurationProperty struct {
	// A unit of time.
	Unit *string `json:"unit" yaml:"unit"`
	// A number of time units.
	Value *float64 `json:"value" yaml:"value"`
}

// An object that represents an access log file.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_FileAccessLogProperty struct {
	// The file path to write access logs to.
	//
	// You can use `/dev/stdout` to send access logs to standard out and configure your Envoy container to use a log driver, such as `awslogs` , to export the access logs to a log storage service such as Amazon CloudWatch Logs. You can also specify a path in the Envoy container's file system to write the files to disk.
	//
	// > The Envoy process must have write permissions to the path that you specify here. Otherwise, Envoy fails to bootstrap properly.
	Path *string `json:"path" yaml:"path"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_GrpcTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
	// An object that represents a per request timeout.
	//
	// The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15 seconds for the source and destination virtual node and the route.
	PerRequest interface{} `json:"perRequest" yaml:"perRequest"`
}

// An object that represents the health check policy for a virtual node's listener.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_HealthCheckProperty struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period in milliseconds between each health check execution.
	IntervalMillis *float64 `json:"intervalMillis" yaml:"intervalMillis"`
	// The protocol for the health check request.
	//
	// If you specify `grpc` , then your service must conform to the [GRPC Health Checking Protocol](https://docs.aws.amazon.com/https://github.com/grpc/grpc/blob/master/doc/health-checking.md) .
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The amount of time to wait when receiving a response from the health check, in milliseconds.
	TimeoutMillis *float64 `json:"timeoutMillis" yaml:"timeoutMillis"`
	// The number of consecutive failed health checks that must occur before declaring a virtual node unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// The destination path for the health check request.
	//
	// This value is only used if the specified protocol is HTTP or HTTP/2. For any other protocol, this value is ignored.
	Path *string `json:"path" yaml:"path"`
	// The destination port for the health check request.
	//
	// This port must match the port defined in the `PortMapping` for the listener.
	Port *float64 `json:"port" yaml:"port"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_HttpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
	// An object that represents a per request timeout.
	//
	// The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15 seconds for the source and destination virtual node and the route.
	PerRequest interface{} `json:"perRequest" yaml:"perRequest"`
}

// An object that represents a listener for a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerProperty struct {
	// The port mapping information for the listener.
	PortMapping interface{} `json:"portMapping" yaml:"portMapping"`
	// The connection pool information for the listener.
	ConnectionPool interface{} `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck interface{} `json:"healthCheck" yaml:"healthCheck"`
	// The outlier detection information for the listener.
	OutlierDetection interface{} `json:"outlierDetection" yaml:"outlierDetection"`
	// An object that represents timeouts for different protocols.
	Timeout interface{} `json:"timeout" yaml:"timeout"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.
	Tls interface{} `json:"tls" yaml:"tls"`
}

// An object that represents timeouts for different protocols.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTimeoutProperty struct {
	// An object that represents types of timeouts.
	Grpc interface{} `json:"grpc" yaml:"grpc"`
	// An object that represents types of timeouts.
	Http interface{} `json:"http" yaml:"http"`
	// An object that represents types of timeouts.
	Http2 interface{} `json:"http2" yaml:"http2"`
	// An object that represents types of timeouts.
	Tcp interface{} `json:"tcp" yaml:"tcp"`
}

// An object that represents an AWS Certificate Manager certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsAcmCertificateProperty struct {
	// The Amazon Resource Name (ARN) for the certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
}

// An object that represents a listener's Transport Layer Security (TLS) certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsCertificateProperty struct {
	// A reference to an object that represents an AWS Certificate Manager certificate.
	Acm interface{} `json:"acm" yaml:"acm"`
	// A reference to an object that represents a local file certificate.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a listener's Secret Discovery Service certificate.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents a local file certificate.
//
// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsFileCertificateProperty struct {
	// The certificate chain for the certificate.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
	// The private key for a certificate stored on the file system of the virtual node that the proxy is running on.
	PrivateKey *string `json:"privateKey" yaml:"privateKey"`
}

// An object that represents the Transport Layer Security (TLS) properties for a listener.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsProperty struct {
	// A reference to an object that represents a listener's Transport Layer Security (TLS) certificate.
	Certificate interface{} `json:"certificate" yaml:"certificate"`
	// Specify one of the following modes.
	//
	// - ** STRICT – Listener only accepts connections with TLS enabled.
	// - ** PERMISSIVE – Listener accepts connections with or without TLS enabled.
	// - ** DISABLED – Listener only accepts connections without TLS.
	Mode *string `json:"mode" yaml:"mode"`
	// A reference to an object that represents a listener's Transport Layer Security (TLS) validation context.
	Validation interface{} `json:"validation" yaml:"validation"`
}

// An object that represents the listener's Secret Discovery Service certificate.
//
// The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsSdsCertificateProperty struct {
	// A reference to an object that represents the name of the secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// An object that represents a listener's Transport Layer Security (TLS) validation context.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// An object that represents a listener's Transport Layer Security (TLS) validation context trust.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ListenerTlsValidationContextTrustProperty struct {
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents the logging information for a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_LoggingProperty struct {
	// The access log configuration for a virtual node.
	AccessLog interface{} `json:"accessLog" yaml:"accessLog"`
}

// An object that represents the outlier detection for a virtual node's listener.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_OutlierDetectionProperty struct {
	// The base amount of time for which a host is ejected.
	BaseEjectionDuration interface{} `json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	Interval interface{} `json:"interval" yaml:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at least one host regardless of the value.
	MaxEjectionPercent *float64 `json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Number of consecutive `5xx` errors required for ejection.
	MaxServerErrors *float64 `json:"maxServerErrors" yaml:"maxServerErrors"`
}

// An object representing a virtual node or virtual router listener port mapping.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_PortMappingProperty struct {
	// The port used for the port mapping.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	//
	// Specify `http` , `http2` , `grpc` , or `tcp` .
	Protocol *string `json:"protocol" yaml:"protocol"`
}

// An object that represents the service discovery information for a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_ServiceDiscoveryProperty struct {
	// Specifies any AWS Cloud Map information for the virtual node.
	AwsCloudMap interface{} `json:"awsCloudMap" yaml:"awsCloudMap"`
	// Specifies the DNS information for the virtual node.
	Dns interface{} `json:"dns" yaml:"dns"`
}

// An object that represents the methods by which a subject alternative name on a peer Transport Layer Security (TLS) certificate can be matched.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_SubjectAlternativeNameMatchersProperty struct {
	// The values sent must match the specified values exactly.
	Exact *[]*string `json:"exact" yaml:"exact"`
}

// An object that represents the subject alternative names secured by the certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_SubjectAlternativeNamesProperty struct {
	// An object that represents the criteria for determining a SANs match.
	Match interface{} `json:"match" yaml:"match"`
}

// An object that represents types of timeouts.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TcpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `json:"idle" yaml:"idle"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TlsValidationContextAcmTrustProperty struct {
	// One or more ACM Amazon Resource Name (ARN)s.
	CertificateAuthorityArns *[]*string `json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TlsValidationContextFileTrustProperty struct {
	// The certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
}

// An object that represents how the proxy will validate its peer during Transport Layer Security (TLS) negotiation.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// An object that represents a Transport Layer Security (TLS) Secret Discovery Service validation context trust.
//
// The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TlsValidationContextSdsTrustProperty struct {
	// A reference to an object that represents the name of the secret for a Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// An object that represents a Transport Layer Security (TLS) validation context trust.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_TlsValidationContextTrustProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
	Acm interface{} `json:"acm" yaml:"acm"`
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `json:"file" yaml:"file"`
	// A reference to an object that represents a Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `json:"sds" yaml:"sds"`
}

// An object that represents the type of virtual node connection pool.
//
// Only one protocol is used at a time and should be the same protocol as the one chosen under port mapping.
//
// If not present the default value for `maxPendingRequests` is `2147483647` .
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeConnectionPoolProperty struct {
	// An object that represents a type of connection pool.
	Grpc interface{} `json:"grpc" yaml:"grpc"`
	// An object that represents a type of connection pool.
	Http interface{} `json:"http" yaml:"http"`
	// An object that represents a type of connection pool.
	Http2 interface{} `json:"http2" yaml:"http2"`
	// An object that represents a type of connection pool.
	Tcp interface{} `json:"tcp" yaml:"tcp"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeGrpcConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeHttp2ConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeHttpConnectionPoolProperty struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster.
	MaxConnections *float64 `json:"maxConnections" yaml:"maxConnections"`
	// Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster.
	MaxPendingRequests *float64 `json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

// An object that represents the specification of a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeSpecProperty struct {
	// A reference to an object that represents the defaults for backends.
	BackendDefaults interface{} `json:"backendDefaults" yaml:"backendDefaults"`
	// The backends that the virtual node is expected to send outbound traffic to.
	Backends interface{} `json:"backends" yaml:"backends"`
	// The listener that the virtual node is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `json:"listeners" yaml:"listeners"`
	// The inbound and outbound access logging information for the virtual node.
	Logging interface{} `json:"logging" yaml:"logging"`
	// The service discovery information for the virtual node.
	//
	// If your virtual node does not expect ingress traffic, you can omit this parameter. If you specify a `listener` , then you must specify service discovery information.
	ServiceDiscovery interface{} `json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

// An object that represents a type of connection pool.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualNodeTcpConnectionPoolProperty struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster.
	MaxConnections *float64 `json:"maxConnections" yaml:"maxConnections"`
}

// An object that represents a virtual service backend for a virtual node.
//
// TODO: EXAMPLE
//
type CfnVirtualNode_VirtualServiceBackendProperty struct {
	// The name of the virtual service that is acting as a virtual node backend.
	VirtualServiceName *string `json:"virtualServiceName" yaml:"virtualServiceName"`
	// A reference to an object that represents the client policy for a backend.
	ClientPolicy interface{} `json:"clientPolicy" yaml:"clientPolicy"`
}

// Properties for defining a `CfnVirtualNode`.
//
// TODO: EXAMPLE
//
type CfnVirtualNodeProps struct {
	// The name of the service mesh to create the virtual node in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The virtual node specification to apply.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual node to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The name to use for the virtual node.
	VirtualNodeName *string `json:"virtualNodeName" yaml:"virtualNodeName"`
}

// A CloudFormation `AWS::AppMesh::VirtualRouter`.
//
// Creates a virtual router within a service mesh.
//
// Specify a `listener` for any inbound traffic that your virtual router receives. Create a virtual router for each protocol and port that you need to route. Virtual routers handle traffic for one or more virtual services within your mesh. After you create your virtual router, create and associate routes for your virtual router that direct incoming requests to different virtual nodes.
//
// For more information about virtual routers, see [Virtual routers](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_routers.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnVirtualRouter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnVirtualRouter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object representing a virtual router listener port mapping.
//
// TODO: EXAMPLE
//
type CfnVirtualRouter_PortMappingProperty struct {
	// The port used for the port mapping.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	//
	// Specify one protocol.
	Protocol *string `json:"protocol" yaml:"protocol"`
}

// An object that represents a virtual router listener.
//
// TODO: EXAMPLE
//
type CfnVirtualRouter_VirtualRouterListenerProperty struct {
	// The port mapping information for the listener.
	PortMapping interface{} `json:"portMapping" yaml:"portMapping"`
}

// An object that represents the specification of a virtual router.
//
// TODO: EXAMPLE
//
type CfnVirtualRouter_VirtualRouterSpecProperty struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `json:"listeners" yaml:"listeners"`
}

// Properties for defining a `CfnVirtualRouter`.
//
// TODO: EXAMPLE
//
type CfnVirtualRouterProps struct {
	// The name of the service mesh to create the virtual router in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The virtual router specification to apply.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual router to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The name to use for the virtual router.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
}

// A CloudFormation `AWS::AppMesh::VirtualService`.
//
// Creates a virtual service within a service mesh.
//
// A virtual service is an abstraction of a real service that is provided by a virtual node directly or indirectly by means of a virtual router. Dependent services call your virtual service by its `virtualServiceName` , and those requests are routed to the virtual node or virtual router that is specified as the provider for the virtual service.
//
// For more information about virtual services, see [Virtual services](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html) .
//
// TODO: EXAMPLE
//
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnVirtualService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
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

func (c *jsiiProxy_CfnVirtualService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents a virtual node service provider.
//
// TODO: EXAMPLE
//
type CfnVirtualService_VirtualNodeServiceProviderProperty struct {
	// The name of the virtual node that is acting as a service provider.
	VirtualNodeName *string `json:"virtualNodeName" yaml:"virtualNodeName"`
}

// An object that represents a virtual node service provider.
//
// TODO: EXAMPLE
//
type CfnVirtualService_VirtualRouterServiceProviderProperty struct {
	// The name of the virtual router that is acting as a service provider.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
}

// An object that represents the provider for a virtual service.
//
// TODO: EXAMPLE
//
type CfnVirtualService_VirtualServiceProviderProperty struct {
	// The virtual node associated with a virtual service.
	VirtualNode interface{} `json:"virtualNode" yaml:"virtualNode"`
	// The virtual router associated with a virtual service.
	VirtualRouter interface{} `json:"virtualRouter" yaml:"virtualRouter"`
}

// An object that represents the specification of a virtual service.
//
// TODO: EXAMPLE
//
type CfnVirtualService_VirtualServiceSpecProperty struct {
	// The App Mesh object that is acting as the provider for a virtual service.
	//
	// You can specify a single virtual node or virtual router.
	Provider interface{} `json:"provider" yaml:"provider"`
}

// Properties for defining a `CfnVirtualService`.
//
// TODO: EXAMPLE
//
type CfnVirtualServiceProps struct {
	// The name of the service mesh to create the virtual service in.
	MeshName *string `json:"meshName" yaml:"meshName"`
	// The virtual service specification to apply.
	Spec interface{} `json:"spec" yaml:"spec"`
	// The name to use for the virtual service.
	VirtualServiceName *string `json:"virtualServiceName" yaml:"virtualServiceName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual service to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Base options for all gateway route specs.
//
// TODO: EXAMPLE
//
type CommonGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Enum of DNS service discovery response type.
//
// TODO: EXAMPLE
//
type DnsResponseType string

const (
	DnsResponseType_LOAD_BALANCER DnsResponseType = "LOAD_BALANCER"
	DnsResponseType_ENDPOINTS DnsResponseType = "ENDPOINTS"
)

// GatewayRoute represents a new or existing gateway route attached to a VirtualGateway and Mesh.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html
//
type GatewayRoute interface {
	awscdk.Resource
	IGatewayRoute
	Env() *awscdk.ResourceEnvironment
	GatewayRouteArn() *string
	GatewayRouteName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualGateway() IVirtualGateway
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_GatewayRoute) Node() constructs.Node {
	var returns constructs.Node
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


func NewGatewayRoute(scope constructs.Construct, id *string, props *GatewayRouteProps) GatewayRoute {
	_init_.Initialize()

	j := jsiiProxy_GatewayRoute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGatewayRoute_Override(g GatewayRoute, scope constructs.Construct, id *string, props *GatewayRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		[]interface{}{scope, id, props},
		g,
	)
}

// Import an existing GatewayRoute given an ARN.
func GatewayRoute_FromGatewayRouteArn(scope constructs.Construct, id *string, gatewayRouteArn *string) IGatewayRoute {
	_init_.Initialize()

	var returns IGatewayRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		"fromGatewayRouteArn",
		[]interface{}{scope, id, gatewayRouteArn},
		&returns,
	)

	return returns
}

// Import an existing GatewayRoute given attributes.
func GatewayRoute_FromGatewayRouteAttributes(scope constructs.Construct, id *string, attrs *GatewayRouteAttributes) IGatewayRoute {
	_init_.Initialize()

	var returns IGatewayRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		"fromGatewayRouteAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GatewayRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GatewayRoute_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
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
func (g *jsiiProxy_GatewayRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Interface with properties necessary to import a reusable GatewayRoute.
//
// TODO: EXAMPLE
//
type GatewayRouteAttributes struct {
	// The name of the GatewayRoute.
	GatewayRouteName *string `json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	VirtualGateway IVirtualGateway `json:"virtualGateway" yaml:"virtualGateway"`
}

// Basic configuration properties for a GatewayRoute.
//
// TODO: EXAMPLE
//
type GatewayRouteBaseProps struct {
	// What protocol the route uses.
	RouteSpec GatewayRouteSpec `json:"routeSpec" yaml:"routeSpec"`
	// The name of the GatewayRoute.
	GatewayRouteName *string `json:"gatewayRouteName" yaml:"gatewayRouteName"`
}

// Used to generate host name matching methods.
//
// TODO: EXAMPLE
//
type GatewayRouteHostnameMatch interface {
	Bind(scope constructs.Construct) *GatewayRouteHostnameMatchConfig
}

// The jsii proxy struct for GatewayRouteHostnameMatch
type jsiiProxy_GatewayRouteHostnameMatch struct {
	_ byte // padding
}

func NewGatewayRouteHostnameMatch_Override(g GatewayRouteHostnameMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		nil, // no parameters
		g,
	)
}

// The value of the host name with the given name must end with the specified characters.
func GatewayRouteHostnameMatch_EndsWith(suffix *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		"endsWith",
		[]interface{}{suffix},
		&returns,
	)

	return returns
}

// The value of the host name must match the specified value exactly.
func GatewayRouteHostnameMatch_Exactly(name *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		"exactly",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the gateway route host name match configuration.
func (g *jsiiProxy_GatewayRouteHostnameMatch) Bind(scope constructs.Construct) *GatewayRouteHostnameMatchConfig {
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
//
// TODO: EXAMPLE
//
type GatewayRouteHostnameMatchConfig struct {
	// GatewayRoute CFN configuration for host name match.
	HostnameMatch *CfnGatewayRoute_GatewayRouteHostnameMatchProperty `json:"hostnameMatch" yaml:"hostnameMatch"`
}

// Properties to define a new GatewayRoute.
//
// TODO: EXAMPLE
//
type GatewayRouteProps struct {
	// What protocol the route uses.
	RouteSpec GatewayRouteSpec `json:"routeSpec" yaml:"routeSpec"`
	// The name of the GatewayRoute.
	GatewayRouteName *string `json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	VirtualGateway IVirtualGateway `json:"virtualGateway" yaml:"virtualGateway"`
}

// Used to generate specs with different protocols for a GatewayRoute.
//
// TODO: EXAMPLE
//
type GatewayRouteSpec interface {
	Bind(scope constructs.Construct) *GatewayRouteSpecConfig
}

// The jsii proxy struct for GatewayRouteSpec
type jsiiProxy_GatewayRouteSpec struct {
	_ byte // padding
}

func NewGatewayRouteSpec_Override(g GatewayRouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		nil, // no parameters
		g,
	)
}

// Creates an gRPC Based GatewayRoute.
func GatewayRouteSpec_Grpc(options *GrpcGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based GatewayRoute.
func GatewayRouteSpec_Http(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based GatewayRoute.
func GatewayRouteSpec_Http2(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
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
func (g *jsiiProxy_GatewayRouteSpec) Bind(scope constructs.Construct) *GatewayRouteSpecConfig {
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
//
// TODO: EXAMPLE
//
type GatewayRouteSpecConfig struct {
	// The spec for a grpc gateway route.
	GrpcSpecConfig *CfnGatewayRoute_GrpcGatewayRouteProperty `json:"grpcSpecConfig" yaml:"grpcSpecConfig"`
	// The spec for an http2 gateway route.
	Http2SpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `json:"http2SpecConfig" yaml:"http2SpecConfig"`
	// The spec for an http gateway route.
	HttpSpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `json:"httpSpecConfig" yaml:"httpSpecConfig"`
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Connection pool properties for gRPC listeners.
//
// TODO: EXAMPLE
//
type GrpcConnectionPool struct {
	// The maximum requests in the pool.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// Represents the properties needed to define GRPC Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
type GrpcGatewayListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *GrpcConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// The criterion for determining a request match for this GatewayRoute.
//
// TODO: EXAMPLE
//
type GrpcGatewayRouteMatch struct {
	// Create host name based gRPC gateway route match.
	Hostname GatewayRouteHostnameMatch `json:"hostname" yaml:"hostname"`
	// Create metadata based gRPC gateway route match.
	//
	// All specified metadata must match for the route to match.
	Metadata *[]HeaderMatch `json:"metadata" yaml:"metadata"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	RewriteRequestHostname *bool `json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
	// Create service name based gRPC gateway route match.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// Properties specific for a gRPC GatewayRoute.
//
// TODO: EXAMPLE
//
type GrpcGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this GatewayRoute.
	Match *GrpcGatewayRouteMatch `json:"match" yaml:"match"`
	// The VirtualService this GatewayRoute directs traffic to.
	RouteTarget IVirtualService `json:"routeTarget" yaml:"routeTarget"`
}

// Properties used to define GRPC Based healthchecks.
//
// TODO: EXAMPLE
//
type GrpcHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period between each health check execution.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// gRPC events.
//
// TODO: EXAMPLE
//
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
type GrpcRetryPolicy struct {
	// The maximum number of retry attempts.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout for each retry attempt.
	RetryTimeout awscdk.Duration `json:"retryTimeout" yaml:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	HttpRetryEvents *[]HttpRetryEvent `json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	TcpRetryEvents *[]TcpRetryEvent `json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
	// gRPC events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	GrpcRetryEvents *[]GrpcRetryEvent `json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
}

// The criterion for determining a request match for this Route.
//
// At least one match type must be selected.
//
// TODO: EXAMPLE
//
type GrpcRouteMatch struct {
	// Create metadata based gRPC route match.
	//
	// All specified metadata must match for the route to match.
	Metadata *[]HeaderMatch `json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If the method name is specified, service name must be also provided.
	MethodName *string `json:"methodName" yaml:"methodName"`
	// Create service name based gRPC route match.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// Properties specific for a GRPC Based Routes.
//
// TODO: EXAMPLE
//
type GrpcRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this Route.
	Match *GrpcRouteMatch `json:"match" yaml:"match"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets" yaml:"weightedTargets"`
	// The retry policy.
	RetryPolicy *GrpcRetryPolicy `json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a grpc timeout.
	Timeout *GrpcTimeout `json:"timeout" yaml:"timeout"`
}

// Represents timeouts for GRPC protocols.
//
// TODO: EXAMPLE
//
type GrpcTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	Idle awscdk.Duration `json:"idle" yaml:"idle"`
	// Represents per request timeout.
	PerRequest awscdk.Duration `json:"perRequest" yaml:"perRequest"`
}

// Represent the GRPC Node Listener prorperty.
//
// TODO: EXAMPLE
//
type GrpcVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *GrpcConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Timeout for GRPC protocol.
	Timeout *GrpcTimeout `json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// Used to generate header matching methods.
//
// TODO: EXAMPLE
//
type HeaderMatch interface {
	Bind(scope constructs.Construct) *HeaderMatchConfig
}

// The jsii proxy struct for HeaderMatch
type jsiiProxy_HeaderMatch struct {
	_ byte // padding
}

func NewHeaderMatch_Override(h HeaderMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		nil, // no parameters
		h,
	)
}

// The value of the header with the given name in the request must not end with the specified characters.
func HeaderMatch_ValueDoesNotEndWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotEndWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not include the specified characters.
func HeaderMatch_ValueDoesNotMatchRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotMatchRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not start with the specified characters.
func HeaderMatch_ValueDoesNotStartWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotStartWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must end with the specified characters.
func HeaderMatch_ValueEndsWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueEndsWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must match the specified value exactly.
func HeaderMatch_ValueIs(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueIs",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not match the specified value exactly.
func HeaderMatch_ValueIsNot(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueIsNot",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must include the specified characters.
func HeaderMatch_ValueMatchesRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueMatchesRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must be in a range of values.
func HeaderMatch_ValuesIsInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valuesIsInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not be in a range of values.
func HeaderMatch_ValuesIsNotInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valuesIsNotInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must start with the specified characters.
func HeaderMatch_ValueStartsWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueStartsWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// Returns the header match configuration.
func (h *jsiiProxy_HeaderMatch) Bind(scope constructs.Construct) *HeaderMatchConfig {
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
//
// TODO: EXAMPLE
//
type HeaderMatchConfig struct {
	// Route CFN configuration for the route header match.
	HeaderMatch *CfnRoute_HttpRouteHeaderProperty `json:"headerMatch" yaml:"headerMatch"`
}

// Contains static factory methods for creating health checks for different protocols.
//
// TODO: EXAMPLE
//
type HealthCheck interface {
	Bind(scope constructs.Construct, options *HealthCheckBindOptions) *HealthCheckConfig
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

func NewHealthCheck_Override(h HealthCheck) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
		nil, // no parameters
		h,
	)
}

// Construct a GRPC health check.
func HealthCheck_Grpc(options *GrpcHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP health check.
func HealthCheck_Http(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP2 health check.
func HealthCheck_Http2(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a TCP health check.
func HealthCheck_Tcp(options *TcpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
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
func (h *jsiiProxy_HealthCheck) Bind(scope constructs.Construct, options *HealthCheckBindOptions) *HealthCheckConfig {
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
//
// TODO: EXAMPLE
//
type HealthCheckBindOptions struct {
	// Port for Health Check interface.
	DefaultPort *float64 `json:"defaultPort" yaml:"defaultPort"`
}

// All Properties for Health Checks for mesh endpoints.
//
// TODO: EXAMPLE
//
type HealthCheckConfig struct {
	// VirtualGateway CFN configuration for Health Checks.
	VirtualGatewayHealthCheck *CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty `json:"virtualGatewayHealthCheck" yaml:"virtualGatewayHealthCheck"`
	// VirtualNode CFN configuration for Health Checks.
	VirtualNodeHealthCheck *CfnVirtualNode_HealthCheckProperty `json:"virtualNodeHealthCheck" yaml:"virtualNodeHealthCheck"`
}

// Connection pool properties for HTTP2 listeners.
//
// TODO: EXAMPLE
//
type Http2ConnectionPool struct {
	// The maximum requests in the pool.
	MaxRequests *float64 `json:"maxRequests" yaml:"maxRequests"`
}

// Represents the properties needed to define HTTP2 Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
type Http2GatewayListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *Http2ConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// Represent the HTTP2 Node Listener prorperty.
//
// TODO: EXAMPLE
//
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	ConnectionPool *Http2ConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	Timeout *HttpTimeout `json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// Connection pool properties for HTTP listeners.
//
// TODO: EXAMPLE
//
type HttpConnectionPool struct {
	// The maximum connections in the pool.
	MaxConnections *float64 `json:"maxConnections" yaml:"maxConnections"`
	// The maximum pending requests in the pool.
	MaxPendingRequests *float64 `json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

// Represents the properties needed to define HTTP Listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
type HttpGatewayListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *HttpConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// The criterion for determining a request match for this GatewayRoute.
//
// TODO: EXAMPLE
//
type HttpGatewayRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the gateway route to match.
	Headers *[]HeaderMatch `json:"headers" yaml:"headers"`
	// The gateway route host name to be matched on.
	Hostname GatewayRouteHostnameMatch `json:"hostname" yaml:"hostname"`
	// The method to match on.
	Method HttpRouteMethod `json:"method" yaml:"method"`
	// Specify how to match requests based on the 'path' part of their URL.
	Path HttpGatewayRoutePathMatch `json:"path" yaml:"path"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	QueryParameters *[]QueryParameterMatch `json:"queryParameters" yaml:"queryParameters"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	RewriteRequestHostname *bool `json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
}

// Defines HTTP gateway route matching based on the URL path of the request.
//
// TODO: EXAMPLE
//
type HttpGatewayRoutePathMatch interface {
	Bind(scope constructs.Construct) *HttpGatewayRoutePathMatchConfig
}

// The jsii proxy struct for HttpGatewayRoutePathMatch
type jsiiProxy_HttpGatewayRoutePathMatch struct {
	_ byte // padding
}

func NewHttpGatewayRoutePathMatch_Override(h HttpGatewayRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
func HttpGatewayRoutePathMatch_Exactly(path *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"exactly",
		[]interface{}{path, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
func HttpGatewayRoutePathMatch_Regex(regex *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"regex",
		[]interface{}{regex, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
func HttpGatewayRoutePathMatch_StartsWith(prefix *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"startsWith",
		[]interface{}{prefix, rewriteTo},
		&returns,
	)

	return returns
}

// Returns the gateway route path match configuration.
func (h *jsiiProxy_HttpGatewayRoutePathMatch) Bind(scope constructs.Construct) *HttpGatewayRoutePathMatchConfig {
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
//
// TODO: EXAMPLE
//
type HttpGatewayRoutePathMatchConfig struct {
	// Gateway route configuration for matching on the prefix of the URL path of the request.
	PrefixPathMatch *string `json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Gateway route configuration for rewriting the prefix of the URL path of the request.
	PrefixPathRewrite *CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty `json:"prefixPathRewrite" yaml:"prefixPathRewrite"`
	// Gateway route configuration for matching on the complete URL path of the request.
	WholePathMatch *CfnGatewayRoute_HttpPathMatchProperty `json:"wholePathMatch" yaml:"wholePathMatch"`
	// Gateway route configuration for rewriting the complete URL path of the request..
	WholePathRewrite *CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty `json:"wholePathRewrite" yaml:"wholePathRewrite"`
}

// Properties specific for HTTP Based GatewayRoutes.
//
// TODO: EXAMPLE
//
type HttpGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The VirtualService this GatewayRoute directs traffic to.
	RouteTarget IVirtualService `json:"routeTarget" yaml:"routeTarget"`
	// The criterion for determining a request match for this GatewayRoute.
	//
	// When path match is defined, this may optionally determine the path rewrite configuration.
	Match *HttpGatewayRouteMatch `json:"match" yaml:"match"`
}

// Properties used to define HTTP Based healthchecks.
//
// TODO: EXAMPLE
//
type HttpHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period between each health check execution.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// The destination path for the health check request.
	Path *string `json:"path" yaml:"path"`
	// The amount of time to wait when receiving a response from the health check.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// HTTP events on which to retry.
//
// TODO: EXAMPLE
//
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
type HttpRetryPolicy struct {
	// The maximum number of retry attempts.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout for each retry attempt.
	RetryTimeout awscdk.Duration `json:"retryTimeout" yaml:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	HttpRetryEvents *[]HttpRetryEvent `json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	TcpRetryEvents *[]TcpRetryEvent `json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

// The criterion for determining a request match for this Route.
//
// TODO: EXAMPLE
//
type HttpRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the route to match.
	Headers *[]HeaderMatch `json:"headers" yaml:"headers"`
	// The HTTP client request method to match on.
	Method HttpRouteMethod `json:"method" yaml:"method"`
	// Specifies how is the request matched based on the path part of its URL.
	Path HttpRoutePathMatch `json:"path" yaml:"path"`
	// The client request protocol to match on.
	//
	// Applicable only for HTTP2 routes.
	Protocol HttpRouteProtocol `json:"protocol" yaml:"protocol"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	QueryParameters *[]QueryParameterMatch `json:"queryParameters" yaml:"queryParameters"`
}

// Supported values for matching routes based on the HTTP request method.
//
// TODO: EXAMPLE
//
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
type HttpRoutePathMatch interface {
	Bind(scope constructs.Construct) *HttpRoutePathMatchConfig
}

// The jsii proxy struct for HttpRoutePathMatch
type jsiiProxy_HttpRoutePathMatch struct {
	_ byte // padding
}

func NewHttpRoutePathMatch_Override(h HttpRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
func HttpRoutePathMatch_Exactly(path *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"exactly",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
func HttpRoutePathMatch_Regex(regex *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"regex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
func HttpRoutePathMatch_StartsWith(prefix *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"startsWith",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the route path match configuration.
func (h *jsiiProxy_HttpRoutePathMatch) Bind(scope constructs.Construct) *HttpRoutePathMatchConfig {
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
//
// TODO: EXAMPLE
//
type HttpRoutePathMatchConfig struct {
	// Route configuration for matching on the prefix of the URL path of the request.
	PrefixPathMatch *string `json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Route configuration for matching on the complete URL path of the request.
	WholePathMatch *CfnRoute_HttpPathMatchProperty `json:"wholePathMatch" yaml:"wholePathMatch"`
}

// Supported :scheme options for HTTP2.
//
// TODO: EXAMPLE
//
type HttpRouteProtocol string

const (
	HttpRouteProtocol_HTTP HttpRouteProtocol = "HTTP"
	HttpRouteProtocol_HTTPS HttpRouteProtocol = "HTTPS"
)

// Properties specific for HTTP Based Routes.
//
// TODO: EXAMPLE
//
type HttpRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets" yaml:"weightedTargets"`
	// The criterion for determining a request match for this Route.
	Match *HttpRouteMatch `json:"match" yaml:"match"`
	// The retry policy.
	RetryPolicy *HttpRetryPolicy `json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a http timeout.
	Timeout *HttpTimeout `json:"timeout" yaml:"timeout"`
}

// Represents timeouts for HTTP protocols.
//
// TODO: EXAMPLE
//
type HttpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	Idle awscdk.Duration `json:"idle" yaml:"idle"`
	// Represents per request timeout.
	PerRequest awscdk.Duration `json:"perRequest" yaml:"perRequest"`
}

// Represent the HTTP Node Listener prorperty.
//
// TODO: EXAMPLE
//
type HttpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *HttpConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	Timeout *HttpTimeout `json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// Interface for which all GatewayRoute based classes MUST implement.
type IGatewayRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the GatewayRoute.
	GatewayRouteArn() *string
	// The name of the GatewayRoute.
	GatewayRouteName() *string
	// The VirtualGateway the GatewayRoute belongs to.
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
type IMesh interface {
	awscdk.IResource
	// Creates a new VirtualGateway in this Mesh.
	//
	// Note that the Gateway is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	// Creates a new VirtualNode in this Mesh.
	//
	// Note that the Node is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	// Creates a new VirtualRouter in this Mesh.
	//
	// Note that the Router is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
	// The Amazon Resource Name (ARN) of the AppMesh mesh.
	MeshArn() *string
	// The name of the AppMesh mesh.
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
type IRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the route.
	RouteArn() *string
	// The name of the route.
	RouteName() *string
	// The VirtualRouter the Route belongs to.
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
type IVirtualGateway interface {
	awscdk.IResource
	// Utility method to add a new GatewayRoute to the VirtualGateway.
	AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualGateway belongs to.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualGateway.
	VirtualGatewayArn() *string
	// Name of the VirtualGateway.
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
type IVirtualNode interface {
	awscdk.IResource
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualNode belongs to.
	Mesh() IMesh
	// The Amazon Resource Name belonging to the VirtualNode.
	//
	// Set this value as the APPMESH_VIRTUAL_NODE_NAME environment variable for
	// your task group's Envoy proxy container in your task definition or pod
	// spec.
	VirtualNodeArn() *string
	// The name of the VirtualNode.
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
type IVirtualRouter interface {
	awscdk.IResource
	// Add a single route to the router.
	AddRoute(id *string, props *RouteBaseProps) Route
	// The Mesh which the VirtualRouter belongs to.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualRouter.
	VirtualRouterArn() *string
	// The name of the VirtualRouter.
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
type IVirtualService interface {
	awscdk.IResource
	// The Mesh which the VirtualService belongs to.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the virtual service.
	VirtualServiceArn() *string
	// The name of the VirtualService.
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
type ListenerTlsOptions struct {
	// Represents TLS certificate.
	Certificate TlsCertificate `json:"certificate" yaml:"certificate"`
	// The TLS mode.
	Mode TlsMode `json:"mode" yaml:"mode"`
	// Represents a listener's TLS validation context.
	//
	// The client certificate will only be validated if the client provides it, enabling mutual TLS.
	MutualTlsValidation *MutualTlsValidation `json:"mutualTlsValidation" yaml:"mutualTlsValidation"`
}

// Define a new AppMesh mesh.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html
//
type Mesh interface {
	awscdk.Resource
	IMesh
	Env() *awscdk.ResourceEnvironment
	MeshArn() *string
	MeshName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_Mesh) Node() constructs.Node {
	var returns constructs.Node
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


func NewMesh(scope constructs.Construct, id *string, props *MeshProps) Mesh {
	_init_.Initialize()

	j := jsiiProxy_Mesh{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMesh_Override(m Mesh, scope constructs.Construct, id *string, props *MeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import an existing mesh by arn.
func Mesh_FromMeshArn(scope constructs.Construct, id *string, meshArn *string) IMesh {
	_init_.Initialize()

	var returns IMesh

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"fromMeshArn",
		[]interface{}{scope, id, meshArn},
		&returns,
	)

	return returns
}

// Import an existing mesh by name.
func Mesh_FromMeshName(scope constructs.Construct, id *string, meshName *string) IMesh {
	_init_.Initialize()

	var returns IMesh

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"fromMeshName",
		[]interface{}{scope, id, meshName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Mesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Mesh_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a VirtualGateway to the Mesh.
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
func (m *jsiiProxy_Mesh) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// A utility enum defined for the egressFilter type property, the default of DROP_ALL, allows traffic only to other resources inside the mesh, or API calls to amazon resources.
//
// TODO: EXAMPLE
//
type MeshFilterType string

const (
	MeshFilterType_ALLOW_ALL MeshFilterType = "ALLOW_ALL"
	MeshFilterType_DROP_ALL MeshFilterType = "DROP_ALL"
)

// The set of properties used when creating a Mesh.
//
// TODO: EXAMPLE
//
type MeshProps struct {
	// Egress filter to be applied to the Mesh.
	EgressFilter MeshFilterType `json:"egressFilter" yaml:"egressFilter"`
	// The name of the Mesh being defined.
	MeshName *string `json:"meshName" yaml:"meshName"`
}

// Represents a TLS certificate that is supported for mutual TLS authentication.
//
// TODO: EXAMPLE
//
type MutualTlsCertificate interface {
	TlsCertificate
	Differentiator() *bool
	Bind(_scope constructs.Construct) *TlsCertificateConfig
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


func NewMutualTlsCertificate_Override(m MutualTlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		nil, // no parameters
		m,
	)
}

// Returns an ACM TLS Certificate.
func MutualTlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
func MutualTlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
func MutualTlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns TLS certificate based provider.
func (m *jsiiProxy_MutualTlsCertificate) Bind(_scope constructs.Construct) *TlsCertificateConfig {
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
type MutualTlsValidation struct {
	// Reference to where to retrieve the trust chain.
	Trust MutualTlsValidationTrust `json:"trust" yaml:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	SubjectAlternativeNames SubjectAlternativeNames `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// Represents a TLS Validation Context Trust that is supported for mutual TLS authentication.
//
// TODO: EXAMPLE
//
type MutualTlsValidationTrust interface {
	TlsValidationTrust
	Differentiator() *bool
	Bind(scope constructs.Construct) *TlsValidationTrustConfig
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


func NewMutualTlsValidationTrust_Override(m MutualTlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		nil, // no parameters
		m,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
func MutualTlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
func MutualTlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
func MutualTlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns Trust context based on trust type.
func (m *jsiiProxy_MutualTlsValidationTrust) Bind(scope constructs.Construct) *TlsValidationTrustConfig {
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
type OutlierDetection struct {
	// The base amount of time for which a host is ejected.
	BaseEjectionDuration awscdk.Duration `json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at
	// least one host regardless of the value.
	MaxEjectionPercent *float64 `json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Number of consecutive 5xx errors required for ejection.
	MaxServerErrors *float64 `json:"maxServerErrors" yaml:"maxServerErrors"`
}

// Used to generate query parameter matching methods.
//
// TODO: EXAMPLE
//
type QueryParameterMatch interface {
	Bind(scope constructs.Construct) *QueryParameterMatchConfig
}

// The jsii proxy struct for QueryParameterMatch
type jsiiProxy_QueryParameterMatch struct {
	_ byte // padding
}

func NewQueryParameterMatch_Override(q QueryParameterMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatch",
		nil, // no parameters
		q,
	)
}

// The value of the query parameter with the given name in the request must match the specified value exactly.
func QueryParameterMatch_ValueIs(queryParameterName *string, queryParameterValue *string) QueryParameterMatch {
	_init_.Initialize()

	var returns QueryParameterMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatch",
		"valueIs",
		[]interface{}{queryParameterName, queryParameterValue},
		&returns,
	)

	return returns
}

// Returns the query parameter match configuration.
func (q *jsiiProxy_QueryParameterMatch) Bind(scope constructs.Construct) *QueryParameterMatchConfig {
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
//
// TODO: EXAMPLE
//
type QueryParameterMatchConfig struct {
	// Route CFN configuration for route query parameter match.
	QueryParameterMatch *CfnRoute_QueryParameterProperty `json:"queryParameterMatch" yaml:"queryParameterMatch"`
}

// Route represents a new or existing route attached to a VirtualRouter and Mesh.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html
//
type Route interface {
	awscdk.Resource
	IRoute
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	RouteArn() *string
	RouteName() *string
	Stack() awscdk.Stack
	VirtualRouter() IVirtualRouter
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_Route) Node() constructs.Node {
	var returns constructs.Node
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


func NewRoute(scope constructs.Construct, id *string, props *RouteProps) Route {
	_init_.Initialize()

	j := jsiiProxy_Route{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRoute_Override(r Route, scope constructs.Construct, id *string, props *RouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing Route given an ARN.
func Route_FromRouteArn(scope constructs.Construct, id *string, routeArn *string) IRoute {
	_init_.Initialize()

	var returns IRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"fromRouteArn",
		[]interface{}{scope, id, routeArn},
		&returns,
	)

	return returns
}

// Import an existing Route given attributes.
func Route_FromRouteAttributes(scope constructs.Construct, id *string, attrs *RouteAttributes) IRoute {
	_init_.Initialize()

	var returns IRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"fromRouteAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Route_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Route_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
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
func (r *jsiiProxy_Route) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Interface with properties ncecessary to import a reusable Route.
//
// TODO: EXAMPLE
//
type RouteAttributes struct {
	// The name of the Route.
	RouteName *string `json:"routeName" yaml:"routeName"`
	// The VirtualRouter the Route belongs to.
	VirtualRouter IVirtualRouter `json:"virtualRouter" yaml:"virtualRouter"`
}

// Base interface properties for all Routes.
//
// TODO: EXAMPLE
//
type RouteBaseProps struct {
	// Protocol specific spec.
	RouteSpec RouteSpec `json:"routeSpec" yaml:"routeSpec"`
	// The name of the route.
	RouteName *string `json:"routeName" yaml:"routeName"`
}

// Properties to define new Routes.
//
// TODO: EXAMPLE
//
type RouteProps struct {
	// Protocol specific spec.
	RouteSpec RouteSpec `json:"routeSpec" yaml:"routeSpec"`
	// The name of the route.
	RouteName *string `json:"routeName" yaml:"routeName"`
	// The service mesh to define the route in.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// The VirtualRouter the Route belongs to.
	VirtualRouter IVirtualRouter `json:"virtualRouter" yaml:"virtualRouter"`
}

// Used to generate specs with different protocols for a RouteSpec.
//
// TODO: EXAMPLE
//
type RouteSpec interface {
	Bind(scope constructs.Construct) *RouteSpecConfig
}

// The jsii proxy struct for RouteSpec
type jsiiProxy_RouteSpec struct {
	_ byte // padding
}

func NewRouteSpec_Override(r RouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
		nil, // no parameters
		r,
	)
}

// Creates a GRPC Based RouteSpec.
func RouteSpec_Grpc(options *GrpcRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based RouteSpec.
func RouteSpec_Http(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based RouteSpec.
func RouteSpec_Http2(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a TCP Based RouteSpec.
func RouteSpec_Tcp(options *TcpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
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
func (r *jsiiProxy_RouteSpec) Bind(scope constructs.Construct) *RouteSpecConfig {
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
//
// TODO: EXAMPLE
//
type RouteSpecConfig struct {
	// The spec for a grpc route.
	GrpcRouteSpec *CfnRoute_GrpcRouteProperty `json:"grpcRouteSpec" yaml:"grpcRouteSpec"`
	// The spec for an http2 route.
	Http2RouteSpec *CfnRoute_HttpRouteProperty `json:"http2RouteSpec" yaml:"http2RouteSpec"`
	// The spec for an http route.
	HttpRouteSpec *CfnRoute_HttpRouteProperty `json:"httpRouteSpec" yaml:"httpRouteSpec"`
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The spec for a tcp route.
	TcpRouteSpec *CfnRoute_TcpRouteProperty `json:"tcpRouteSpec" yaml:"tcpRouteSpec"`
}

// Base options for all route specs.
//
// TODO: EXAMPLE
//
type RouteSpecOptionsBase struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Provides the Service Discovery method a VirtualNode uses.
//
// TODO: EXAMPLE
//
type ServiceDiscovery interface {
	Bind(scope constructs.Construct) *ServiceDiscoveryConfig
}

// The jsii proxy struct for ServiceDiscovery
type jsiiProxy_ServiceDiscovery struct {
	_ byte // padding
}

func NewServiceDiscovery_Override(s ServiceDiscovery) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		nil, // no parameters
		s,
	)
}

// Returns Cloud Map based service discovery.
func ServiceDiscovery_CloudMap(service awsservicediscovery.IService, instanceAttributes *map[string]*string) ServiceDiscovery {
	_init_.Initialize()

	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		"cloudMap",
		[]interface{}{service, instanceAttributes},
		&returns,
	)

	return returns
}

// Returns DNS based service discovery.
func ServiceDiscovery_Dns(hostname *string, responseType DnsResponseType) ServiceDiscovery {
	_init_.Initialize()

	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		"dns",
		[]interface{}{hostname, responseType},
		&returns,
	)

	return returns
}

// Binds the current object when adding Service Discovery to a VirtualNode.
func (s *jsiiProxy_ServiceDiscovery) Bind(scope constructs.Construct) *ServiceDiscoveryConfig {
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
//
// TODO: EXAMPLE
//
type ServiceDiscoveryConfig struct {
	// Cloud Map based Service Discovery.
	Cloudmap *CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty `json:"cloudmap" yaml:"cloudmap"`
	// DNS based Service Discovery.
	Dns *CfnVirtualNode_DnsServiceDiscoveryProperty `json:"dns" yaml:"dns"`
}

// Used to generate Subject Alternative Names Matchers.
//
// TODO: EXAMPLE
//
type SubjectAlternativeNames interface {
	Bind(scope constructs.Construct) *SubjectAlternativeNamesMatcherConfig
}

// The jsii proxy struct for SubjectAlternativeNames
type jsiiProxy_SubjectAlternativeNames struct {
	_ byte // padding
}

func NewSubjectAlternativeNames_Override(s SubjectAlternativeNames) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNames",
		nil, // no parameters
		s,
	)
}

// The values of the SAN must match the specified values exactly.
func SubjectAlternativeNames_MatchingExactly(names ...*string) SubjectAlternativeNames {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range names {
		args = append(args, a)
	}

	var returns SubjectAlternativeNames

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNames",
		"matchingExactly",
		args,
		&returns,
	)

	return returns
}

// Returns Subject Alternative Names Matcher based on method type.
func (s *jsiiProxy_SubjectAlternativeNames) Bind(scope constructs.Construct) *SubjectAlternativeNamesMatcherConfig {
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
//
// TODO: EXAMPLE
//
type SubjectAlternativeNamesMatcherConfig struct {
	// VirtualNode CFN configuration for subject alternative names secured by the certificate.
	SubjectAlternativeNamesMatch *CfnVirtualNode_SubjectAlternativeNameMatchersProperty `json:"subjectAlternativeNamesMatch" yaml:"subjectAlternativeNamesMatch"`
}

// Connection pool properties for TCP listeners.
//
// TODO: EXAMPLE
//
type TcpConnectionPool struct {
	// The maximum connections in the pool.
	MaxConnections *float64 `json:"maxConnections" yaml:"maxConnections"`
}

// Properties used to define TCP Based healthchecks.
//
// TODO: EXAMPLE
//
type TcpHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period between each health check execution.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// TCP events on which you may retry.
//
// TODO: EXAMPLE
//
type TcpRetryEvent string

const (
	TcpRetryEvent_CONNECTION_ERROR TcpRetryEvent = "CONNECTION_ERROR"
)

// Properties specific for a TCP Based Routes.
//
// TODO: EXAMPLE
//
type TcpRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `json:"priority" yaml:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `json:"weightedTargets" yaml:"weightedTargets"`
	// An object that represents a tcp timeout.
	Timeout *TcpTimeout `json:"timeout" yaml:"timeout"`
}

// Represents timeouts for TCP protocols.
//
// TODO: EXAMPLE
//
type TcpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	Idle awscdk.Duration `json:"idle" yaml:"idle"`
}

// Represent the TCP Node Listener prorperty.
//
// TODO: EXAMPLE
//
type TcpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *TcpConnectionPool `json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `json:"port" yaml:"port"`
	// Timeout for TCP protocol.
	Timeout *TcpTimeout `json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `json:"tls" yaml:"tls"`
}

// Represents a TLS certificate.
//
// TODO: EXAMPLE
//
type TlsCertificate interface {
	Bind(_scope constructs.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for TlsCertificate
type jsiiProxy_TlsCertificate struct {
	_ byte // padding
}

func NewTlsCertificate_Override(t TlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		nil, // no parameters
		t,
	)
}

// Returns an ACM TLS Certificate.
func TlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
func TlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
func TlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns TLS certificate based provider.
func (t *jsiiProxy_TlsCertificate) Bind(_scope constructs.Construct) *TlsCertificateConfig {
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
//
// TODO: EXAMPLE
//
type TlsCertificateConfig struct {
	// The CFN shape for a TLS certificate.
	TlsCertificate *CfnVirtualNode_ListenerTlsCertificateProperty `json:"tlsCertificate" yaml:"tlsCertificate"`
}

// Represents the properties needed to define client policy.
//
// TODO: EXAMPLE
//
type TlsClientPolicy struct {
	// Represents the object for TLS validation context.
	Validation *TlsValidation `json:"validation" yaml:"validation"`
	// Whether the policy is enforced.
	Enforce *bool `json:"enforce" yaml:"enforce"`
	// Represents a client TLS certificate.
	//
	// The certificate will be sent only if the server requests it, enabling mutual TLS.
	MutualTlsCertificate MutualTlsCertificate `json:"mutualTlsCertificate" yaml:"mutualTlsCertificate"`
	// TLS is enforced on the ports specified here.
	//
	// If no ports are specified, TLS will be enforced on all the ports.
	Ports *[]*float64 `json:"ports" yaml:"ports"`
}

// Enum of supported TLS modes.
//
// TODO: EXAMPLE
//
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
type TlsValidation struct {
	// Reference to where to retrieve the trust chain.
	Trust TlsValidationTrust `json:"trust" yaml:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	SubjectAlternativeNames SubjectAlternativeNames `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// Defines the TLS Validation Context Trust.
//
// TODO: EXAMPLE
//
type TlsValidationTrust interface {
	Bind(scope constructs.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for TlsValidationTrust
type jsiiProxy_TlsValidationTrust struct {
	_ byte // padding
}

func NewTlsValidationTrust_Override(t TlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		nil, // no parameters
		t,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
func TlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
func TlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
func TlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

// Returns Trust context based on trust type.
func (t *jsiiProxy_TlsValidationTrust) Bind(scope constructs.Construct) *TlsValidationTrustConfig {
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
//
// TODO: EXAMPLE
//
type TlsValidationTrustConfig struct {
	// VirtualNode CFN configuration for client policy's TLS Validation Trust.
	TlsValidationTrust *CfnVirtualNode_TlsValidationContextTrustProperty `json:"tlsValidationTrust" yaml:"tlsValidationTrust"`
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
type VirtualGateway interface {
	awscdk.Resource
	IVirtualGateway
	Env() *awscdk.ResourceEnvironment
	Listeners() *[]*VirtualGatewayListenerConfig
	Mesh() IMesh
	Node() constructs.Node
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
	ToString() *string
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

func (j *jsiiProxy_VirtualGateway) Node() constructs.Node {
	var returns constructs.Node
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


func NewVirtualGateway(scope constructs.Construct, id *string, props *VirtualGatewayProps) VirtualGateway {
	_init_.Initialize()

	j := jsiiProxy_VirtualGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVirtualGateway_Override(v VirtualGateway, scope constructs.Construct, id *string, props *VirtualGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualGateway given an ARN.
func VirtualGateway_FromVirtualGatewayArn(scope constructs.Construct, id *string, virtualGatewayArn *string) IVirtualGateway {
	_init_.Initialize()

	var returns IVirtualGateway

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		"fromVirtualGatewayArn",
		[]interface{}{scope, id, virtualGatewayArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualGateway given its attributes.
func VirtualGateway_FromVirtualGatewayAttributes(scope constructs.Construct, id *string, attrs *VirtualGatewayAttributes) IVirtualGateway {
	_init_.Initialize()

	var returns IVirtualGateway

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		"fromVirtualGatewayAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VirtualGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VirtualGateway_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Utility method to add a new GatewayRoute to the VirtualGateway.
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
func (v *jsiiProxy_VirtualGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Unterface with properties necessary to import a reusable VirtualGateway.
//
// TODO: EXAMPLE
//
type VirtualGatewayAttributes struct {
	// The Mesh that the VirtualGateway belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// The name of the VirtualGateway.
	VirtualGatewayName *string `json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

// Basic configuration properties for a VirtualGateway.
//
// TODO: EXAMPLE
//
type VirtualGatewayBaseProps struct {
	// Access Logging Configuration for the VirtualGateway.
	AccessLog AccessLog `json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	Listeners *[]VirtualGatewayListener `json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	VirtualGatewayName *string `json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

// Represents the properties needed to define listeners for a VirtualGateway.
//
// TODO: EXAMPLE
//
type VirtualGatewayListener interface {
	Bind(scope constructs.Construct) *VirtualGatewayListenerConfig
}

// The jsii proxy struct for VirtualGatewayListener
type jsiiProxy_VirtualGatewayListener struct {
	_ byte // padding
}

func NewVirtualGatewayListener_Override(v VirtualGatewayListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualGateway.
func VirtualGatewayListener_Grpc(options *GrpcGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualGateway.
func VirtualGatewayListener_Http(options *HttpGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualGateway.
func VirtualGatewayListener_Http2(options *Http2GatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
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
func (v *jsiiProxy_VirtualGatewayListener) Bind(scope constructs.Construct) *VirtualGatewayListenerConfig {
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
//
// TODO: EXAMPLE
//
type VirtualGatewayListenerConfig struct {
	// Single listener config for a VirtualGateway.
	Listener *CfnVirtualGateway_VirtualGatewayListenerProperty `json:"listener" yaml:"listener"`
}

// Properties used when creating a new VirtualGateway.
//
// TODO: EXAMPLE
//
type VirtualGatewayProps struct {
	// Access Logging Configuration for the VirtualGateway.
	AccessLog AccessLog `json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	Listeners *[]VirtualGatewayListener `json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	VirtualGatewayName *string `json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The Mesh which the VirtualGateway belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
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
type VirtualNode interface {
	awscdk.Resource
	IVirtualNode
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() constructs.Node
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
	ToString() *string
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

func (j *jsiiProxy_VirtualNode) Node() constructs.Node {
	var returns constructs.Node
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


func NewVirtualNode(scope constructs.Construct, id *string, props *VirtualNodeProps) VirtualNode {
	_init_.Initialize()

	j := jsiiProxy_VirtualNode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVirtualNode_Override(v VirtualNode, scope constructs.Construct, id *string, props *VirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualNode given an ARN.
func VirtualNode_FromVirtualNodeArn(scope constructs.Construct, id *string, virtualNodeArn *string) IVirtualNode {
	_init_.Initialize()

	var returns IVirtualNode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		"fromVirtualNodeArn",
		[]interface{}{scope, id, virtualNodeArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualNode given its name.
func VirtualNode_FromVirtualNodeAttributes(scope constructs.Construct, id *string, attrs *VirtualNodeAttributes) IVirtualNode {
	_init_.Initialize()

	var returns IVirtualNode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		"fromVirtualNodeAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VirtualNode_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a Virtual Services that this node is expected to send outbound traffic to.
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
func (v *jsiiProxy_VirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Interface with properties necessary to import a reusable VirtualNode.
//
// TODO: EXAMPLE
//
type VirtualNodeAttributes struct {
	// The Mesh that the VirtualNode belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// The name of the VirtualNode.
	VirtualNodeName *string `json:"virtualNodeName" yaml:"virtualNodeName"`
}

// Basic configuration properties for a VirtualNode.
//
// TODO: EXAMPLE
//
type VirtualNodeBaseProps struct {
	// Access Logging Configuration for the virtual node.
	AccessLog AccessLog `json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `json:"backendDefaults" yaml:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	Backends *[]Backend `json:"backends" yaml:"backends"`
	// Initial listener for the virtual node.
	Listeners *[]VirtualNodeListener `json:"listeners" yaml:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	ServiceDiscovery ServiceDiscovery `json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// The name of the VirtualNode.
	VirtualNodeName *string `json:"virtualNodeName" yaml:"virtualNodeName"`
}

// Defines listener for a VirtualNode.
//
// TODO: EXAMPLE
//
type VirtualNodeListener interface {
	Bind(scope constructs.Construct) *VirtualNodeListenerConfig
}

// The jsii proxy struct for VirtualNodeListener
type jsiiProxy_VirtualNodeListener struct {
	_ byte // padding
}

func NewVirtualNodeListener_Override(v VirtualNodeListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		nil, // no parameters
		v,
	)
}

// Returns an GRPC Listener for a VirtualNode.
func VirtualNodeListener_Grpc(props *GrpcVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"grpc",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualNode.
func VirtualNodeListener_Http(props *HttpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"http",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualNode.
func VirtualNodeListener_Http2(props *Http2VirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"http2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an TCP Listener for a VirtualNode.
func VirtualNodeListener_Tcp(props *TcpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"tcp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Binds the current object when adding Listener to a VirtualNode.
func (v *jsiiProxy_VirtualNodeListener) Bind(scope constructs.Construct) *VirtualNodeListenerConfig {
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
//
// TODO: EXAMPLE
//
type VirtualNodeListenerConfig struct {
	// Single listener config for a VirtualNode.
	Listener *CfnVirtualNode_ListenerProperty `json:"listener" yaml:"listener"`
}

// The properties used when creating a new VirtualNode.
//
// TODO: EXAMPLE
//
type VirtualNodeProps struct {
	// Access Logging Configuration for the virtual node.
	AccessLog AccessLog `json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `json:"backendDefaults" yaml:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	Backends *[]Backend `json:"backends" yaml:"backends"`
	// Initial listener for the virtual node.
	Listeners *[]VirtualNodeListener `json:"listeners" yaml:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	ServiceDiscovery ServiceDiscovery `json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// The name of the VirtualNode.
	VirtualNodeName *string `json:"virtualNodeName" yaml:"virtualNodeName"`
	// The Mesh which the VirtualNode belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
}

// TODO: EXAMPLE
//
type VirtualRouter interface {
	awscdk.Resource
	IVirtualRouter
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualRouterArn() *string
	VirtualRouterName() *string
	AddRoute(id *string, props *RouteBaseProps) Route
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_VirtualRouter) Node() constructs.Node {
	var returns constructs.Node
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


func NewVirtualRouter(scope constructs.Construct, id *string, props *VirtualRouterProps) VirtualRouter {
	_init_.Initialize()

	j := jsiiProxy_VirtualRouter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVirtualRouter_Override(v VirtualRouter, scope constructs.Construct, id *string, props *VirtualRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualRouter given an ARN.
func VirtualRouter_FromVirtualRouterArn(scope constructs.Construct, id *string, virtualRouterArn *string) IVirtualRouter {
	_init_.Initialize()

	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"fromVirtualRouterArn",
		[]interface{}{scope, id, virtualRouterArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualRouter given attributes.
func VirtualRouter_FromVirtualRouterAttributes(scope constructs.Construct, id *string, attrs *VirtualRouterAttributes) IVirtualRouter {
	_init_.Initialize()

	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"fromVirtualRouterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VirtualRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VirtualRouter_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a single route to the router.
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
func (v *jsiiProxy_VirtualRouter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Interface with properties ncecessary to import a reusable VirtualRouter.
//
// TODO: EXAMPLE
//
type VirtualRouterAttributes struct {
	// The Mesh which the VirtualRouter belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
}

// Interface with base properties all routers willl inherit.
//
// TODO: EXAMPLE
//
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
}

// Represents the properties needed to define listeners for a VirtualRouter.
//
// TODO: EXAMPLE
//
type VirtualRouterListener interface {
	Bind(scope constructs.Construct) *VirtualRouterListenerConfig
}

// The jsii proxy struct for VirtualRouterListener
type jsiiProxy_VirtualRouterListener struct {
	_ byte // padding
}

func NewVirtualRouterListener_Override(v VirtualRouterListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualRouter.
func VirtualRouterListener_Grpc(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"grpc",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualRouter.
func VirtualRouterListener_Http(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"http",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualRouter.
func VirtualRouterListener_Http2(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"http2",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns a TCP Listener for a VirtualRouter.
func VirtualRouterListener_Tcp(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
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
func (v *jsiiProxy_VirtualRouterListener) Bind(scope constructs.Construct) *VirtualRouterListenerConfig {
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
//
// TODO: EXAMPLE
//
type VirtualRouterListenerConfig struct {
	// Single listener config for a VirtualRouter.
	Listener *CfnVirtualRouter_VirtualRouterListenerProperty `json:"listener" yaml:"listener"`
}

// The properties used when creating a new VirtualRouter.
//
// TODO: EXAMPLE
//
type VirtualRouterProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `json:"virtualRouterName" yaml:"virtualRouterName"`
	// The Mesh which the VirtualRouter belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
}

// VirtualService represents a service inside an AppMesh.
//
// It routes traffic either to a Virtual Node or to a Virtual Router.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html
//
type VirtualService interface {
	awscdk.Resource
	IVirtualService
	Env() *awscdk.ResourceEnvironment
	Mesh() IMesh
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	VirtualServiceArn() *string
	VirtualServiceName() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_VirtualService) Node() constructs.Node {
	var returns constructs.Node
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


func NewVirtualService(scope constructs.Construct, id *string, props *VirtualServiceProps) VirtualService {
	_init_.Initialize()

	j := jsiiProxy_VirtualService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVirtualService_Override(v VirtualService, scope constructs.Construct, id *string, props *VirtualServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualService given an ARN.
func VirtualService_FromVirtualServiceArn(scope constructs.Construct, id *string, virtualServiceArn *string) IVirtualService {
	_init_.Initialize()

	var returns IVirtualService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		"fromVirtualServiceArn",
		[]interface{}{scope, id, virtualServiceArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualService given its attributes.
func VirtualService_FromVirtualServiceAttributes(scope constructs.Construct, id *string, attrs *VirtualServiceAttributes) IVirtualService {
	_init_.Initialize()

	var returns IVirtualService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		"fromVirtualServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VirtualService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VirtualService_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualService",
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
func (v *jsiiProxy_VirtualService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// Returns a string representation of this construct.
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

// Interface with properties ncecessary to import a reusable VirtualService.
//
// TODO: EXAMPLE
//
type VirtualServiceAttributes struct {
	// The Mesh which the VirtualService belongs to.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// The name of the VirtualService, it is recommended this follows the fully-qualified domain name format.
	VirtualServiceName *string `json:"virtualServiceName" yaml:"virtualServiceName"`
}

// Represents the properties needed to define a Virtual Service backend.
//
// TODO: EXAMPLE
//
type VirtualServiceBackendOptions struct {
	// TLS properties for  Client policy for the backend.
	TlsClientPolicy *TlsClientPolicy `json:"tlsClientPolicy" yaml:"tlsClientPolicy"`
}

// The properties applied to the VirtualService being defined.
//
// TODO: EXAMPLE
//
type VirtualServiceProps struct {
	// The VirtualNode or VirtualRouter which the VirtualService uses as its provider.
	VirtualServiceProvider VirtualServiceProvider `json:"virtualServiceProvider" yaml:"virtualServiceProvider"`
	// The name of the VirtualService.
	//
	// It is recommended this follows the fully-qualified domain name format,
	// such as "my-service.default.svc.cluster.local".
	//
	// Example value: `service.domain.local`
	VirtualServiceName *string `json:"virtualServiceName" yaml:"virtualServiceName"`
}

// Represents the properties needed to define the provider for a VirtualService.
//
// TODO: EXAMPLE
//
type VirtualServiceProvider interface {
	Bind(_construct constructs.Construct) *VirtualServiceProviderConfig
}

// The jsii proxy struct for VirtualServiceProvider
type jsiiProxy_VirtualServiceProvider struct {
	_ byte // padding
}

func NewVirtualServiceProvider_Override(v VirtualServiceProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		nil, // no parameters
		v,
	)
}

// Returns an Empty Provider for a VirtualService.
//
// This provides no routing capabilities
// and should only be used as a placeholder
func VirtualServiceProvider_None(mesh IMesh) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		"none",
		[]interface{}{mesh},
		&returns,
	)

	return returns
}

// Returns a VirtualNode based Provider for a VirtualService.
func VirtualServiceProvider_VirtualNode(virtualNode IVirtualNode) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		"virtualNode",
		[]interface{}{virtualNode},
		&returns,
	)

	return returns
}

// Returns a VirtualRouter based Provider for a VirtualService.
func VirtualServiceProvider_VirtualRouter(virtualRouter IVirtualRouter) VirtualServiceProvider {
	_init_.Initialize()

	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		"virtualRouter",
		[]interface{}{virtualRouter},
		&returns,
	)

	return returns
}

// Enforces mutual exclusivity for VirtualService provider types.
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
//
// TODO: EXAMPLE
//
type VirtualServiceProviderConfig struct {
	// Mesh the Provider is using.
	Mesh IMesh `json:"mesh" yaml:"mesh"`
	// Virtual Node based provider.
	VirtualNodeProvider *CfnVirtualService_VirtualNodeServiceProviderProperty `json:"virtualNodeProvider" yaml:"virtualNodeProvider"`
	// Virtual Router based provider.
	VirtualRouterProvider *CfnVirtualService_VirtualRouterServiceProviderProperty `json:"virtualRouterProvider" yaml:"virtualRouterProvider"`
}

// Properties for the Weighted Targets in the route.
//
// TODO: EXAMPLE
//
type WeightedTarget struct {
	// The VirtualNode the route points to.
	VirtualNode IVirtualNode `json:"virtualNode" yaml:"virtualNode"`
	// The weight for the target.
	Weight *float64 `json:"weight" yaml:"weight"`
}

