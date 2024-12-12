package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EKS control plane.
//
// The Amazon EKS control plane consists of control plane instances that run the Kubernetes software, such as `etcd` and the API server. The control plane runs in an account managed by AWS , and the Kubernetes API is exposed by the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is single tenant and unique. It runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones and fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS also provisions elastic network interfaces in your VPC subnets to provide connectivity from the control plane instances to the nodes (for example, to support `kubectl exec` , `logs` , and `proxy` data flows).
//
// Amazon EKS nodes run in your AWS account and connect to your cluster's control plane over the Kubernetes API server endpoint and a certificate file that is created for your cluster.
//
// You can use the `endpointPublicAccess` and `endpointPrivateAccess` parameters to enable or disable public and private access to your cluster's Kubernetes API server endpoint. By default, public access is enabled, and private access is disabled. For more information, see [Amazon EKS Cluster Endpoint Access Control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
//
// You can use the `logging` parameter to enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs. By default, cluster control plane logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS Cluster Control Plane Logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the **Amazon EKS User Guide** .
//
// > CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported control plane logs. For more information, see [CloudWatch Pricing](https://docs.aws.amazon.com/cloudwatch/pricing/) .
//
// In most cases, it takes several minutes to create a cluster. After you create an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate with the API server and launch nodes into your cluster. For more information, see [Allowing users to access your cluster](https://docs.aws.amazon.com/eks/latest/userguide/cluster-auth.html) and [Launching Amazon EKS nodes](https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_eks.NewCfnCluster(this, jsii.String("MyCfnCluster"), &CfnClusterProps{
//   	ResourcesVpcConfig: &ResourcesVpcConfigProperty{
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		EndpointPrivateAccess: jsii.Boolean(false),
//   		EndpointPublicAccess: jsii.Boolean(false),
//   		PublicAccessCidrs: []*string{
//   			jsii.String("publicAccessCidrs"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AccessConfig: &AccessConfigProperty{
//   		AuthenticationMode: jsii.String("authenticationMode"),
//   		BootstrapClusterCreatorAdminPermissions: jsii.Boolean(false),
//   	},
//   	BootstrapSelfManagedAddons: jsii.Boolean(false),
//   	ComputeConfig: &ComputeConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		NodePools: []*string{
//   			jsii.String("nodePools"),
//   		},
//   		NodeRoleArn: jsii.String("nodeRoleArn"),
//   	},
//   	EncryptionConfig: []interface{}{
//   		&EncryptionConfigProperty{
//   			Provider: &ProviderProperty{
//   				KeyArn: jsii.String("keyArn"),
//   			},
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   	},
//   	KubernetesNetworkConfig: &KubernetesNetworkConfigProperty{
//   		ElasticLoadBalancing: &ElasticLoadBalancingProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IpFamily: jsii.String("ipFamily"),
//   		ServiceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   		ServiceIpv6Cidr: jsii.String("serviceIpv6Cidr"),
//   	},
//   	Logging: &LoggingProperty{
//   		ClusterLogging: &ClusterLoggingProperty{
//   			EnabledTypes: []interface{}{
//   				&LoggingTypeConfigProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OutpostConfig: &OutpostConfigProperty{
//   		ControlPlaneInstanceType: jsii.String("controlPlaneInstanceType"),
//   		OutpostArns: []*string{
//   			jsii.String("outpostArns"),
//   		},
//
//   		// the properties below are optional
//   		ControlPlanePlacement: &ControlPlanePlacementProperty{
//   			GroupName: jsii.String("groupName"),
//   		},
//   	},
//   	RemoteNetworkConfig: &RemoteNetworkConfigProperty{
//   		RemoteNodeNetworks: []interface{}{
//   			&RemoteNodeNetworkProperty{
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RemotePodNetworks: []interface{}{
//   			&RemotePodNetworkProperty{
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   			},
//   		},
//   	},
//   	StorageConfig: &StorageConfigProperty{
//   		BlockStorage: &BlockStorageProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UpgradePolicy: &UpgradePolicyProperty{
//   		SupportType: jsii.String("supportType"),
//   	},
//   	Version: jsii.String("version"),
//   	ZonalShiftConfig: &ZonalShiftConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The access configuration for the cluster.
	AccessConfig() interface{}
	SetAccessConfig(val interface{})
	// The ARN of the cluster, such as `arn:aws:eks:us-west-2:666666666666:cluster/prod` .
	AttrArn() *string
	// The `certificate-authority-data` for your cluster.
	AttrCertificateAuthorityData() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	//
	// Managed node groups use this security group for control plane to data plane communication.
	//
	// This parameter is only returned by Amazon EKS clusters that support managed node groups. For more information, see [Managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in the *Amazon EKS User Guide* .
	AttrClusterSecurityGroupId() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	AttrEncryptionConfigKeyArn() *string
	// The endpoint for your Kubernetes API server, such as `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com` .
	AttrEndpoint() *string
	// The ID of your local Amazon EKS cluster on an AWS Outpost.
	//
	// This property isn't available for an Amazon EKS cluster on the AWS cloud.
	AttrId() *string
	// The CIDR block that Kubernetes Service IP addresses are assigned from if you created a `1.21` or later cluster with version `>1.10.1` or later of the Amazon VPC CNI add-on and specified `ipv6` for *ipFamily* when you created the cluster. Kubernetes assigns Service addresses from the unique local address range ( `fc00::/7` ) because you can't specify a custom `IPv6` CIDR block when you create the cluster.
	AttrKubernetesNetworkConfigServiceIpv6Cidr() *string
	// The issuer URL for the OIDC identity provider of the cluster, such as `https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E` . If you need to remove `https://` from this output value, you can include the following code in your template.
	//
	// `!Select [1, !Split ["//", !GetAtt EKSCluster.OpenIdConnectIssuerUrl]]`
	AttrOpenIdConnectIssuerUrl() *string
	// If you set this value to `False` when creating a cluster, the default networking add-ons will not be installed.
	BootstrapSelfManagedAddons() interface{}
	SetBootstrapSelfManagedAddons(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Indicates the current configuration of the compute capability on your EKS Auto Mode cluster.
	ComputeConfig() interface{}
	SetComputeConfig(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The encryption configuration for the cluster.
	EncryptionConfig() interface{}
	SetEncryptionConfig(val interface{})
	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig() interface{}
	SetKubernetesNetworkConfig(val interface{})
	// The logging configuration for your cluster.
	Logging() interface{}
	SetLogging(val interface{})
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
	// The unique name to give to your cluster.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// An object representing the configuration of your local Amazon EKS cluster on an AWS Outpost.
	OutpostConfig() interface{}
	SetOutpostConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The configuration in the cluster for EKS Hybrid Nodes.
	RemoteNetworkConfig() interface{}
	SetRemoteNetworkConfig(val interface{})
	// The VPC configuration that's used by the cluster control plane.
	ResourcesVpcConfig() interface{}
	SetResourcesVpcConfig(val interface{})
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Indicates the current configuration of the block storage capability on your EKS Auto Mode cluster.
	StorageConfig() interface{}
	SetStorageConfig(val interface{})
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The metadata that you apply to the cluster to assist with categorization and organization.
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
	// This value indicates if extended support is enabled or disabled for the cluster.
	UpgradePolicy() interface{}
	SetUpgradePolicy(val interface{})
	// The desired Kubernetes version for your cluster.
	Version() *string
	SetVersion(val *string)
	// The configuration for zonal shift for the cluster.
	ZonalShiftConfig() interface{}
	SetZonalShiftConfig(val interface{})
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCluster) AccessConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrKubernetesNetworkConfigServiceIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKubernetesNetworkConfigServiceIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ComputeConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EncryptionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KubernetesNetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OutpostConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) RemoteNetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ResourcesVpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) StorageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpgradePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ZonalShiftConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}


func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	if err := validateNewCfnClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster)SetAccessConfig(val interface{}) {
	if err := j.validateSetAccessConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetComputeConfig(val interface{}) {
	if err := j.validateSetComputeConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEncryptionConfig(val interface{}) {
	if err := j.validateSetEncryptionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetKubernetesNetworkConfig(val interface{}) {
	if err := j.validateSetKubernetesNetworkConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesNetworkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetLogging(val interface{}) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetOutpostConfig(val interface{}) {
	if err := j.validateSetOutpostConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetRemoteNetworkConfig(val interface{}) {
	if err := j.validateSetRemoteNetworkConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteNetworkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetResourcesVpcConfig(val interface{}) {
	if err := j.validateSetResourcesVpcConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesVpcConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetStorageConfig(val interface{}) {
	if err := j.validateSetStorageConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetUpgradePolicy(val interface{}) {
	if err := j.validateSetUpgradePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetZonalShiftConfig(val interface{}) {
	if err := j.validateSetZonalShiftConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalShiftConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCluster_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.CfnCluster",
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCluster) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

