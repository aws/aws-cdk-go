package awsekslegacy

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties to instantiate the Cluster.
//
// Example:
//   eks.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	DefaultCapacity: jsii.Number(10),
//   	DefaultCapacityInstance: ec2.NewInstanceType(jsii.String("m2.xlarge")),
//   })
//
// Experimental.
type ClusterProps struct {
	// Name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Number of instances to allocate as an initial capacity for this cluster.
	//
	// Instance type can be configured through `defaultCapacityInstanceType`,
	// which defaults to `m5.large`.
	//
	// Use `cluster.addCapacity` to add additional customized capacity. Set this
	// to `0` is you wish to avoid the initial capacity allocation.
	// Experimental.
	DefaultCapacity *float64 `field:"optional" json:"defaultCapacity" yaml:"defaultCapacity"`
	// The instance type to use for the default capacity.
	//
	// This will only be taken
	// into account if `defaultCapacity` is > 0.
	// Experimental.
	DefaultCapacityInstance awsec2.InstanceType `field:"optional" json:"defaultCapacityInstance" yaml:"defaultCapacityInstance"`
	// Allows defining `kubectrl`-related resources on this cluster.
	//
	// If this is disabled, it will not be possible to use the following
	// capabilities:
	// - `addResource`
	// - `addRoleMapping`
	// - `addUserMapping`
	// - `addMastersRole` and `props.mastersRole`
	//
	// If this is disabled, the cluster can only be managed by issuing `kubectl`
	// commands from a session that uses the IAM role/user that created the
	// account.
	//
	// _NOTE_: changing this value will destoy the cluster. This is because a
	// managable cluster must be created using an AWS CloudFormation custom
	// resource which executes with an IAM role owned by the CDK app.
	// Experimental.
	KubectlEnabled *bool `field:"optional" json:"kubectlEnabled" yaml:"kubectlEnabled"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// ```ts
	// const vpcSubnets = [
	//    { subnetType: ec2.SubnetType.PRIVATE_WITH_NAT }
	// ]
	// ```.
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

