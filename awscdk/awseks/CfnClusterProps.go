package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html
//
type CfnClusterProps struct {
	// The VPC configuration that's used by the cluster control plane.
	//
	// Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-resourcesvpcconfig
	//
	ResourcesVpcConfig interface{} `field:"required" json:"resourcesVpcConfig" yaml:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// For more information, see [Amazon EKS Service IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the **Amazon EKS User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The access configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-accessconfig
	//
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// If you set this value to `False` when creating a cluster, the default networking add-ons will not be installed.
	//
	// The default networking addons include vpc-cni, coredns, and kube-proxy.
	//
	// Use this option when you plan to install third-party alternative add-ons or self-manage the default networking add-ons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-bootstrapselfmanagedaddons
	//
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// The encryption configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The Kubernetes network configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-kubernetesnetworkconfig
	//
	KubernetesNetworkConfig interface{} `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// The logging configuration for your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The unique name to give to your cluster.
	//
	// The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphanumeric character and can't be longer than 100 characters. The name must be unique within the AWS Region and AWS account that you're creating the cluster in. Note that underscores can't be used in AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object representing the configuration of your local Amazon EKS cluster on an AWS Outpost.
	//
	// This object isn't available for clusters on the AWS cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-outpostconfig
	//
	OutpostConfig interface{} `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// The metadata that you apply to the cluster to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Cluster tags don't propagate to any other resources associated with the cluster.
	//
	// > You must have the `eks:TagResource` and `eks:UntagResource` permissions for your [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the AWS CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// This value indicates if extended support is enabled or disabled for the cluster.
	//
	// [Learn more about EKS Extended Support in the EKS User Guide.](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-upgradepolicy
	//
	UpgradePolicy interface{} `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the default version available in Amazon EKS is used.
	//
	// > The default version might not be the latest version available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The configuration for zonal shift for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-zonalshiftconfig
	//
	ZonalShiftConfig interface{} `field:"optional" json:"zonalShiftConfig" yaml:"zonalShiftConfig"`
}

