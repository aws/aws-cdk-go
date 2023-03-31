package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AppMesh::Route`.
//
// Creates a route that is associated with a virtual router.
//
// You can route several different protocols and define a retry policy for a route. Traffic can be routed to one or more virtual nodes.
//
// For more information about routes, see [Routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoute := awscdk.Aws_appmesh.NewCfnRoute(this, jsii.String("MyCfnRoute"), &cfnRouteProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &routeSpecProperty{
//   		grpcRoute: &grpcRouteProperty{
//   			action: &grpcRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &grpcRouteMatchProperty{
//   				metadata: []interface{}{
//   					&grpcRouteMetadataProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &grpcRouteMetadataMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				methodName: jsii.String("methodName"),
//   				port: jsii.Number(123),
//   				serviceName: jsii.String("serviceName"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &grpcRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				grpcRetryEvents: []*string{
//   					jsii.String("grpcRetryEvents"),
//   				},
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &grpcTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		http2Route: &httpRouteProperty{
//   			action: &httpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &httpRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &headerMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				method: jsii.String("method"),
//   				path: &httpPathMatchProperty{
//   					exact: jsii.String("exact"),
//   					regex: jsii.String("regex"),
//   				},
//   				port: jsii.Number(123),
//   				prefix: jsii.String("prefix"),
//   				queryParameters: []interface{}{
//   					&queryParameterProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						match: &httpQueryParameterMatchProperty{
//   							exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &httpRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		httpRoute: &httpRouteProperty{
//   			action: &httpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &httpRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &headerMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				method: jsii.String("method"),
//   				path: &httpPathMatchProperty{
//   					exact: jsii.String("exact"),
//   					regex: jsii.String("regex"),
//   				},
//   				port: jsii.Number(123),
//   				prefix: jsii.String("prefix"),
//   				queryParameters: []interface{}{
//   					&queryParameterProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						match: &httpQueryParameterMatchProperty{
//   							exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &httpRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		priority: jsii.Number(123),
//   		tcpRoute: &tcpRouteProperty{
//   			action: &tcpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			match: &tcpRouteMatchProperty{
//   				port: jsii.Number(123),
//   			},
//   			timeout: &tcpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	virtualRouterName: jsii.String("virtualRouterName"),
//
//   	// the properties below are optional
//   	meshOwner: jsii.String("meshOwner"),
//   	routeName: jsii.String("routeName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRoute interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The full Amazon Resource Name (ARN) for the route.
	AttrArn() *string
	// The name of the service mesh that the route resides in.
	AttrMeshName() *string
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrMeshOwner() *string
	// The AWS IAM account ID of the resource owner.
	//
	// If the account ID is not your own, then it's the ID of the mesh owner or of another account that the mesh is shared with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrResourceOwner() *string
	// The name of the route.
	AttrRouteName() *string
	// The unique identifier for the route.
	AttrUid() *string
	// The name of the virtual router that the route is associated with.
	AttrVirtualRouterName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the service mesh to create the route in.
	MeshName() *string
	SetMeshName(val *string)
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner() *string
	SetMeshOwner(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The name to use for the route.
	RouteName() *string
	SetRouteName(val *string)
	// The route specification to apply.
	Spec() interface{}
	SetSpec(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Optional metadata that you can apply to the route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The name of the virtual router in which to create the route.
	//
	// If the virtual router is in a shared mesh, then you must be the owner of the virtual router resource.
	VirtualRouterName() *string
	SetVirtualRouterName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
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

	if err := validateNewCfnRouteParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnRoute)SetMeshName(val *string) {
	if err := j.validateSetMeshNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnRoute)SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnRoute)SetRouteName(val *string) {
	_jsii_.Set(
		j,
		"routeName",
		val,
	)
}

func (j *jsiiProxy_CfnRoute)SetSpec(val interface{}) {
	if err := j.validateSetSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnRoute)SetVirtualRouterName(val *string) {
	if err := j.validateSetVirtualRouterNameParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnRoute_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnRoute_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnRoute) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRoute) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRoute) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRoute) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRoute) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRoute) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRoute) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRoute) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRoute) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRoute) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRoute) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnRoute) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRoute) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRoute) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

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

func (c *jsiiProxy_CfnRoute) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnRoute) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

