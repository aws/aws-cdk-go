package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for configuring an EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var kubernetesVersion kubernetesVersion
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   commonClusterOptions := &CommonClusterOptions{
//   	Version: kubernetesVersion,
//
//   	// the properties below are optional
//   	ClusterName: jsii.String("clusterName"),
//   	OutputClusterName: jsii.Boolean(false),
//   	OutputConfigCommand: jsii.Boolean(false),
//   	Role: role,
//   	SecurityGroup: securityGroup,
//   	Vpc: vpc,
//   	VpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []iSubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type CommonClusterOptions struct {
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
	// the cluster name and, if applicable, the ARN of the masters IAM role.
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
}

