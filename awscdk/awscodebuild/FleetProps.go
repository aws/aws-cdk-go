package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of a CodeBuild Fleet.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fleet := codebuild.NewFleet(this, jsii.String("MyFleet"), &FleetProps{
//   	BaseCapacity: jsii.Number(1),
//   	ComputeType: codebuild.FleetComputeType_CUSTOM_INSTANCE_TYPE,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	ComputeConfiguration: &ComputeConfiguration{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_MEDIUM),
//   		// By default, 64 GiB of disk space is included. Any value optionally
//   		// specified here is _incremental_ on top of the included disk space.
//   		Disk: awscdk.Size_Gibibytes(jsii.Number(10)),
//   	},
//   })
//
type FleetProps struct {
	// The number of machines allocated to the compute ﬂeet. Deﬁnes the number of builds that can run in parallel.
	//
	// Minimum value of 1.
	BaseCapacity *float64 `field:"required" json:"baseCapacity" yaml:"baseCapacity"`
	// The instance type of the compute fleet.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.ComputeType.html
	//
	ComputeType FleetComputeType `field:"required" json:"computeType" yaml:"computeType"`
	// The build environment (operating system/architecture/accelerator) type made available to projects using this fleet.
	EnvironmentType EnvironmentType `field:"required" json:"environmentType" yaml:"environmentType"`
	// The compute configuration of the compute fleet.
	//
	// This is only permitted if `computeType` is set to ATTRIBUTE_BASED or
	// CUSTOM_INSTANCE_TYPE. In such cases, this is required.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.types
	//
	// Default: - do not specify compute configuration.
	//
	ComputeConfiguration *ComputeConfiguration `field:"optional" json:"computeConfiguration" yaml:"computeConfiguration"`
	// The name of the Fleet.
	// Default: - CloudFormation generated name.
	//
	FleetName *string `field:"optional" json:"fleetName" yaml:"fleetName"`
	// The compute fleet overflow behavior.
	//
	// For overflow behavior `QUEUE`, overflow builds need to wait on the existing fleet instances to become available.
	//
	// For overflow behavior `ON_DEMAND`, overflow builds run on CodeBuild on-demand.
	// Default: undefined - AWS CodeBuild default behavior is QUEUE.
	//
	OverflowBehavior FleetOverflowBehavior `field:"optional" json:"overflowBehavior" yaml:"overflowBehavior"`
	// Service Role assumed by Fleet instances.
	//
	// This Role is not used by Project builds running on Fleet instances; Project
	// builds assume the `role` on Project instead.
	// Default: - A role will be created if any permissions are granted.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// What security groups to associate with the fleet's network interfaces. If none are provided, one will be created automatically.
	//
	// Only used if `vpc` is supplied.
	// Default: - A security group will be automatically created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// To access AWS services, your fleet needs to be in one of the following types of subnets:
	//
	// 1. Subnets with access to the internet (of type PRIVATE_WITH_EGRESS).
	// 2. Private subnets unconnected to the internet, but with [VPC endpoints](https://docs.aws.amazon.com/codebuild/latest/userguide/use-vpc-endpoints-with-codebuild.html) for the necessary services.
	//
	// If you don't specify a subnet selection, the default behavior is to use PRIVATE_WITH_EGRESS subnets first if they exist,
	// then PRIVATE_WITHOUT_EGRESS, and finally PUBLIC subnets. If your VPC doesn't have PRIVATE_WITH_EGRESS subnets but you need
	// AWS service access, add VPC Endpoints to your private subnets.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html
	//
	// Default: - private subnets if available else public subnets.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC network to place fleet instance network interfaces.
	//
	// Specify this if the fleet needs to access resources in a VPC.
	// Default: - No VPC is specified.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

