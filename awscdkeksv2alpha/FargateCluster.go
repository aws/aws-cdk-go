package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an EKS cluster that runs entirely on AWS Fargate.
//
// The cluster is created with a default Fargate Profile that matches the
// "default" and "kube-system" namespaces. You can add additional profiles using
// `addFargateProfile`.
//
// Example:
//   cluster := eks.NewFargateCluster(this, jsii.String("FargateCluster"), &FargateClusterProps{
//   	Version: eks.KubernetesVersion_V1_34(),
//   })
//
// Experimental.
type FargateCluster interface {
	Cluster
	// The ALB Controller construct defined for this cluster.
	//
	// Will be undefined if `albController` wasn't configured.
	// Experimental.
	AlbController() AlbController
	// The AWS generated ARN for the Cluster resource.
	//
	// For example, `arn:aws:eks:us-west-2:666666666666:cluster/prod`.
	// Experimental.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn() *string
	// The endpoint URL for the Cluster.
	//
	// This is the URL inside the kubeconfig file to use with kubectl
	//
	// For example, `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com`
	// Experimental.
	ClusterEndpoint() *string
	// The Name of the created EKS Cluster.
	// Experimental.
	ClusterName() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer url.
	//
	// If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	// Experimental.
	ClusterOpenIdConnectIssuerUrl() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId() *string
	// Manages connection rules (Security Group Rules) for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The auto scaling group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is not `EC2` or
	// `defaultCapacityType` is `EC2` but default capacity is set to 0.
	// Experimental.
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	// The node group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is `EC2` or
	// `defaultCapacityType` is `NODEGROUP` but default capacity is set to 0.
	// Experimental.
	DefaultNodegroup() Nodegroup
	// Fargate Profile that was created with the cluster.
	// Experimental.
	DefaultProfile() FargateProfile
	// Retrieves the EKS Pod Identity Agent addon for the EKS cluster.
	//
	// The EKS Pod Identity Agent is responsible for managing the temporary credentials
	// used by pods in the cluster to access AWS resources. It runs as a DaemonSet on
	// each node and provides the necessary credentials to the pods based on their
	// associated service account.
	// Experimental.
	EksPodIdentityAgent() IAddon
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	// Experimental.
	Env() *interfaces.ResourceEnvironment
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: - IpFamily.IP_V4
	//
	// Experimental.
	IpFamily() IpFamily
	// Experimental.
	KubectlProvider() IKubectlProvider
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// An `OpenIdConnectProvider` resource associated with this cluster, and which can be used to link this cluster to AWS IAM.
	//
	// A provider will only be defined if this property is accessed (lazy initialization).
	// Experimental.
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Determines if Kubernetes resources can be pruned automatically.
	// Experimental.
	Prune() *bool
	// IAM role assumed by the EKS Control Plane.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc() awsec2.IVpc
	// Add nodes to this EKS cluster.
	//
	// The nodes will automatically be configured with the right VPC and AMI
	// for the instance type and Kubernetes version.
	//
	// Note that if you specify `updateType: RollingUpdate` or `updateType: ReplacingUpdate`, your nodes might be replaced at deploy
	// time without notice in case the recommended AMI for your machine image type has been updated by AWS.
	// The default behavior for `updateType` is `None`, which means only new instances will be launched using the new AMI.
	// Experimental.
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	// Experimental.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Adds a Fargate profile to this cluster.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
	//
	// Experimental.
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct.
	// Experimental.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesResource` object.
	// Experimental.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Add managed nodegroup to this Amazon EKS cluster.
	//
	// This method will create a new managed nodegroup and add into the capacity.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
	//
	// Experimental.
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	// Creates a new service account with corresponding IAM Role (IRSA).
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Connect capacity in the form of an existing AutoScalingGroup to the EKS cluster.
	//
	// The AutoScalingGroup must be running an EKS-optimized AMI containing the
	// /etc/eks/bootstrap.sh script. This method will configure Security Groups,
	// add the right policies to the instance role, apply the right tags, and add
	// the required user data to the instance's launch configuration.
	//
	// Prefer to use `addAutoScalingGroupCapacity` if possible.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
	//
	// Experimental.
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	// Experimental.
	GeneratePhysicalName() *string
	// Fetch the load balancer address of an ingress backed by a load balancer.
	// Experimental.
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
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
	// Fetch the load balancer address of a service of type 'LoadBalancer'.
	// Experimental.
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
	// Grants the specified IAM principal access to the EKS cluster based on the provided access policies.
	//
	// This method creates an `AccessEntry` construct that grants the specified IAM principal the access permissions
	// defined by the provided `IAccessPolicy` array. This allows the IAM principal to perform the actions permitted
	// by the access policies within the EKS cluster.
	// Experimental.
	GrantAccess(id *string, principal *string, accessPolicies *[]IAccessPolicy)
	// Grants the specified IAM principal cluster admin access to the EKS cluster.
	//
	// This method creates an `AccessEntry` construct that grants the specified IAM principal the cluster admin
	// access permissions. This allows the IAM principal to perform the actions permitted
	// by the cluster admin acces.
	//
	// Returns: the access entry construct.
	// Experimental.
	GrantClusterAdmin(id *string, principal *string) AccessEntry
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FargateCluster
type jsiiProxy_FargateCluster struct {
	jsiiProxy_Cluster
}

func (j *jsiiProxy_FargateCluster) AlbController() AlbController {
	var returns AlbController
	_jsii_.Get(
		j,
		"albController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultCapacity() awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup
	_jsii_.Get(
		j,
		"defaultCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultNodegroup() Nodegroup {
	var returns Nodegroup
	_jsii_.Get(
		j,
		"defaultNodegroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultProfile() FargateProfile {
	var returns FargateProfile
	_jsii_.Get(
		j,
		"defaultProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) EksPodIdentityAgent() IAddon {
	var returns IAddon
	_jsii_.Get(
		j,
		"eksPodIdentityAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) IpFamily() IpFamily {
	var returns IpFamily
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlProvider() IKubectlProvider {
	var returns IKubectlProvider
	_jsii_.Get(
		j,
		"kubectlProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateCluster(scope constructs.Construct, id *string, props *FargateClusterProps) FargateCluster {
	_init_.Initialize()

	if err := validateNewFargateClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FargateCluster{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateCluster_Override(f FargateCluster, scope constructs.Construct, id *string, props *FargateClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing cluster.
// Experimental.
func FargateCluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	if err := validateFargateCluster_FromClusterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ICluster

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
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
// Experimental.
func FargateCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFargateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FargateCluster_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFargateCluster_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FargateCluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFargateCluster_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func FargateCluster_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FargateCluster) AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup {
	if err := f.validateAddAutoScalingGroupCapacityParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		f,
		"addAutoScalingGroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	if err := f.validateAddCdk8sChartParameters(id, chart, options); err != nil {
		panic(err)
	}
	var returns KubernetesManifest

	_jsii_.Invoke(
		f,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile {
	if err := f.validateAddFargateProfileParameters(id, options); err != nil {
		panic(err)
	}
	var returns FargateProfile

	_jsii_.Invoke(
		f,
		"addFargateProfile",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	if err := f.validateAddHelmChartParameters(id, options); err != nil {
		panic(err)
	}
	var returns HelmChart

	_jsii_.Invoke(
		f,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	if err := f.validateAddManifestParameters(id); err != nil {
		panic(err)
	}
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		f,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup {
	if err := f.validateAddNodegroupCapacityParameters(id, options); err != nil {
		panic(err)
	}
	var returns Nodegroup

	_jsii_.Invoke(
		f,
		"addNodegroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	if err := f.validateAddServiceAccountParameters(id, options); err != nil {
		panic(err)
	}
	var returns ServiceAccount

	_jsii_.Invoke(
		f,
		"addServiceAccount",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FargateCluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	if err := f.validateConnectAutoScalingGroupCapacityParameters(autoScalingGroup, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

func (f *jsiiProxy_FargateCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string {
	if err := f.validateGetIngressLoadBalancerAddressParameters(ingressName, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getIngressLoadBalancerAddress",
		[]interface{}{ingressName, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string {
	if err := f.validateGetServiceLoadBalancerAddressParameters(serviceName, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getServiceLoadBalancerAddress",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) GrantAccess(id *string, principal *string, accessPolicies *[]IAccessPolicy) {
	if err := f.validateGrantAccessParameters(id, principal, accessPolicies); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"grantAccess",
		[]interface{}{id, principal, accessPolicies},
	)
}

func (f *jsiiProxy_FargateCluster) GrantClusterAdmin(id *string, principal *string) AccessEntry {
	if err := f.validateGrantClusterAdminParameters(id, principal); err != nil {
		panic(err)
	}
	var returns AccessEntry

	_jsii_.Invoke(
		f,
		"grantClusterAdmin",
		[]interface{}{id, principal},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

