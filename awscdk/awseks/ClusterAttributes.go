package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Attributes for EKS clusters.
//
// Example:
//   var cluster cluster
//   var asg autoScalingGroup
//
//   importedCluster := eks.cluster_FromClusterAttributes(this, jsii.String("ImportedCluster"), &ClusterAttributes{
//   	ClusterName: cluster.ClusterName,
//   	ClusterSecurityGroupId: cluster.ClusterSecurityGroupId,
//   })
//
//   importedCluster.ConnectAutoScalingGroupCapacity(asg, &AutoScalingGroupOptions{
//   })
//
type ClusterAttributes struct {
	// The physical name of the Cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
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
	// The certificate-authority-data for your cluster.
	// Default: - if not specified `cluster.clusterCertificateAuthorityData` will
	// throw an error.
	//
	ClusterCertificateAuthorityData *string `field:"optional" json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Default: - if not specified `cluster.clusterEncryptionConfigKeyArn` will
	// throw an error.
	//
	ClusterEncryptionConfigKeyArn *string `field:"optional" json:"clusterEncryptionConfigKeyArn" yaml:"clusterEncryptionConfigKeyArn"`
	// The API Server endpoint URL.
	// Default: - if not specified `cluster.clusterEndpoint` will throw an error.
	//
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// A security group id to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	// Default: - No security group.
	//
	ClusterHandlerSecurityGroupId *string `field:"optional" json:"clusterHandlerSecurityGroupId" yaml:"clusterHandlerSecurityGroupId"`
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Default: - if not specified `cluster.clusterSecurityGroupId` will throw an
	// error.
	//
	ClusterSecurityGroupId *string `field:"optional" json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: - IpFamily.IP_V4
	//
	IpFamily IpFamily `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Environment variables to use when running `kubectl` against this cluster.
	// Default: - no additional variables.
	//
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands
	// to the cluster.
	// Default: - if not specified, the default role created by a lambda function will
	// be used.
	//
	KubectlLambdaRole awsiam.IRole `field:"optional" json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
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
	// Default: - No default layer will be provided.
	//
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Default: Size.gibibytes(1)
	//
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// Subnets to host the `kubectl` compute resources.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Default: - k8s endpoint is expected to be accessible publicly.
	//
	KubectlPrivateSubnetIds *[]*string `field:"optional" json:"kubectlPrivateSubnetIds" yaml:"kubectlPrivateSubnetIds"`
	// KubectlProvider for issuing kubectl commands.
	// Default: - Default CDK provider.
	//
	KubectlProvider IKubectlProvider `field:"optional" json:"kubectlProvider" yaml:"kubectlProvider"`
	// An IAM role with cluster administrator and "system:masters" permissions.
	// Default: - if not specified, it not be possible to issue `kubectl` commands
	// against an imported cluster.
	//
	KubectlRoleArn *string `field:"optional" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
	// A security group to use for `kubectl` execution.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Default: - k8s endpoint is expected to be accessible publicly.
	//
	KubectlSecurityGroupId *string `field:"optional" json:"kubectlSecurityGroupId" yaml:"kubectlSecurityGroupId"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// The handler expects the layer to include the following node_modules:
	//
	// proxy-agent.
	// Default: - a layer bundled with this module.
	//
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// An Open ID Connect provider for this cluster that can be used to configure service accounts.
	//
	// You can either import an existing provider using `iam.OpenIdConnectProvider.fromProviderArn`,
	// or create a new provider using `new eks.OpenIdConnectProvider`
	// Default: - if not specified `cluster.openIdConnectProvider` and `cluster.addServiceAccount` will throw an error.
	//
	OpenIdConnectProvider awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProvider" yaml:"openIdConnectProvider"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Default: true.
	//
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Additional security groups associated with this cluster.
	// Default: - if not specified, no additional security groups will be
	// considered in `cluster.connections`.
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC in which this Cluster was created.
	// Default: - if not specified `cluster.vpc` will throw an error
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

