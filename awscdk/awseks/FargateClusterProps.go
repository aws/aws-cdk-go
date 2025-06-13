package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Configuration props for EKS Fargate.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"
//
//
//   cluster := eks.NewFargateCluster(this, jsii.String("MyCluster"), &FargateClusterProps{
//   	Version: eks.KubernetesVersion_V1_32(),
//   	KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
//   })
//
type FargateClusterProps struct {
	// The Kubernetes version to run in the cluster.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	// Default: - Automatically generated name.
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Default: false.
	//
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and the ARN of the masters IAM role.
	//
	// Note: If mastersRole is not specified, this property will be ignored and no config command will be emitted.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_eks-readme.html#masters-role
	//
	// Default: true.
	//
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Default: - A role is automatically created for you.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Default: - A security group is automatically created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Default: - a VPC with default configuration will be created and can be accessed through `cluster.vpc`.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE_WITH_EGRESS }]`
	// Default: - All public and private subnets.
	//
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// An AWS Lambda Layer which includes `kubectl` and Helm.
	//
	// This layer is used by the kubectl handler to apply manifests and install
	// helm charts. You must pick an appropriate releases of one of the
	// `@aws-cdk/layer-kubectl-vXX` packages, that works with the version of
	// Kubernetes you have chosen.
	//
	// The handler expects the layer to include the following executables:
	//
	// ```
	// /opt/helm/helm
	// /opt/kubectl/kubectl
	// ```.
	KubectlLayer awslambda.ILayerVersion `field:"required" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Default: - The controller is not installed.
	//
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// The desired authentication mode for the cluster.
	// Default: AuthenticationMode.CONFIG_MAP
	//
	AuthenticationMode AuthenticationMode `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// An AWS Lambda layer that contains the `aws` CLI.
	//
	// The handler expects the layer to include the following executables:
	//
	// ```
	// /opt/awscli/aws
	// ```.
	// Default: - a default layer with the AWS CLI 1.x
	//
	AwscliLayer awslambda.ILayerVersion `field:"optional" json:"awscliLayer" yaml:"awscliLayer"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Default: - No environment variables.
	//
	ClusterHandlerEnvironment *map[string]*string `field:"optional" json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Default: - No security group.
	//
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// The cluster log types which you want to enable.
	// Default: - none.
	//
	ClusterLogging *[]ClusterLoggingTypes `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Default: CoreDnsComputeType.EC2 (for `FargateCluster` the default is FARGATE)
	//
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Default: EndpointAccess.PUBLIC_AND_PRIVATE
	//
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: - IpFamily.IP_V4
	//
	IpFamily IpFamily `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Default: - No environment variables.
	//
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// Amount of memory to allocate to the provider's lambda function.
	// Default: Size.gibibytes(1)
	//
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Default: - no masters role.
	//
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-node-proxy-agent" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally define it in your app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'proxy-agent-layer', {
	//   code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//   compatibleRuntimes: [lambda.Runtime.NODEJS_LATEST],
	// });
	// ```.
	// Default: - a layer bundled with this module.
	//
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Default: false.
	//
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Default: false.
	//
	PlaceClusterHandlerInVpc *bool `field:"optional" json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Default: true.
	//
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// IPv4 CIDR blocks defining the expected address range of hybrid nodes that will join the cluster.
	// Default: - none.
	//
	RemoteNodeNetworks *[]*RemoteNodeNetwork `field:"optional" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// IPv4 CIDR blocks for Pods running Kubernetes webhooks on hybrid nodes.
	// Default: - none.
	//
	RemotePodNetworks *[]*RemotePodNetwork `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Default: - By default, Kubernetes stores all secret object data within etcd and
	//   all etcd volumes used by Amazon EKS are encrypted at the disk-level
	//   using AWS-Managed encryption keys.
	//
	SecretsEncryptionKey awskms.IKey `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Default: - Kubernetes assigns addresses from either the
	//   10.100.0.0/16 or 172.20.0.0/16 CIDR blocks
	//
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// Fargate Profile to create along with the cluster.
	// Default: - A profile called "default" with 'default' and 'kube-system'
	//   selectors will be created if this is left undefined.
	//
	DefaultProfile *FargateProfileOptions `field:"optional" json:"defaultProfile" yaml:"defaultProfile"`
}

