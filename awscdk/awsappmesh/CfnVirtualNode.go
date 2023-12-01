package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualNode := awscdk.Aws_appmesh.NewCfnVirtualNode(this, jsii.String("MyCfnVirtualNode"), &CfnVirtualNodeProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &VirtualNodeSpecProperty{
//   		BackendDefaults: &BackendDefaultsProperty{
//   			ClientPolicy: &ClientPolicyProperty{
//   				Tls: &ClientPolicyTlsProperty{
//   					Validation: &TlsValidationContextProperty{
//   						Trust: &TlsValidationContextTrustProperty{
//   							Acm: &TlsValidationContextAcmTrustProperty{
//   								CertificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Certificate: &ClientTlsCertificateProperty{
//   						File: &ListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &ListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Enforce: jsii.Boolean(false),
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Backends: []interface{}{
//   			&BackendProperty{
//   				VirtualService: &VirtualServiceBackendProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//
//   					// the properties below are optional
//   					ClientPolicy: &ClientPolicyProperty{
//   						Tls: &ClientPolicyTlsProperty{
//   							Validation: &TlsValidationContextProperty{
//   								Trust: &TlsValidationContextTrustProperty{
//   									Acm: &TlsValidationContextAcmTrustProperty{
//   										CertificateAuthorityArns: []*string{
//   											jsii.String("certificateAuthorityArns"),
//   										},
//   									},
//   									File: &TlsValidationContextFileTrustProperty{
//   										CertificateChain: jsii.String("certificateChain"),
//   									},
//   									Sds: &TlsValidationContextSdsTrustProperty{
//   										SecretName: jsii.String("secretName"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   									Match: &SubjectAlternativeNameMatchersProperty{
//   										Exact: []*string{
//   											jsii.String("exact"),
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Certificate: &ClientTlsCertificateProperty{
//   								File: &ListenerTlsFileCertificateProperty{
//   									CertificateChain: jsii.String("certificateChain"),
//   									PrivateKey: jsii.String("privateKey"),
//   								},
//   								Sds: &ListenerTlsSdsCertificateProperty{
//   									SecretName: jsii.String("secretName"),
//   								},
//   							},
//   							Enforce: jsii.Boolean(false),
//   							Ports: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Listeners: []interface{}{
//   			&ListenerProperty{
//   				PortMapping: &PortMappingProperty{
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   					Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Http: &VirtualNodeHttpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						MaxPendingRequests: jsii.Number(123),
//   					},
//   					Http2: &VirtualNodeHttp2ConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Tcp: &VirtualNodeTcpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//   					},
//   				},
//   				HealthCheck: &HealthCheckProperty{
//   					HealthyThreshold: jsii.Number(123),
//   					IntervalMillis: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					TimeoutMillis: jsii.Number(123),
//   					UnhealthyThreshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					Path: jsii.String("path"),
//   					Port: jsii.Number(123),
//   				},
//   				OutlierDetection: &OutlierDetectionProperty{
//   					BaseEjectionDuration: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					Interval: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					MaxEjectionPercent: jsii.Number(123),
//   					MaxServerErrors: jsii.Number(123),
//   				},
//   				Timeout: &ListenerTimeoutProperty{
//   					Grpc: &GrpcTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Http: &HttpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Http2: &HttpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Tcp: &TcpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Tls: &ListenerTlsProperty{
//   					Certificate: &ListenerTlsCertificateProperty{
//   						Acm: &ListenerTlsAcmCertificateProperty{
//   							CertificateArn: jsii.String("certificateArn"),
//   						},
//   						File: &ListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &ListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					Validation: &ListenerTlsValidationContextProperty{
//   						Trust: &ListenerTlsValidationContextTrustProperty{
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Logging: &LoggingProperty{
//   			AccessLog: &AccessLogProperty{
//   				File: &FileAccessLogProperty{
//   					Path: jsii.String("path"),
//
//   					// the properties below are optional
//   					Format: &LoggingFormatProperty{
//   						Json: []interface{}{
//   							&JsonFormatRefProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Text: jsii.String("text"),
//   					},
//   				},
//   			},
//   		},
//   		ServiceDiscovery: &ServiceDiscoveryProperty{
//   			AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   				NamespaceName: jsii.String("namespaceName"),
//   				ServiceName: jsii.String("serviceName"),
//
//   				// the properties below are optional
//   				Attributes: []interface{}{
//   					&AwsCloudMapInstanceAttributeProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				IpPreference: jsii.String("ipPreference"),
//   			},
//   			Dns: &DnsServiceDiscoveryProperty{
//   				Hostname: jsii.String("hostname"),
//
//   				// the properties below are optional
//   				IpPreference: jsii.String("ipPreference"),
//   				ResponseType: jsii.String("responseType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	MeshOwner: jsii.String("meshOwner"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualNodeName: jsii.String("virtualNodeName"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html
//
type CfnVirtualNode interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The full Amazon Resource Name (ARN) for the virtual node.
	AttrArn() *string
	AttrId() *string
	// The name of the service mesh that the virtual node resides in.
	AttrMeshName() *string
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrMeshOwner() *string
	// The AWS IAM account ID of the resource owner.
	//
	// If the account ID is not your own, then it's the ID of the mesh owner or of another account that the mesh is shared with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrResourceOwner() *string
	// The unique identifier for the virtual node.
	AttrUid() *string
	// The name of the virtual node.
	AttrVirtualNodeName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of the service mesh to create the virtual node in.
	MeshName() *string
	SetMeshName(val *string)
	// The AWS IAM account ID of the service mesh owner.
	MeshOwner() *string
	SetMeshOwner(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The virtual node specification to apply.
	Spec() interface{}
	SetSpec(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Optional metadata that you can apply to the virtual node to assist with categorization and organization.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The name to use for the virtual node.
	VirtualNodeName() *string
	SetVirtualNodeName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVirtualNode
type jsiiProxy_CfnVirtualNode struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnVirtualNode) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (j *jsiiProxy_CfnVirtualNode) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnVirtualNode) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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


func NewCfnVirtualNode(scope constructs.Construct, id *string, props *CfnVirtualNodeProps) CfnVirtualNode {
	_init_.Initialize()

	if err := validateNewCfnVirtualNodeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualNode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnVirtualNode_Override(c CfnVirtualNode, scope constructs.Construct, id *string, props *CfnVirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualNode)SetMeshName(val *string) {
	if err := j.validateSetMeshNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode)SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode)SetSpec(val interface{}) {
	if err := j.validateSetSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode)SetVirtualNodeName(val *string) {
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

	if err := validateCfnVirtualNode_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnVirtualNode_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualNode_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isCfnResource",
		[]interface{}{x},
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
func CfnVirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualNode_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnVirtualNode) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVirtualNode) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnVirtualNode) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVirtualNode) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualNode) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVirtualNode) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnVirtualNode) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

