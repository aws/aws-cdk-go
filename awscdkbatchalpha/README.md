# AWS Batch Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

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
awsManagedEnvironment := batch.NewComputeEnvironment(this, jsii.String("AWS-Managed-Compute-Env"), &computeEnvironmentProps{
	computeResources: &computeResources{
		vpc: vpc,
	},
})

customerManagedEnvironment := batch.NewComputeEnvironment(this, jsii.String("Customer-Managed-Compute-Env"), &computeEnvironmentProps{
	managed: jsii.Boolean(false),
})
```

### Spot-Based Compute Environment

It is possible to have AWS Batch submit spotfleet requests for obtaining compute resources. Below is an example of how this can be done:

```go
vpc := ec2.NewVpc(this, jsii.String("VPC"))

spotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MySpotEnvironment"), &computeEnvironmentProps{
	computeResources: &computeResources{
		type: batch.computeResourceType_SPOT,
		bidPercentage: jsii.Number(75),
		 // Bids for resources at 75% of the on-demand price
		vpc: vpc,
	},
})
```

### Fargate Compute Environment

It is possible to have AWS Batch submit jobs to be run on Fargate compute resources. Below is an example of how this can be done:

```go
vpc := ec2.NewVpc(this, jsii.String("VPC"))

fargateSpotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MyFargateEnvironment"), &computeEnvironmentProps{
	computeResources: &computeResources{
		type: batch.computeResourceType_FARGATE_SPOT,
		vpc: vpc,
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

```go
myLaunchTemplate := ec2.NewCfnLaunchTemplate(this, jsii.String("LaunchTemplate"), &cfnLaunchTemplateProps{
	launchTemplateName: jsii.String("extra-storage-template"),
	launchTemplateData: &launchTemplateDataProperty{
		blockDeviceMappings: []interface{}{
			&blockDeviceMappingProperty{
				deviceName: jsii.String("/dev/xvdcz"),
				ebs: &ebsProperty{
					encrypted: jsii.Boolean(true),
					volumeSize: jsii.Number(100),
					volumeType: jsii.String("gp2"),
				},
			},
		},
	},
})
```

And provide `launchTemplateName`:

```go
var vpc vpc
var myLaunchTemplate cfnLaunchTemplate


myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
	computeResources: &computeResources{
		launchTemplate: &launchTemplateSpecification{
			launchTemplateName: string(myLaunchTemplate.launchTemplateName),
		},
		vpc: vpc,
	},
	computeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
})
```

Or provide `launchTemplateId` instead:

```go
var vpc vpc
var myLaunchTemplate cfnLaunchTemplate


myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
	computeResources: &computeResources{
		launchTemplate: &launchTemplateSpecification{
			launchTemplateId: string(myLaunchTemplate.ref),
		},
		vpc: vpc,
	},
	computeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
})
```

### Importing an existing Compute Environment

To import an existing batch compute environment, call `ComputeEnvironment.fromComputeEnvironmentArn()`.

Below is an example:

```go
computeEnv := batch.computeEnvironment.fromComputeEnvironmentArn(this, jsii.String("imported-compute-env"), jsii.String("arn:aws:batch:us-east-1:555555555555:compute-environment/My-Compute-Env"))
```

### Change the baseline AMI of the compute resources

Occasionally, you will need to deviate from the default processing AMI.

ECS Optimized Amazon Linux 2 example:

```go
var vpc vpc

myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
	computeResources: &computeResources{
		image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
			generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
		}),
		vpc: vpc,
	},
})
```

Custom based AMI example:

```go
var vpc vpc

myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
	computeResources: &computeResources{
		image: ec2.machineImage.genericLinux(map[string]*string{
			"[aws-region]": jsii.String("[ami-ID]"),
		}),
		vpc: vpc,
	},
})
```

## Job Queue

Jobs are always submitted to a specific queue. This means that you have to create a queue before you can start submitting jobs. Each queue is mapped to at least one (and no more than three) compute environment. When the job is scheduled for execution, AWS Batch will select the compute environment based on ordinal priority and available capacity in each environment.

```go
var computeEnvironment computeEnvironment

jobQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
	computeEnvironments: []jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			// Defines a collection of compute resources to handle assigned batch jobs
			computeEnvironment: computeEnvironment,
			// Order determines the allocation order for jobs (i.e. Lower means higher preference for job assignment)
			order: jsii.Number(1),
		},
	},
})
```

### Priorty-Based Queue Example

Sometimes you might have jobs that are more important than others, and when submitted, should take precedence over the existing jobs. To achieve this, you can create a priority based execution strategy, by assigning each queue its own priority:

```go
var sharedComputeEnvs computeEnvironment

highPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
	computeEnvironments: []jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			computeEnvironment: sharedComputeEnvs,
			order: jsii.Number(1),
		},
	},
	priority: jsii.Number(2),
})

lowPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
	computeEnvironments: []*jobQueueComputeEnvironment{
		&jobQueueComputeEnvironment{
			computeEnvironment: sharedComputeEnvs,
			order: jsii.Number(1),
		},
	},
	priority: jsii.Number(1),
})
```

By making sure to use the same compute environments between both job queues, we will give precedence to the `highPrioQueue` for the assigning of jobs to available compute environments.

### Importing an existing Job Queue

To import an existing batch job queue, call `JobQueue.fromJobQueueArn()`.

Below is an example:

```go
jobQueue := batch.jobQueue.fromJobQueueArn(this, jsii.String("imported-job-queue"), jsii.String("arn:aws:batch:us-east-1:555555555555:job-queue/High-Prio-Queue"))
```

## Job Definition

A Batch Job definition helps AWS Batch understand important details about how to run your application in the scope of a Batch Job. This involves key information like resource requirements, what containers to run, how the compute environment should be prepared, and more. Below is a simple example of how to create a job definition:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"


repo := ecr.repository.fromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))

batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &jobDefinitionProps{
	container: &jobDefinitionContainer{
		image: ecs.NewEcrImage(repo, jsii.String("latest")),
	},
})
```

### Using a local Docker project

Below is an example of how you can create a Batch Job Definition from a local Docker application.

```go
batch.NewJobDefinition(this, jsii.String("batch-job-def-from-local"), &jobDefinitionProps{
	container: &jobDefinitionContainer{
		// todo-list is a directory containing a Dockerfile to build the application
		image: ecs.containerImage.fromAsset(jsii.String("../todo-list")),
	},
})
```

### Providing custom log configuration

You can provide custom log driver and its configuration for the container.

```go
import ssm "github.com/aws/aws-cdk-go/awscdk"


batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
	container: &jobDefinitionContainer{
		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
		logConfiguration: &logConfiguration{
			logDriver: batch.logDriver_AWSLOGS,
			options: map[string]*string{
				"awslogs-region": jsii.String("us-east-1"),
			},
			secretOptions: []exposedSecret{
				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
			},
		},
	},
})
```

### Using the secret on secrets manager

You can set the environment variables from secrets manager.

```go
dbSecret := secretsmanager.NewSecret(this, jsii.String("secret"))

batch.NewJobDefinition(this, jsii.String("batch-job-def-secrets"), &jobDefinitionProps{
	container: &jobDefinitionContainer{
		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
		secrets: map[string]secret{
			"PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
		},
	},
})
```

### Importing an existing Job Definition

#### From ARN

To import an existing batch job definition from its ARN, call `JobDefinition.fromJobDefinitionArn()`.

Below is an example:

```go
job := batch.jobDefinition.fromJobDefinitionArn(this, jsii.String("imported-job-definition"), jsii.String("arn:aws:batch:us-east-1:555555555555:job-definition/my-job-definition"))
```

#### From Name

To import an existing batch job definition from its name, call `JobDefinition.fromJobDefinitionName()`.
If name is specified without a revision then the latest active revision is used.

Below is an example:

```go
// Without revision
job1 := batch.jobDefinition.fromJobDefinitionName(this, jsii.String("imported-job-definition"), jsii.String("my-job-definition"))

// With revision
job2 := batch.jobDefinition.fromJobDefinitionName(this, jsii.String("imported-job-definition"), jsii.String("my-job-definition:3"))
```
