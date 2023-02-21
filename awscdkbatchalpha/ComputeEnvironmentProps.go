// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a new Compute Environment.
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
type ComputeEnvironmentProps struct {
	// A name for the compute environment.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	ComputeEnvironmentName *string `field:"optional" json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// The details of the required compute resources for the managed compute environment.
	//
	// If specified, and this is an unmanaged compute environment, will throw an error.
	//
	// By default, AWS Batch managed compute environments use a recent, approved version of the
	// Amazon ECS-optimized AMI for compute resources.
	// Experimental.
	ComputeResources *ComputeResources `field:"optional" json:"computeResources" yaml:"computeResources"`
	// The state of the compute environment.
	//
	// If the state is set to true, then the compute
	// environment accepts jobs from a queue and can scale out automatically based on queues.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Determines if AWS should manage the allocation of compute resources for processing jobs.
	//
	// If set to false, then you are in charge of providing the compute resource details.
	// Experimental.
	Managed *bool `field:"optional" json:"managed" yaml:"managed"`
	// The IAM role used by Batch to make calls to other AWS services on your behalf for managing the resources that you use with the service.
	//
	// By default, this role is created for you using
	// the AWS managed service policy for Batch.
	// Experimental.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
}

