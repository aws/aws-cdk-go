package awscdkeksv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Options for configuring an EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalHelmChartValues interface{}
//   var albControllerVersion AlbControllerVersion
//   var endpointAccess EndpointAccess
//   var keyRef IKeyRef
//   var kubernetesVersion KubernetesVersion
//   var layerVersion LayerVersion
//   var policy interface{}
//   var role Role
//   var securityGroup SecurityGroup
//   var size Size
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
//
//   clusterCommonOptions := &ClusterCommonOptions{
//   	Version: kubernetesVersion,
//
//   	// the properties below are optional
//   	AlbController: &AlbControllerOptions{
//   		Version: albControllerVersion,
//
//   		// the properties below are optional
//   		AdditionalHelmChartValues: map[string]interface{}{
//   			"additionalHelmChartValuesKey": additionalHelmChartValues,
//   		},
//   		OverwriteServiceAccount: jsii.Boolean(false),
//   		Policy: policy,
//   		Repository: jsii.String("repository"),
//   	},
//   	ClusterLogging: []ClusterLoggingTypes{
//   		eks_v2_alpha.ClusterLoggingTypes_API,
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	CoreDnsComputeType: eks_v2_alpha.CoreDnsComputeType_EC2,
//   	EndpointAccess: endpointAccess,
//   	IpFamily: eks_v2_alpha.IpFamily_IP_V4,
//   	KubectlProviderOptions: &KubectlProviderOptions{
//   		KubectlLayer: layerVersion,
//
//   		// the properties below are optional
//   		AwscliLayer: layerVersion,
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Memory: size,
//   		PrivateSubnets: []ISubnet{
//   			subnet,
//   		},
//   		Role: role,
//   		SecurityGroup: securityGroup,
//   	},
//   	MastersRole: role,
//   	Prune: jsii.Boolean(false),
//   	Role: role,
//   	SecretsEncryptionKey: keyRef,
//   	SecurityGroup: securityGroup,
//   	ServiceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Vpc: vpc,
//   	VpcSubnets: []SubnetSelection{
//   		&SubnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []SubnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []ISubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
// Experimental.
type ClusterCommonOptions struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Default: - The controller is not installed.
	//
	// Experimental.
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// The cluster log types which you want to enable.
	// Default: - none.
	//
	// Experimental.
	ClusterLogging *[]ClusterLoggingTypes `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
	// Name for the cluster.
	// Default: - Automatically generated name.
	//
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Default: CoreDnsComputeType.EC2 (for `FargateCluster` the default is FARGATE)
	//
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Default: EndpointAccess.PUBLIC_AND_PRIVATE
	//
	// Experimental.
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-ipFamily
	//
	// Default: IpFamily.IP_V4
	//
	// Experimental.
	IpFamily IpFamily `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Options for creating the kubectl provider - a lambda function that executes `kubectl` and `helm` against the cluster.
	//
	// If defined, `kubectlLayer` is a required property.
	//
	// If not defined, kubectl provider will not be created by default.
	// Experimental.
	KubectlProviderOptions *KubectlProviderOptions `field:"optional" json:"kubectlProviderOptions" yaml:"kubectlProviderOptions"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Default: - no masters role.
	//
	// Experimental.
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Default: true.
	//
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Default: - A role is automatically created for you.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Default: - By default, Kubernetes stores all secret object data within etcd and
	//   all etcd volumes used by Amazon EKS are encrypted at the disk-level
	//   using AWS-Managed encryption keys.
	//
	// Experimental.
	SecretsEncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// Security Group to use for Control Plane ENIs.
	// Default: - A security group is automatically created.
	//
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Default: - Kubernetes assigns addresses from either the
	//   10.100.0.0/16 or 172.20.0.0/16 CIDR blocks
	//
	// Experimental.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// The tags assigned to the EKS cluster.
	// Default: - none.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The VPC in which to create the Cluster.
	// Default: - a VPC with default configuration will be created and can be accessed through `cluster.vpc`.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE_WITH_EGRESS }]`
	// Default: - All public and private subnets.
	//
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

