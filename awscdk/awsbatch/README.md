# AWS Batch Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

AWS Batch is a batch processing tool for efficiently running hundreds of thousands computing jobs in AWS. Batch can dynamically provision different types of compute resources based on the resource requirements of submitted jobs.

AWS Batch simplifies the planning, scheduling, and executions of your batch workloads across a full range of compute services like [Amazon EC2](https://aws.amazon.com/ec2/) and [Spot Resources](https://aws.amazon.com/ec2/spot/).

Batch achieves this by utilizing queue processing of batch job requests. To successfully submit a job for execution, you need the following resources:

1. [Job Definition](#job-definition) - *Group various job properties (container image, resource requirements, env variables...) into a single definition. These definitions are used at job submission time.*
2. [Compute Environment](#compute-environment) - *the execution runtime of submitted batch jobs*
3. [Job Queue](#job-queue) - *the queue where batch jobs can be submitted to via AWS SDK/CLI*

For more information on **AWS Batch** visit the [AWS Docs for Batch](https://docs.aws.amazon.com/batch/index.html).

## Compute Environment

At the core of AWS Batch is the compute environment. All batch jobs are processed within a compute environment, which uses resource like OnDemand/Spot EC2 instances or Fargate.

In **MANAGED** mode, AWS will handle the provisioning of compute resources to accommodate the demand. Otherwise, in **UNMANAGED** mode, you will need to manage the provisioning of those resources.

Below is an example of each available type of compute environment:

```go
var vpc vpc


// default is managed
awsManagedEnvironment := batch.NewComputeEnvironment(this, jsii.String("AWS-Managed-Compute-Env"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		Vpc: *Vpc,
	},
})

customerManagedEnvironment := batch.NewComputeEnvironment(this, jsii.String("Customer-Managed-Compute-Env"), &ComputeEnvironmentProps{
	Managed: jsii.Boolean(false),
})
```

### Spot-Based Compute Environment

It is possible to have AWS Batch submit spotfleet requests for obtaining compute resources. Below is an example of how this can be done:

```go
vpc := ec2.NewVpc(this, jsii.String("VPC"))

spotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MySpotEnvironment"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		Type: batch.ComputeResourceType_SPOT,
		BidPercentage: jsii.Number(75),
		 // Bids for resources at 75% of the on-demand price
		Vpc: *Vpc,
	},
})
```

### Fargate Compute Environment

It is possible to have AWS Batch submit jobs to be run on Fargate compute resources. Below is an example of how this can be done:

```go
vpc := ec2.NewVpc(this, jsii.String("VPC"))

fargateSpotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MyFargateEnvironment"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		Type: batch.ComputeResourceType_FARGATE_SPOT,
		Vpc: *Vpc,
	},
})
```

### Understanding Progressive Allocation Strategies

AWS Batch uses an [allocation strategy](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) to determine what compute resource will efficiently handle incoming job requests. By default, **BEST_FIT** will pick an available compute instance based on vCPU requirements. If none exist, the job will wait until resources become available. However, with this strategy, you may have jobs waiting in the queue unnecessarily despite having more powerful instances available. Below is an example of how that situation might look like:

```plaintext
Compute Environment:

1. m5.xlarge => 4 vCPU
2. m5.2xlarge => 8 vCPU
```

```plaintext
Job Queue:
---------
| A | B |
---------

Job Requirements:
A => 4 vCPU - ALLOCATED TO m5.xlarge
B => 2 vCPU - WAITING
```

In this situation, Batch will allocate **Job A** to compute resource #1 because it is the most cost efficient resource that matches the vCPU requirement. However, with this `BEST_FIT` strategy, **Job B** will not be allocated to our other available compute resource even though it is strong enough to handle it. Instead, it will wait until the first job is finished processing or wait a similar `m5.xlarge` resource to be provisioned.

The alternative would be to use the `BEST_FIT_PROGRESSIVE` strategy in order for the remaining job to be handled in larger containers regardless of vCPU requirement and costs.

### Launch template support

Simply define your Launch Template:

```text
// This example is only available in TypeScript
const myLaunchTemplate = new ec2.CfnLaunchTemplate(this, 'LaunchTemplate', {
  launchTemplateName: 'extra-storage-template',
  launchTemplateData: {
    blockDeviceMappings: [
      {
        deviceName: '/dev/xvdcz',
        ebs: {
          encrypted: true,
          volumeSize: 100,
          volumeType: 'gp2',
        },
      },
    ],
  },
});
```

and use it:

```go
var vpc vpc
var myLaunchTemplate cfnLaunchTemplate


myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		LaunchTemplate: &LaunchTemplateSpecification{
			LaunchTemplateName: string(myLaunchTemplate.LaunchTemplateName),
		},
		Vpc: *Vpc,
	},
	ComputeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
})
```

### Importing an existing Compute Environment

To import an existing batch compute environment, call `ComputeEnvironment.fromComputeEnvironmentArn()`.

Below is an example:

```go
computeEnv := batch.ComputeEnvironment_FromComputeEnvironmentArn(this, jsii.String("imported-compute-env"), jsii.String("arn:aws:batch:us-east-1:555555555555:compute-environment/My-Compute-Env"))
```

### Change the baseline AMI of the compute resources

Occasionally, you will need to deviate from the default processing AMI.

ECS Optimized Amazon Linux 2 example:

```go
var vpc vpc

myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		Image: ecs.NewEcsOptimizedAmi(&EcsOptimizedAmiProps{
			Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
		}),
		Vpc: *Vpc,
	},
})
```

Custom based AMI example:

```go
var vpc vpc

myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
	ComputeResources: &ComputeResources{
		Image: ec2.MachineImage_GenericLinux(map[string]*string{
			"[aws-region]": jsii.String("[ami-ID]"),
		}),
		Vpc: *Vpc,
	},
})
```

## Job Queue

Jobs are always submitted to a specific queue. This means that you have to create a queue before you can start submitting jobs. Each queue is mapped to at least one (and no more than three) compute environment. When the job is scheduled for execution, AWS Batch will select the compute environment based on ordinal priority and available capacity in each environment.

```go
var computeEnvironment computeEnvironment

jobQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
	ComputeEnvironments: []jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			// Defines a collection of compute resources to handle assigned batch jobs
			ComputeEnvironment: *ComputeEnvironment,
			// Order determines the allocation order for jobs (i.e. Lower means higher preference for job assignment)
			Order: jsii.Number(1),
		},
	},
})
```

### Priorty-Based Queue Example

Sometimes you might have jobs that are more important than others, and when submitted, should take precedence over the existing jobs. To achieve this, you can create a priority based execution strategy, by assigning each queue its own priority:

```go
var sharedComputeEnvs computeEnvironment

highPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
	ComputeEnvironments: []jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			ComputeEnvironment: sharedComputeEnvs,
			Order: jsii.Number(1),
		},
	},
	Priority: jsii.Number(2),
})

lowPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
	ComputeEnvironments: []*jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			ComputeEnvironment: sharedComputeEnvs,
			Order: jsii.Number(1),
		},
	},
	Priority: jsii.Number(1),
})
```

By making sure to use the same compute environments between both job queues, we will give precedence to the `highPrioQueue` for the assigning of jobs to available compute environments.

### Importing an existing Job Queue

To import an existing batch job queue, call `JobQueue.fromJobQueueArn()`.

Below is an example:

```go
jobQueue := batch.JobQueue_FromJobQueueArn(this, jsii.String("imported-job-queue"), jsii.String("arn:aws:batch:us-east-1:555555555555:job-queue/High-Prio-Queue"))
```

## Job Definition

A Batch Job definition helps AWS Batch understand important details about how to run your application in the scope of a Batch Job. This involves key information like resource requirements, what containers to run, how the compute environment should be prepared, and more. Below is a simple example of how to create a job definition:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"


repo := ecr.Repository_FromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))

batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &JobDefinitionProps{
	Container: &JobDefinitionContainer{
		Image: ecs.NewEcrImage(repo, jsii.String("latest")),
	},
})
```

### Using a local Docker project

Below is an example of how you can create a Batch Job Definition from a local Docker application.

```go
batch.NewJobDefinition(this, jsii.String("batch-job-def-from-local"), &JobDefinitionProps{
	Container: &JobDefinitionContainer{
		// todo-list is a directory containing a Dockerfile to build the application
		Image: ecs.ContainerImage_FromAsset(jsii.String("../todo-list")),
	},
})
```

### Providing custom log configuration

You can provide custom log driver and its configuration for the container.

```go
import ssm "github.com/aws/aws-cdk-go/awscdk"


batch.NewJobDefinition(this, jsii.String("job-def"), &JobDefinitionProps{
	Container: &JobDefinitionContainer{
		Image: ecs.EcrImage_FromRegistry(jsii.String("docker/whalesay")),
		LogConfiguration: &LogConfiguration{
			LogDriver: batch.LogDriver_AWSLOGS,
			Options: map[string]*string{
				"awslogs-region": jsii.String("us-east-1"),
			},
			SecretOptions: []exposedSecret{
				batch.*exposedSecret_FromParametersStore(jsii.String("xyz"), ssm.StringParameter_FromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
			},
		},
	},
})
```

### Importing an existing Job Definition

#### From ARN

To import an existing batch job definition from its ARN, call `JobDefinition.fromJobDefinitionArn()`.

Below is an example:

```go
job := batch.JobDefinition_FromJobDefinitionArn(this, jsii.String("imported-job-definition"), jsii.String("arn:aws:batch:us-east-1:555555555555:job-definition/my-job-definition"))
```

#### From Name

To import an existing batch job definition from its name, call `JobDefinition.fromJobDefinitionName()`.
If name is specified without a revision then the latest active revision is used.

Below is an example:

```go
// Without revision
job1 := batch.JobDefinition_FromJobDefinitionName(this, jsii.String("imported-job-definition"), jsii.String("my-job-definition"))

// With revision
job2 := batch.JobDefinition_FromJobDefinitionName(this, jsii.String("imported-job-definition"), jsii.String("my-job-definition:3"))
```
