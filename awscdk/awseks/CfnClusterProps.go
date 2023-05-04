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
//   	Version: jsii.String("version"),
//   }
//
type CfnClusterProps struct {
	// The VPC configuration that's used by the cluster control plane.
	//
	// Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
	//
	// > Updates require replacement of the `SecurityGroupIds` and `SubnetIds` sub-properties.
	ResourcesVpcConfig interface{} `field:"required" json:"resourcesVpcConfig" yaml:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// For more information, see [Amazon EKS Service IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the **Amazon EKS User Guide** .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The encryption configuration for the cluster.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig interface{} `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// The logging configuration for your cluster.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The unique name to give to your cluster.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object representing the configuration of your local Amazon EKS cluster on an AWS Outpost.
	//
	// This object isn't available for clusters on the AWS cloud.
	OutpostConfig interface{} `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// The metadata that you apply to the cluster to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Cluster tags don't propagate to any other resources associated with the cluster.
	//
	// > You must have the `eks:TagResource` and `eks:UntagResource` permissions for your [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the AWS CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the default version available in Amazon EKS is used.
	//
	// > The default version might not be the latest version available.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

