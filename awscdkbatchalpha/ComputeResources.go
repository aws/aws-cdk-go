// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining the structure of the batch compute cluster.
//
// Example:
//   var vpc vpc
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
//   	ComputeResources: &ComputeResources{
//   		Image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
//   			generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		Vpc: *Vpc,
//   	},
//   })
//
// Experimental.
type ComputeResources struct {
	// The VPC network that all compute resources will be connected to.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated.
	//
	// This could be due to availability of the instance type in
	// the region or Amazon EC2 service limits. If this is not specified, the default for the EC2
	// ComputeResourceType is BEST_FIT, which will use only the best fitting instance type, waiting for
	// additional capacity if it's not available. This allocation strategy keeps costs lower but can limit
	// scaling. If you are using Spot Fleets with BEST_FIT then the Spot Fleet IAM Role must be specified.
	// BEST_FIT_PROGRESSIVE will select an additional instance type that is large enough to meet the
	// requirements of the jobs in the queue, with a preference for an instance type with a lower cost.
	// The default value for the SPOT instance type is SPOT_CAPACITY_OPTIMIZED, which is only available for
	// for this type of compute resources and will select an additional instance type that is large enough
	// to meet the requirements of the jobs in the queue, with a preference for an instance type that is
	// less likely to be interrupted.
	// Experimental.
	AllocationStrategy AllocationStrategy `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for
	// that instance type before instances are launched. For example, if your maximum percentage is 20%,
	// then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. You always
	// pay the lowest (market) price and never more than your maximum percentage. If you leave this field empty,
	// the default value is 100% of the On-Demand price.
	// Experimental.
	BidPercentage *float64 `field:"optional" json:"bidPercentage" yaml:"bidPercentage"`
	// Key-value pair tags to be applied to resources that are launched in the compute environment.
	//
	// For AWS Batch, these take the form of "String1": "String2", where String1 is the tag key and
	// String2 is the tag value—for example, { "Name": "AWS Batch Instance - C4OnDemand" }.
	// Experimental.
	ComputeResourcesTags *map[string]*string `field:"optional" json:"computeResourcesTags" yaml:"computeResourcesTags"`
	// The desired number of EC2 vCPUS in the compute environment.
	// Experimental.
	DesiredvCpus *float64 `field:"optional" json:"desiredvCpus" yaml:"desiredvCpus"`
	// The EC2 key pair that is used for instances launched in the compute environment.
	//
	// If no key is defined, then SSH access is not allowed to provisioned compute resources.
	// Experimental.
	Ec2KeyPair *string `field:"optional" json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
	// Experimental.
	Image awsec2.IMachineImage `field:"optional" json:"image" yaml:"image"`
	// The Amazon ECS instance profile applied to Amazon EC2 instances in a compute environment.
	//
	// You can specify
	// the short name or full Amazon Resource Name (ARN) of an instance profile. For example, ecsInstanceRole or
	// arn:aws:iam::<aws_account_id>:instance-profile/ecsInstanceRole . For more information, see Amazon ECS
	// Instance Role in the AWS Batch User Guide.
	// Experimental.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The types of EC2 instances that may be launched in the compute environment.
	//
	// You can specify instance
	// families to launch any instance type within those families (for example, c4 or p3), or you can specify
	// specific sizes within a family (such as c4.8xlarge). You can also choose optimal to pick instance types
	// (from the C, M, and R instance families) on the fly that match the demand of your job queues.
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// An optional launch template to associate with your compute resources.
	//
	// For more information, see README file.
	// Experimental.
	LaunchTemplate *LaunchTemplateSpecification `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of EC2 vCPUs that an environment can reach.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU.
	// Experimental.
	MaxvCpus *float64 `field:"optional" json:"maxvCpus" yaml:"maxvCpus"`
	// The minimum number of EC2 vCPUs that an environment should maintain (even if the compute environment state is DISABLED).
	//
	// Each vCPU is equivalent to 1,024 CPU shares. By keeping this set to 0 you will not have instance time wasted when
	// there is no work to be run. If you set this above zero you will maintain that number of vCPUs at all times.
	// Experimental.
	MinvCpus *float64 `field:"optional" json:"minvCpus" yaml:"minvCpus"`
	// The Amazon EC2 placement group to associate with your compute resources.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Up to 5 EC2 security group(s) associated with instances launched in the compute environment.
	//
	// This parameter is mutually exclusive with launchTemplate.useNetworkInterfaceSecurityGroups
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment.
	// For more information, see Amazon EC2 Spot Fleet Role in the AWS Batch User Guide.
	// Experimental.
	SpotFleetRole awsiam.IRole `field:"optional" json:"spotFleetRole" yaml:"spotFleetRole"`
	// The type of compute environment: ON_DEMAND, SPOT, FARGATE, or FARGATE_SPOT.
	// Experimental.
	Type ComputeResourceType `field:"optional" json:"type" yaml:"type"`
	// The VPC subnets into which the compute resources are launched.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

