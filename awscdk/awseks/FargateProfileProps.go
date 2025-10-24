package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration props for EKS Fargate Profiles.
//
// Example:
//   var cluster Cluster
//
//   eks.NewFargateProfile(this, jsii.String("MyProfile"), &FargateProfileProps{
//   	Cluster: Cluster,
//   	Selectors: []Selector{
//   		&Selector{
//   			Namespace: jsii.String("default"),
//   		},
//   	},
//   })
//
type FargateProfileProps struct {
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace.
	//
	// At least one selector is required and you may specify up to five selectors.
	Selectors *[]*Selector `field:"required" json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	// Default: - generated.
	//
	FargateProfileName *string `field:"optional" json:"fargateProfileName" yaml:"fargateProfileName"`
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
	//
	// Default: - a role will be automatically created.
	//
	PodExecutionRole awsiam.IRole `field:"optional" json:"podExecutionRole" yaml:"podExecutionRole"`
	// Select which subnets to launch your pods into.
	//
	// At this time, pods running
	// on Fargate are not assigned public IP addresses, so only private subnets
	// (with no direct route to an Internet Gateway) are allowed.
	//
	// You must specify the VPC to customize the subnet selection.
	// Default: - all private subnets of the VPC are selected.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC from which to select subnets to launch your pods into.
	//
	// By default, all private subnets are selected. You can customize this using
	// `subnetSelection`.
	// Default: - all private subnets used by the EKS cluster.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The EKS cluster to apply the Fargate profile to.
	//
	// [disable-awslint:ref-via-interface].
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

