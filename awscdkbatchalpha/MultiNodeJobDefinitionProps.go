package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Props to configure a MultiNodeJobDefinition.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiNodeJob := batch.NewMultiNodeJobDefinition(this, jsii.String("JobDefinition"), &MultiNodeJobDefinitionProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R4, ec2.InstanceSize_LARGE),
//   	Containers: []multiNodeContainer{
//   		&multiNodeContainer{
//   			Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("mainMPIContainer"), &EcsEc2ContainerDefinitionProps{
//   				Image: ecs.ContainerImage_FromRegistry(jsii.String("yourregsitry.com/yourMPIImage:latest")),
//   				Cpu: jsii.Number(256),
//   				Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   			StartNode: jsii.Number(0),
//   			EndNode: jsii.Number(5),
//   		},
//   	},
//   })
//   // convenience method
//   multiNodeJob.AddContainer(&multiNodeContainer{
//   	StartNode: jsii.Number(6),
//   	EndNode: jsii.Number(10),
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("multiContainer"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Cpu: jsii.Number(256),
//   		Memory: cdk.Size_*Mebibytes(jsii.Number(2048)),
//   	}),
//   })
//
// Experimental.
type MultiNodeJobDefinitionProps struct {
	// The name of this job definition.
	// Experimental.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Defines the retry behavior for this job.
	// Experimental.
	RetryStrategies *[]RetryStrategy `field:"optional" json:"retryStrategies" yaml:"retryStrategies"`
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Experimental.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The instance type that this job definition will run.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The containers that this multinode job will run.
	// See: https://aws.amazon.com/blogs/compute/building-a-tightly-coupled-molecular-dynamics-workflow-with-multi-node-parallel-jobs-in-aws-batch/
	//
	// Experimental.
	Containers *[]*MultiNodeContainer `field:"optional" json:"containers" yaml:"containers"`
	// The index of the main node in this job.
	//
	// The main node is responsible for orchestration.
	// Experimental.
	MainNode *float64 `field:"optional" json:"mainNode" yaml:"mainNode"`
	// Whether to propogate tags from the JobDefinition to the ECS task that Batch spawns.
	// Experimental.
	PropagateTags *bool `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

