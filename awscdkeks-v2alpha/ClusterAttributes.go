package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Attributes for EKS clusters.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv31"
//
//
//   handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   // get the serivceToken from the custom resource provider
//   functionArn := lambda.Function_FromFunctionName(this, jsii.String("ProviderOnEventFunc"), jsii.String("ProviderframeworkonEvent-XXX")).FunctionArn
//   kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
//   	ServiceToken: functionArn,
//   	Role: handlerRole,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("cluster"),
//   	KubectlProvider: KubectlProvider,
//   })
//
// Experimental.
type ClusterAttributes struct {
	// The physical name of the Cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The certificate-authority-data for your cluster.
	// Default: - if not specified `cluster.clusterCertificateAuthorityData` will
	// throw an error.
	//
	// Experimental.
	ClusterCertificateAuthorityData *string `field:"optional" json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Default: - if not specified `cluster.clusterEncryptionConfigKeyArn` will
	// throw an error.
	//
	// Experimental.
	ClusterEncryptionConfigKeyArn *string `field:"optional" json:"clusterEncryptionConfigKeyArn" yaml:"clusterEncryptionConfigKeyArn"`
	// The API Server endpoint URL.
	// Default: - if not specified `cluster.clusterEndpoint` will throw an error.
	//
	// Experimental.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Default: - if not specified `cluster.clusterSecurityGroupId` will throw an
	// error.
	//
	// Experimental.
	ClusterSecurityGroupId *string `field:"optional" json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: - IpFamily.IP_V4
	//
	// Experimental.
	IpFamily IpFamily `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// KubectlProvider for issuing kubectl commands.
	// Default: - Default CDK provider.
	//
	// Experimental.
	KubectlProvider IKubectlProvider `field:"optional" json:"kubectlProvider" yaml:"kubectlProvider"`
	// Options for creating the kubectl provider - a lambda function that executes `kubectl` and `helm` against the cluster.
	//
	// If defined, `kubectlLayer` is a required property.
	//
	// If not defined, kubectl provider will not be created by default.
	// Experimental.
	KubectlProviderOptions *KubectlProviderOptions `field:"optional" json:"kubectlProviderOptions" yaml:"kubectlProviderOptions"`
	// An Open ID Connect provider for this cluster that can be used to configure service accounts.
	//
	// You can either import an existing provider using `iam.OpenIdConnectProvider.fromProviderArn`,
	// or create a new provider using `new eks.OpenIdConnectProvider`
	// Default: - if not specified `cluster.openIdConnectProvider` and `cluster.addServiceAccount` will throw an error.
	//
	// Experimental.
	OpenIdConnectProvider awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProvider" yaml:"openIdConnectProvider"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Default: true.
	//
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Additional security groups associated with this cluster.
	// Default: - if not specified, no additional security groups will be
	// considered in `cluster.connections`.
	//
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC in which this Cluster was created.
	// Default: - if not specified `cluster.vpc` will throw an error
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

