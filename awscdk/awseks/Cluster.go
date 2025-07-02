package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Cluster represents a managed Kubernetes Service (EKS).
//
// This is a fully managed cluster of API Servers (control-plane)
// The user is still required to create the worker nodes.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
//
//   // or
//   var vpc vpc
//
//
//   eks.NewCluster(this, jsii.String("MyCluster"), &ClusterProps{
//   	KubectlMemory: awscdk.Size_Gibibytes(jsii.Number(4)),
//   	Version: eks.KubernetesVersion_V1_33(),
//   	KubectlLayer: kubectlv33.NewKubectlV33Layer(this, jsii.String("kubectl")),
//   })
//   eks.Cluster_FromClusterAttributes(this, jsii.String("MyCluster"), &ClusterAttributes{
//   	KubectlMemory: awscdk.Size_*Gibibytes(jsii.Number(4)),
//   	Vpc: Vpc,
//   	ClusterName: jsii.String("cluster-name"),
//   })
//
type Cluster interface {
	awscdk.Resource
	ICluster
	// An IAM role with administrative permissions to create or update the cluster.
	//
	// This role also has `systems:master` permissions.
	AdminRole() awsiam.Role
	// The ALB Controller construct defined for this cluster.
	//
	// Will be undefined if `albController` wasn't configured.
	AlbController() AlbController
	// The authentication mode for the Amazon EKS cluster.
	//
	// The authentication mode determines how users and applications authenticate to the Kubernetes API server.
	// Default: CONFIG_MAP.
	//
	AuthenticationMode() AuthenticationMode
	// Lazily creates the AwsAuth resource, which manages AWS authentication mapping.
	AwsAuth() AwsAuth
	// An AWS Lambda layer that contains the `aws` CLI.
	//
	// If not defined, a default layer will be used containing the AWS CLI 1.x.
	AwscliLayer() awslambda.ILayerVersion
	// The AWS generated ARN for the Cluster resource.
	//
	// For example, `arn:aws:eks:us-west-2:666666666666:cluster/prod`.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	ClusterEncryptionConfigKeyArn() *string
	// The endpoint URL for the Cluster.
	//
	// This is the URL inside the kubeconfig file to use with kubectl
	//
	// For example, `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com`
	ClusterEndpoint() *string
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Default: - No security group.
	//
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	// The Name of the created EKS Cluster.
	ClusterName() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	ClusterOpenIdConnectIssuer() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer url.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	ClusterOpenIdConnectIssuerUrl() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	ClusterSecurityGroupId() *string
	// Manages connection rules (Security Group Rules) for the cluster.
	Connections() awsec2.Connections
	// The auto scaling group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is not `EC2` or
	// `defaultCapacityType` is `EC2` but default capacity is set to 0.
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	// The node group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is `EC2` or
	// `defaultCapacityType` is `NODEGROUP` but default capacity is set to 0.
	DefaultNodegroup() Nodegroup
	// Retrieves the EKS Pod Identity Agent addon for the EKS cluster.
	//
	// The EKS Pod Identity Agent is responsible for managing the temporary credentials
	// used by pods in the cluster to access AWS resources. It runs as a DaemonSet on
	// each node and provides the necessary credentials to the pods based on their
	// associated service account.
	EksPodIdentityAgent() IAddon
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: - IpFamily.IP_V4
	//
	IpFamily() IpFamily
	// Custom environment variables when running `kubectl` against this cluster.
	KubectlEnvironment() *map[string]*string
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands to the cluster.
	// Default: - if not specified, the default role created by a lambda function will
	// be used.
	//
	KubectlLambdaRole() awsiam.IRole
	// An AWS Lambda layer that includes `kubectl` and `helm`.
	KubectlLayer() awslambda.ILayerVersion
	// The amount of memory allocated to the kubectl provider's lambda function.
	KubectlMemory() awscdk.Size
	// Subnets to host the `kubectl` compute resources.
	// Default: - If not specified, the k8s endpoint is expected to be accessible
	// publicly.
	//
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	KubectlRole() awsiam.IRole
	// A security group to use for `kubectl` execution.
	// Default: - If not specified, the k8s endpoint is expected to be accessible
	// publicly.
	//
	KubectlSecurityGroup() awsec2.ISecurityGroup
	// The tree node.
	Node() constructs.Node
	// The AWS Lambda layer that contains the NPM dependency `proxy-agent`.
	//
	// If
	// undefined, a SAR app that contains this layer will be used.
	OnEventLayer() awslambda.ILayerVersion
	// An `OpenIdConnectProvider` resource associated with this cluster, and which can be used to link this cluster to AWS IAM.
	//
	// A provider will only be defined if this property is accessed (lazy initialization).
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// Determines if Kubernetes resources can be pruned automatically.
	Prune() *bool
	// IAM role assumed by the EKS Control Plane.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC in which this Cluster was created.
	Vpc() awsec2.IVpc
	// Add nodes to this EKS cluster.
	//
	// The nodes will automatically be configured with the right VPC and AMI
	// for the instance type and Kubernetes version.
	//
	// Note that if you specify `updateType: RollingUpdate` or `updateType: ReplacingUpdate`, your nodes might be replaced at deploy
	// time without notice in case the recommended AMI for your machine image type has been updated by AWS.
	// The default behavior for `updateType` is `None`, which means only new instances will be launched using the new AMI.
	//
	// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
	// In addition, the [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
	// daemon will be installed on all spot instances to handle
	// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Adds a Fargate profile to this cluster.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
	//
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesResource` object.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Add managed nodegroup to this Amazon EKS cluster.
	//
	// This method will create a new managed nodegroup and add into the capacity.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
	//
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	// Creates a new service account with corresponding IAM Role (IRSA).
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
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
	// Connect capacity in the form of an existing AutoScalingGroup to the EKS cluster.
	//
	// The AutoScalingGroup must be running an EKS-optimized AMI containing the
	// /etc/eks/bootstrap.sh script. This method will configure Security Groups,
	// add the right policies to the instance role, apply the right tags, and add
	// the required user data to the instance's launch configuration.
	//
	// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
	// If kubectl is enabled, the
	// [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
	// daemon will be installed on all spot instances to handle
	// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
	//
	// Prefer to use `addAutoScalingGroupCapacity` if possible.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
	//
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	GeneratePhysicalName() *string
	// Fetch the load balancer address of an ingress backed by a load balancer.
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
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
	// Fetch the load balancer address of a service of type 'LoadBalancer'.
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
	// Grants the specified IAM principal access to the EKS cluster based on the provided access policies.
	//
	// This method creates an `AccessEntry` construct that grants the specified IAM principal the access permissions
	// defined by the provided `IAccessPolicy` array. This allows the IAM principal to perform the actions permitted
	// by the access policies within the EKS cluster.
	GrantAccess(id *string, principal *string, accessPolicies *[]IAccessPolicy)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) AdminRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"adminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AlbController() AlbController {
	var returns AlbController
	_jsii_.Get(
		j,
		"albController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AuthenticationMode() AuthenticationMode {
	var returns AuthenticationMode
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AwsAuth() AwsAuth {
	var returns AwsAuth
	_jsii_.Get(
		j,
		"awsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AwscliLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"awscliLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterHandlerSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterHandlerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterOpenIdConnectIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultCapacity() awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup
	_jsii_.Get(
		j,
		"defaultCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultNodegroup() Nodegroup {
	var returns Nodegroup
	_jsii_.Get(
		j,
		"defaultNodegroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EksPodIdentityAgent() IAddon {
	var returns IAddon
	_jsii_.Get(
		j,
		"eksPodIdentityAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IpFamily() IpFamily {
	var returns IpFamily
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlEnvironment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kubectlEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlLambdaRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlLambdaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"kubectlLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlMemory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"kubectlMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlPrivateSubnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"kubectlPrivateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"kubectlSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) OnEventLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"onEventLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Initiates an EKS Cluster with the supplied arguments.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Initiates an EKS Cluster with the supplied arguments.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing cluster.
func Cluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	if err := validateCluster_FromClusterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ICluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Cluster",
		"fromClusterAttributes",
		[]interface{}{scope, id, attrs},
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
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Cluster_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Cluster",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Cluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Cluster_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.Cluster",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cluster) AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup {
	if err := c.validateAddAutoScalingGroupCapacityParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		c,
		"addAutoScalingGroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	if err := c.validateAddCdk8sChartParameters(id, chart, options); err != nil {
		panic(err)
	}
	var returns KubernetesManifest

	_jsii_.Invoke(
		c,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile {
	if err := c.validateAddFargateProfileParameters(id, options); err != nil {
		panic(err)
	}
	var returns FargateProfile

	_jsii_.Invoke(
		c,
		"addFargateProfile",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	if err := c.validateAddHelmChartParameters(id, options); err != nil {
		panic(err)
	}
	var returns HelmChart

	_jsii_.Invoke(
		c,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	if err := c.validateAddManifestParameters(id); err != nil {
		panic(err)
	}
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		c,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup {
	if err := c.validateAddNodegroupCapacityParameters(id, options); err != nil {
		panic(err)
	}
	var returns Nodegroup

	_jsii_.Invoke(
		c,
		"addNodegroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	if err := c.validateAddServiceAccountParameters(id, options); err != nil {
		panic(err)
	}
	var returns ServiceAccount

	_jsii_.Invoke(
		c,
		"addServiceAccount",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Cluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	if err := c.validateConnectAutoScalingGroupCapacityParameters(autoScalingGroup, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string {
	if err := c.validateGetIngressLoadBalancerAddressParameters(ingressName, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getIngressLoadBalancerAddress",
		[]interface{}{ingressName, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string {
	if err := c.validateGetServiceLoadBalancerAddressParameters(serviceName, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getServiceLoadBalancerAddress",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GrantAccess(id *string, principal *string, accessPolicies *[]IAccessPolicy) {
	if err := c.validateGrantAccessParameters(id, principal, accessPolicies); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"grantAccess",
		[]interface{}{id, principal, accessPolicies},
	)
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

