package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// An EKS cluster.
// Experimental.
type ICluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	// Experimental.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct.
	// Experimental.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesManifest` object.
	// Experimental.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Creates a new service account with corresponding IAM Role (IRSA).
	// Experimental.
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
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
	// Experimental.
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	// The unique ARN assigned to the service by AWS in the form of arn:aws:eks:.
	// Experimental.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn() *string
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint() *string
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	// The physical name of the Cluster.
	// Experimental.
	ClusterName() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId() *string
	// Custom environment variables when running `kubectl` against this cluster.
	// Experimental.
	KubectlEnvironment() *map[string]*string
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands to the cluster.
	// Experimental.
	KubectlLambdaRole() awsiam.IRole
	// An AWS Lambda layer that includes `kubectl`, `helm` and the `aws` CLI.
	//
	// If not defined, a default layer will be used.
	// Experimental.
	KubectlLayer() awslambda.ILayerVersion
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory() awscdk.Size
	// Subnets to host the `kubectl` compute resources.
	//
	// If this is undefined, the k8s endpoint is expected to be accessible
	// publicly.
	// Experimental.
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	// Kubectl Provider for issuing kubectl commands against it.
	//
	// If not defined, a default provider will be used.
	// Experimental.
	KubectlProvider() IKubectlProvider
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	// Experimental.
	KubectlRole() awsiam.IRole
	// A security group to use for `kubectl` execution.
	//
	// If this is undefined, the k8s endpoint is expected to be accessible
	// publicly.
	// Experimental.
	KubectlSecurityGroup() awsec2.ISecurityGroup
	// An AWS Lambda layer that includes the NPM dependency `proxy-agent`.
	//
	// If not defined, a default layer will be used.
	// Experimental.
	OnEventLayer() awslambda.ILayerVersion
	// The Open ID Connect Provider of the cluster used to configure Service Accounts.
	// Experimental.
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Indicates whether Kubernetes resources can be automatically pruned.
	//
	// When
	// this is enabled (default), prune labels will be allocated and injected to
	// each resource. These labels will then be used when issuing the `kubectl
	// apply` operation with the `--prune` switch.
	// Experimental.
	Prune() *bool
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	var returns KubernetesManifest

	_jsii_.Invoke(
		i,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	var returns HelmChart

	_jsii_.Invoke(
		i,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		i,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	var returns ServiceAccount

	_jsii_.Invoke(
		i,
		"addServiceAccount",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		i,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterHandlerSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterHandlerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlEnvironment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kubectlEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlLambdaRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlLambdaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"kubectlLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlMemory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"kubectlMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlPrivateSubnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"kubectlPrivateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlProvider() IKubectlProvider {
	var returns IKubectlProvider
	_jsii_.Get(
		j,
		"kubectlProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"kubectlSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) OnEventLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"onEventLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

