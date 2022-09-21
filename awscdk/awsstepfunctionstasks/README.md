# Tasks for AWS Step Functions

[AWS Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) is a web service that enables you to coordinate the
components of distributed applications and microservices using visual workflows.
You build applications from individual components that each perform a discrete
function, or task, allowing you to scale and change applications quickly.

A [Task](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-task-state.html) state represents a single unit of work performed by a state machine.
All work in your state machine is performed by tasks.  This module contains a collection of classes that allow you to call various AWS services
from your Step Functions state machine.

Be sure to familiarize yourself with the [`aws-stepfunctions` module documentation](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_stepfunctions-readme.html) first.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Table Of Contents

* [Tasks for AWS Step Functions](#tasks-for-aws-step-functions)

  * [Table Of Contents](#table-of-contents)
  * [Evaluate Expression](#evaluate-expression)
  * [API Gateway](#api-gateway)

    * [Call REST API Endpoint](#call-rest-api-endpoint)
    * [Call HTTP API Endpoint](#call-http-api-endpoint)
  * [AWS SDK](#aws-sdk)
  * [Athena](#athena)

    * [StartQueryExecution](#startqueryexecution)
    * [GetQueryExecution](#getqueryexecution)
    * [GetQueryResults](#getqueryresults)
    * [StopQueryExecution](#stopqueryexecution)
  * [Batch](#batch)

    * [SubmitJob](#submitjob)
  * [CodeBuild](#codebuild)

    * [StartBuild](#startbuild)
  * [DynamoDB](#dynamodb)

    * [GetItem](#getitem)
    * [PutItem](#putitem)
    * [DeleteItem](#deleteitem)
    * [UpdateItem](#updateitem)
  * [ECS](#ecs)

    * [RunTask](#runtask)

      * [EC2](#ec2)
      * [Fargate](#fargate)
  * [EMR](#emr)

    * [Create Cluster](#create-cluster)
    * [Termination Protection](#termination-protection)
    * [Terminate Cluster](#terminate-cluster)
    * [Add Step](#add-step)
    * [Cancel Step](#cancel-step)
    * [Modify Instance Fleet](#modify-instance-fleet)
    * [Modify Instance Group](#modify-instance-group)
  * [EMR on EKS](#emr-on-eks)

    * [Create Virtual Cluster](#create-virtual-cluster)
    * [Delete Virtual Cluster](#delete-virtual-cluster)
    * [Start Job Run](#start-job-run)
  * [EKS](#eks)

    * [Call](#call)
  * [EventBridge](#eventbridge)

    * [Put Events](#put-events)
  * [Glue](#glue)
  * [Glue DataBrew](#glue-databrew)
  * [Lambda](#lambda)
  * [SageMaker](#sagemaker)

    * [Create Training Job](#create-training-job)
    * [Create Transform Job](#create-transform-job)
    * [Create Endpoint](#create-endpoint)
    * [Create Endpoint Config](#create-endpoint-config)
    * [Create Model](#create-model)
    * [Update Endpoint](#update-endpoint)
  * [SNS](#sns)
  * [Step Functions](#step-functions)

    * [Start Execution](#start-execution)
    * [Invoke Activity](#invoke-activity)
  * [SQS](#sqs)

## Paths

Learn more about input and output processing in Step Functions [here](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html)

## Evaluate Expression

Use the `EvaluateExpression` to perform simple operations referencing state paths. The
`expression` referenced in the task will be evaluated in a Lambda function
(`eval()`). This allows you to not have to write Lambda code for simple operations.

Example: convert a wait time from milliseconds to seconds, concat this in a message and wait:

```go
convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &evaluateExpressionProps{
	expression: jsii.String("$.waitMilliseconds / 1000"),
	resultPath: jsii.String("$.waitSeconds"),
})

createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &evaluateExpressionProps{
	// Note: this is a string inside a string.
	expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
	runtime: lambda.runtime_NODEJS_14_X(),
	resultPath: jsii.String("$.message"),
})

publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &snsPublishProps{
	topic: sns.NewTopic(this, jsii.String("cool-topic")),
	message: sfn.taskInput.fromJsonPathAt(jsii.String("$.message")),
	resultPath: jsii.String("$.sns"),
})

wait := sfn.NewWait(this, jsii.String("Wait"), &waitProps{
	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
})

sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
	definition: convertToSeconds.next(createMessage).next(publishMessage).next(wait),
})
```

The `EvaluateExpression` supports a `runtime` prop to specify the Lambda
runtime to use to evaluate the expression. Currently, only runtimes
of the Node.js family are supported.

## API Gateway

Step Functions supports [API Gateway](https://docs.aws.amazon.com/step-functions/latest/dg/connect-api-gateway.html) through the service integration pattern.

HTTP APIs are designed for low-latency, cost-effective integrations with AWS services, including AWS Lambda, and HTTP endpoints.
HTTP APIs support OIDC and OAuth 2.0 authorization, and come with built-in support for CORS and automatic deployments.
Previous-generation REST APIs currently offer more features. More details can be found [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vs-rest.html).

### Call REST API Endpoint

The `CallApiGatewayRestApiEndpoint` calls the REST API endpoint.

```go
import apigateway "github.com/aws/aws-cdk-go/awscdk"

restApi := apigateway.NewRestApi(this, jsii.String("MyRestApi"))

invokeTask := tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Call REST API"), &callApiGatewayRestApiEndpointProps{
	api: restApi,
	stageName: jsii.String("prod"),
	method: tasks.httpMethod_GET,
})
```

Be aware that the header values must be arrays. When passing the Task Token
in the headers field `WAIT_FOR_TASK_TOKEN` integration, use
`JsonPath.array()` to wrap the token in an array:

```go
import apigateway "github.com/aws/aws-cdk-go/awscdk"
var api restApi


tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Endpoint"), &callApiGatewayRestApiEndpointProps{
	api: api,
	stageName: jsii.String("Stage"),
	method: tasks.httpMethod_PUT,
	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
	headers: sfn.taskInput.fromObject(map[string]interface{}{
		"TaskToken": sfn.JsonPath.array(sfn.JsonPath.taskToken),
	}),
})
```

### Call HTTP API Endpoint

The `CallApiGatewayHttpApiEndpoint` calls the HTTP API endpoint.

```go
// Example automatically generated from non-compiling source. May contain errors.
import apigatewayv2 "github.com/aws/aws-cdk-go/awscdk"

httpApi := apigatewayv2.NewHttpApi(this, jsii.String("MyHttpApi"))

invokeTask := tasks.NewCallApiGatewayHttpApiEndpoint(this, jsii.String("Call HTTP API"), &callApiGatewayHttpApiEndpointProps{
	apiId: httpApi.apiId,
	apiStack: awscdk.*stack.of(httpApi),
	method: tasks.httpMethod_GET,
})
```

### AWS SDK

Step Functions supports calling [AWS service's API actions](https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html)
through the service integration pattern.

You can use Step Functions' AWS SDK integrations to call any of the over two hundred AWS services
directly from your state machine, giving you access to over nine thousand API actions.

```go
var myBucket bucket

getObject := tasks.NewCallAwsService(this, jsii.String("GetObject"), &callAwsServiceProps{
	service: jsii.String("s3"),
	action: jsii.String("getObject"),
	parameters: map[string]interface{}{
		"Bucket": myBucket.bucketName,
		"Key": sfn.JsonPath.stringAt(jsii.String("$.key")),
	},
	iamResources: []*string{
		myBucket.arnForObjects(jsii.String("*")),
	},
})
```

Use camelCase for actions and PascalCase for parameter names.

The task automatically adds an IAM statement to the state machine role's policy based on the
service and action called. The resources for this statement must be specified in `iamResources`.

Use the `iamAction` prop to manually specify the IAM action name in the case where the IAM
action name does not match with the API service/action name:

```go
listBuckets := tasks.NewCallAwsService(this, jsii.String("ListBuckets"), &callAwsServiceProps{
	service: jsii.String("s3"),
	action: jsii.String("listBuckets"),
	iamResources: []*string{
		jsii.String("*"),
	},
	iamAction: jsii.String("s3:ListAllMyBuckets"),
})
```

Use the `additionalIamStatements` prop to pass additional IAM statements that will be added to the
state machine role's policy. Use it in the case where the call requires more than a single statement
to be executed:

```go
// Example automatically generated from non-compiling source. May contain errors.
detectLabels := tasks.NewCallAwsService(stack, jsii.String("DetectLabels"), &callAwsServiceProps{
	service: jsii.String("rekognition"),
	action: jsii.String("detectLabels"),
	iamResources: []*string{
		jsii.String("*"),
	},
	additionalIamStatements: []policyStatement{
		iam.NewPolicyStatement(&policyStatementProps{
			actions: []*string{
				jsii.String("s3:getObject"),
			},
			resources: []*string{
				jsii.String("arn:aws:s3:::my-bucket/*"),
			},
		}),
	},
})
```

## Athena

Step Functions supports [Athena](https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html) through the service integration pattern.

### StartQueryExecution

The [StartQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_StartQueryExecution.html) API runs the SQL query statement.

```go
startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &athenaStartQueryExecutionProps{
	queryString: sfn.jsonPath.stringAt(jsii.String("$.queryString")),
	queryExecutionContext: &queryExecutionContext{
		databaseName: jsii.String("mydatabase"),
	},
	resultConfiguration: &resultConfiguration{
		encryptionConfiguration: &encryptionConfiguration{
			encryptionOption: tasks.encryptionOption_S3_MANAGED,
		},
		outputLocation: &location{
			bucketName: jsii.String("query-results-bucket"),
			objectKey: jsii.String("folder"),
		},
	},
})
```

### GetQueryExecution

The [GetQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetQueryExecution.html) API gets information about a single execution of a query.

```go
getQueryExecutionJob := tasks.NewAthenaGetQueryExecution(this, jsii.String("Get Query Execution"), &athenaGetQueryExecutionProps{
	queryExecutionId: sfn.jsonPath.stringAt(jsii.String("$.QueryExecutionId")),
})
```

### GetQueryResults

The [GetQueryResults](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetQueryResults.html) API that streams the results of a single query execution specified by QueryExecutionId from S3.

```go
getQueryResultsJob := tasks.NewAthenaGetQueryResults(this, jsii.String("Get Query Results"), &athenaGetQueryResultsProps{
	queryExecutionId: sfn.jsonPath.stringAt(jsii.String("$.QueryExecutionId")),
})
```

### StopQueryExecution

The [StopQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_StopQueryExecution.html) API that stops a query execution.

```go
stopQueryExecutionJob := tasks.NewAthenaStopQueryExecution(this, jsii.String("Stop Query Execution"), &athenaStopQueryExecutionProps{
	queryExecutionId: sfn.jsonPath.stringAt(jsii.String("$.QueryExecutionId")),
})
```

## Batch

Step Functions supports [Batch](https://docs.aws.amazon.com/step-functions/latest/dg/connect-batch.html) through the service integration pattern.

### SubmitJob

The [SubmitJob](https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html) API submits an AWS Batch job from a job definition.

```go
// Example automatically generated from non-compiling source. May contain errors.
import batch "github.com/aws/aws-cdk-go/awscdk"
var batchJobDefinition batch.JobDefinition
var batchQueue batch.JobQueue


task := tasks.NewBatchSubmitJob(this, jsii.String("Submit Job"), &batchSubmitJobProps{
	jobDefinitionArn: batchJobDefinition.jobDefinitionArn,
	jobName: jsii.String("MyJob"),
	jobQueueArn: batchQueue.jobQueueArn,
})
```

## CodeBuild

Step Functions supports [CodeBuild](https://docs.aws.amazon.com/step-functions/latest/dg/connect-codebuild.html) through the service integration pattern.

### StartBuild

[StartBuild](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuild.html) starts a CodeBuild Project by Project Name.

```go
import codebuild "github.com/aws/aws-cdk-go/awscdk"


codebuildProject := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	projectName: jsii.String("MyTestProject"),
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("echo \"Hello, CodeBuild!\""),
				},
			},
		},
	}),
})

task := tasks.NewCodeBuildStartBuild(this, jsii.String("Task"), &codeBuildStartBuildProps{
	project: codebuildProject,
	integrationPattern: sfn.integrationPattern_RUN_JOB,
	environmentVariablesOverride: map[string]buildEnvironmentVariable{
		"ZONE": &buildEnvironmentVariable{
			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
			"value": sfn.JsonPath.stringAt(jsii.String("$.envVariables.zone")),
		},
	},
})
```

## DynamoDB

You can call DynamoDB APIs from a `Task` state.
Read more about calling DynamoDB APIs [here](https://docs.aws.amazon.com/step-functions/latest/dg/connect-ddb.html)

### GetItem

The [GetItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html) operation returns a set of attributes for the item with the given primary key.

```go
var myTable table

tasks.NewDynamoGetItem(this, jsii.String("Get Item"), &dynamoGetItemProps{
	key: map[string]dynamoAttributeValue{
		"messageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
	},
	table: myTable,
})
```

### PutItem

The [PutItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html) operation creates a new item, or replaces an old item with a new item.

```go
var myTable table

tasks.NewDynamoPutItem(this, jsii.String("PutItem"), &dynamoPutItemProps{
	item: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
		"Text": tasks.*dynamoAttributeValue.fromString(sfn.JsonPath.stringAt(jsii.String("$.bar"))),
		"TotalCount": tasks.*dynamoAttributeValue.fromNumber(jsii.Number(10)),
	},
	table: myTable,
})
```

### DeleteItem

The [DeleteItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html) operation deletes a single item in a table by primary key.

```go
var myTable table

tasks.NewDynamoDeleteItem(this, jsii.String("DeleteItem"), &dynamoDeleteItemProps{
	key: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
	},
	table: myTable,
	resultPath: sfn.jsonPath_DISCARD(),
})
```

### UpdateItem

The [UpdateItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html) operation edits an existing item's attributes, or adds a new item
to the table if it does not already exist.

```go
var myTable table

tasks.NewDynamoUpdateItem(this, jsii.String("UpdateItem"), &dynamoUpdateItemProps{
	key: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
	},
	table: myTable,
	expressionAttributeValues: map[string]*dynamoAttributeValue{
		":val": tasks.*dynamoAttributeValue.numberFromString(sfn.JsonPath.stringAt(jsii.String("$.Item.TotalCount.N"))),
		":rand": tasks.*dynamoAttributeValue.fromNumber(jsii.Number(20)),
	},
	updateExpression: jsii.String("SET TotalCount = :val + :rand"),
})
```

## ECS

Step Functions supports [ECS/Fargate](https://docs.aws.amazon.com/step-functions/latest/dg/connect-ecs.html) through the service integration pattern.

### RunTask

[RunTask](https://docs.aws.amazon.com/step-functions/latest/dg/connect-ecs.html) starts a new task using the specified task definition.

#### EC2

The EC2 launch type allows you to run your containerized applications on a cluster
of Amazon EC2 instances that you manage.

When a task that uses the EC2 launch type is launched, Amazon ECS must determine where
to place the task based on the requirements specified in the task definition, such as
CPU and memory. Similarly, when you scale down the task count, Amazon ECS must determine
which tasks to terminate. You can apply task placement strategies and constraints to
customize how Amazon ECS places and terminates tasks. Learn more about [task placement](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement.html)

The latest ACTIVE revision of the passed task definition is used for running the task.

The following example runs a job from a task definition on EC2

```go
vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
	isDefault: jsii.Boolean(true),
})

cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &clusterProps{
	vpc: vpc,
})
cluster.addCapacity(jsii.String("DefaultAutoScalingGroup"), &addCapacityOptions{
	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	},
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
	compatibility: ecs.compatibility_EC2,
})

taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
	memoryLimitMiB: jsii.Number(256),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &ecsRunTaskProps{
	integrationPattern: sfn.integrationPattern_RUN_JOB,
	cluster: cluster,
	taskDefinition: taskDefinition,
	launchTarget: tasks.NewEcsEc2LaunchTarget(&ecsEc2LaunchTargetOptions{
		placementStrategies: []placementStrategy{
			ecs.*placementStrategy.spreadAcrossInstances(),
			ecs.*placementStrategy.packedByCpu(),
			ecs.*placementStrategy.randomly(),
		},
		placementConstraints: []placementConstraint{
			ecs.*placementConstraint.memberOf(jsii.String("blieptuut")),
		},
	}),
})
```

#### Fargate

AWS Fargate is a serverless compute engine for containers that works with Amazon
Elastic Container Service (ECS). Fargate makes it easy for you to focus on building
your applications. Fargate removes the need to provision and manage servers, lets you
specify and pay for resources per application, and improves security through application
isolation by design. Learn more about [Fargate](https://aws.amazon.com/fargate/)

The Fargate launch type allows you to run your containerized applications without the need
to provision and manage the backend infrastructure. Just register your task definition and
Fargate launches the container for you. The latest ACTIVE revision of the passed
task definition is used for running the task. Learn more about
[Fargate Versioning](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTaskDefinition.html)

The following example runs a job from a task definition on Fargate

```go
vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
	isDefault: jsii.Boolean(true),
})

cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &clusterProps{
	vpc: vpc,
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
	memoryMiB: jsii.String("512"),
	cpu: jsii.String("256"),
	compatibility: ecs.compatibility_FARGATE,
})

containerDefinition := taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
	memoryLimitMiB: jsii.Number(256),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &ecsRunTaskProps{
	integrationPattern: sfn.integrationPattern_RUN_JOB,
	cluster: cluster,
	taskDefinition: taskDefinition,
	assignPublicIp: jsii.Boolean(true),
	containerOverrides: []containerOverride{
		&containerOverride{
			containerDefinition: containerDefinition,
			environment: []taskEnvironmentVariable{
				&taskEnvironmentVariable{
					name: jsii.String("SOME_KEY"),
					value: sfn.jsonPath.stringAt(jsii.String("$.SomeKey")),
				},
			},
		},
	},
	launchTarget: tasks.NewEcsFargateLaunchTarget(),
})
```

## EMR

Step Functions supports Amazon EMR through the service integration pattern.
The service integration APIs correspond to Amazon EMR APIs but differ in the
parameters that are used.

[Read more](https://docs.aws.amazon.com/step-functions/latest/dg/connect-emr.html) about the differences when using these service integrations.

### Create Cluster

Creates and starts running a cluster (job flow).
Corresponds to the [`runJobFlow`](https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html) API in EMR.

```go
clusterRole := iam.NewRole(this, jsii.String("ClusterRole"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
})

serviceRole := iam.NewRole(this, jsii.String("ServiceRole"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
})

autoScalingRole := iam.NewRole(this, jsii.String("AutoScalingRole"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
})

autoScalingRole.assumeRolePolicy.addStatements(
iam.NewPolicyStatement(&policyStatementProps{
	effect: iam.effect_ALLOW,
	principals: []iPrincipal{
		iam.NewServicePrincipal(jsii.String("application-autoscaling.amazonaws.com")),
	},
	actions: []*string{
		jsii.String("sts:AssumeRole"),
	},
}))

tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &emrCreateClusterProps{
	instances: &instancesConfigProperty{
	},
	clusterRole: clusterRole,
	name: sfn.taskInput.fromJsonPathAt(jsii.String("$.ClusterName")).value,
	serviceRole: serviceRole,
	autoScalingRole: autoScalingRole,
})
```

If you want to run multiple steps in [parallel](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-concurrent-steps.html),
you can specify the `stepConcurrencyLevel` property. The concurrency range is between 1
and 256 inclusive, where the default concurrency of 1 means no step concurrency is allowed.
`stepConcurrencyLevel` requires the EMR release label to be 5.28.0 or above.

```go
tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &emrCreateClusterProps{
	instances: &instancesConfigProperty{
	},
	name: sfn.taskInput.fromJsonPathAt(jsii.String("$.ClusterName")).value,
	stepConcurrencyLevel: jsii.Number(10),
})
```

### Termination Protection

Locks a cluster (job flow) so the EC2 instances in the cluster cannot be
terminated by user intervention, an API call, or a job-flow error.

Corresponds to the [`setTerminationProtection`](https://docs.aws.amazon.com/step-functions/latest/dg/connect-emr.html) API in EMR.

```go
tasks.NewEmrSetClusterTerminationProtection(this, jsii.String("Task"), &emrSetClusterTerminationProtectionProps{
	clusterId: jsii.String("ClusterId"),
	terminationProtected: jsii.Boolean(false),
})
```

### Terminate Cluster

Shuts down a cluster (job flow).
Corresponds to the [`terminateJobFlows`](https://docs.aws.amazon.com/emr/latest/APIReference/API_TerminateJobFlows.html) API in EMR.

```go
tasks.NewEmrTerminateCluster(this, jsii.String("Task"), &emrTerminateClusterProps{
	clusterId: jsii.String("ClusterId"),
})
```

### Add Step

Adds a new step to a running cluster.
Corresponds to the [`addJobFlowSteps`](https://docs.aws.amazon.com/emr/latest/APIReference/API_AddJobFlowSteps.html) API in EMR.

```go
tasks.NewEmrAddStep(this, jsii.String("Task"), &emrAddStepProps{
	clusterId: jsii.String("ClusterId"),
	name: jsii.String("StepName"),
	jar: jsii.String("Jar"),
	actionOnFailure: tasks.actionOnFailure_CONTINUE,
})
```

### Cancel Step

Cancels a pending step in a running cluster.
Corresponds to the [`cancelSteps`](https://docs.aws.amazon.com/emr/latest/APIReference/API_CancelSteps.html) API in EMR.

```go
tasks.NewEmrCancelStep(this, jsii.String("Task"), &emrCancelStepProps{
	clusterId: jsii.String("ClusterId"),
	stepId: jsii.String("StepId"),
})
```

### Modify Instance Fleet

Modifies the target On-Demand and target Spot capacities for the instance
fleet with the specified InstanceFleetName.

Corresponds to the [`modifyInstanceFleet`](https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceFleet.html) API in EMR.

```go
tasks.NewEmrModifyInstanceFleetByName(this, jsii.String("Task"), &emrModifyInstanceFleetByNameProps{
	clusterId: jsii.String("ClusterId"),
	instanceFleetName: jsii.String("InstanceFleetName"),
	targetOnDemandCapacity: jsii.Number(2),
	targetSpotCapacity: jsii.Number(0),
})
```

### Modify Instance Group

Modifies the number of nodes and configuration settings of an instance group.

Corresponds to the [`modifyInstanceGroups`](https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceGroups.html) API in EMR.

```go
tasks.NewEmrModifyInstanceGroupByName(this, jsii.String("Task"), &emrModifyInstanceGroupByNameProps{
	clusterId: jsii.String("ClusterId"),
	instanceGroupName: sfn.jsonPath.stringAt(jsii.String("$.InstanceGroupName")),
	instanceGroup: &instanceGroupModifyConfigProperty{
		instanceCount: jsii.Number(1),
	},
})
```

## EMR on EKS

Step Functions supports Amazon EMR on EKS through the service integration pattern.
The service integration APIs correspond to Amazon EMR on EKS APIs, but differ in the parameters that are used.

[Read more](https://docs.aws.amazon.com/step-functions/latest/dg/connect-emr-eks.html) about the differences when using these service integrations.

[Setting up](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up.html) the EKS cluster is required.

### Create Virtual Cluster

The [CreateVirtualCluster](https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_CreateVirtualCluster.html) API creates a single virtual cluster that's mapped to a single Kubernetes namespace.

The EKS cluster containing the Kubernetes namespace where the virtual cluster will be mapped can be passed in from the task input.

```go
tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &emrContainersCreateVirtualClusterProps{
	eksCluster: tasks.eksClusterInput.fromTaskInput(sfn.taskInput.fromText(jsii.String("clusterId"))),
})
```

The EKS cluster can also be passed in directly.

```go
import eks "github.com/aws/aws-cdk-go/awscdk"

var eksCluster cluster


tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &emrContainersCreateVirtualClusterProps{
	eksCluster: tasks.eksClusterInput.fromCluster(eksCluster),
})
```

By default, the Kubernetes namespace that a virtual cluster maps to is "default", but a specific namespace within an EKS cluster can be selected.

```go
tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &emrContainersCreateVirtualClusterProps{
	eksCluster: tasks.eksClusterInput.fromTaskInput(sfn.taskInput.fromText(jsii.String("clusterId"))),
	eksNamespace: jsii.String("specified-namespace"),
})
```

### Delete Virtual Cluster

The [DeleteVirtualCluster](https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_DeleteVirtualCluster.html) API deletes a virtual cluster.

```go
tasks.NewEmrContainersDeleteVirtualCluster(this, jsii.String("Delete a Virtual Cluster"), &emrContainersDeleteVirtualClusterProps{
	virtualClusterId: sfn.taskInput.fromJsonPathAt(jsii.String("$.virtualCluster")),
})
```

### Start Job Run

The [StartJobRun](https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_StartJobRun.html) API starts a job run. A job is a unit of work that you submit to Amazon EMR on EKS for execution. The work performed by the job can be defined by a Spark jar, PySpark script, or SparkSQL query. A job run is an execution of the job on the virtual cluster.

Required setup:

* If not done already, follow the [steps](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up.html) to setup EMR on EKS and [create an EKS Cluster](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-eks-readme.html#quick-start).
* Enable [Cluster access](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-cluster-access.html)
* Enable [IAM Role access](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-enable-IAM.html)

The following actions must be performed if the virtual cluster ID is supplied from the task input. Otherwise, if it is supplied statically in the state machine definition, these actions will be done automatically.

* Create an [IAM role](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-iam.Role.html)
* Update the [Role Trust Policy](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-trust-policy.html) of the Job Execution Role.

The job can be configured with spark submit parameters:

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
	jobDriver: &jobDriver{
		sparkSubmitJobDriver: &sparkSubmitJobDriver{
			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
})
```

Configuring the job can also be done via application configuration:

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
	jobName: jsii.String("EMR-Containers-Job"),
	jobDriver: &jobDriver{
		sparkSubmitJobDriver: &sparkSubmitJobDriver{
			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
		},
	},
	applicationConfig: []applicationConfiguration{
		&applicationConfiguration{
			classification: tasks.classification_SPARK_DEFAULTS(),
			properties: map[string]*string{
				"spark.executor.instances": jsii.String("1"),
				"spark.executor.memory": jsii.String("512M"),
			},
		},
	},
})
```

Job monitoring can be enabled if `monitoring.logging` is set true. This automatically generates an S3 bucket and CloudWatch logs.

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
	jobDriver: &jobDriver{
		sparkSubmitJobDriver: &sparkSubmitJobDriver{
			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
	monitoring: &monitoring{
		logging: jsii.Boolean(true),
	},
})
```

Otherwise, providing monitoring for jobs with existing log groups and log buckets is also available.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
logBucket := s3.NewBucket(this, jsii.String("S3 Bucket"))

tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
	jobDriver: &jobDriver{
		sparkSubmitJobDriver: &sparkSubmitJobDriver{
			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
	monitoring: &monitoring{
		logGroup: logGroup,
		logBucket: logBucket,
	},
})
```

Users can provide their own existing Job Execution Role.

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
	virtualCluster: tasks.virtualClusterInput.fromTaskInput(sfn.taskInput.fromJsonPathAt(jsii.String("$.VirtualClusterId"))),
	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
	jobName: jsii.String("EMR-Containers-Job"),
	executionRole: iam.role.fromRoleArn(this, jsii.String("Job-Execution-Role"), jsii.String("arn:aws:iam::xxxxxxxxxxxx:role/JobExecutionRole")),
	jobDriver: &jobDriver{
		sparkSubmitJobDriver: &sparkSubmitJobDriver{
			entryPoint: sfn.*taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
})
```

## EKS

Step Functions supports Amazon EKS through the service integration pattern.
The service integration APIs correspond to Amazon EKS APIs.

[Read more](https://docs.aws.amazon.com/step-functions/latest/dg/connect-eks.html) about the differences when using these service integrations.

### Call

Read and write Kubernetes resource objects via a Kubernetes API endpoint.
Corresponds to the [`call`](https://docs.aws.amazon.com/step-functions/latest/dg/connect-eks.html) API in Step Functions Connector.

The following code snippet includes a Task state that uses eks:call to list the pods.

```go
import eks "github.com/aws/aws-cdk-go/awscdk"


myEksCluster := eks.NewCluster(this, jsii.String("my sample cluster"), &clusterProps{
	version: eks.kubernetesVersion_V1_18(),
	clusterName: jsii.String("myEksCluster"),
})

tasks.NewEksCall(this, jsii.String("Call a EKS Endpoint"), &eksCallProps{
	cluster: myEksCluster,
	httpMethod: tasks.httpMethods_GET,
	httpPath: jsii.String("/api/v1/namespaces/default/pods"),
})
```

## EventBridge

Step Functions supports Amazon EventBridge through the service integration pattern.
The service integration APIs correspond to Amazon EventBridge APIs.

[Read more](https://docs.aws.amazon.com/step-functions/latest/dg/connect-eventbridge.html) about the differences when using these service integrations.

### Put Events

Send events to an EventBridge bus.
Corresponds to the [`put-events`](https://docs.aws.amazon.com/step-functions/latest/dg/connect-eventbridge.html) API in Step Functions Connector.

The following code snippet includes a Task state that uses events:putevents to send an event to the default bus.

```go
import events "github.com/aws/aws-cdk-go/awscdk"


myEventBus := events.NewEventBus(this, jsii.String("EventBus"), &eventBusProps{
	eventBusName: jsii.String("MyEventBus1"),
})

tasks.NewEventBridgePutEvents(this, jsii.String("Send an event to EventBridge"), &eventBridgePutEventsProps{
	entries: []eventBridgePutEventsEntry{
		&eventBridgePutEventsEntry{
			detail: sfn.taskInput.fromObject(map[string]interface{}{
				"Message": jsii.String("Hello from Step Functions!"),
			}),
			eventBus: myEventBus,
			detailType: jsii.String("MessageFromStepFunctions"),
			source: jsii.String("step.functions"),
		},
	},
})
```

## Glue

Step Functions supports [AWS Glue](https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html) through the service integration pattern.

You can call the [`StartJobRun`](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-StartJobRun) API from a `Task` state.

```go
tasks.NewGlueStartJobRun(this, jsii.String("Task"), &glueStartJobRunProps{
	glueJobName: jsii.String("my-glue-job"),
	arguments: sfn.taskInput.fromObject(map[string]interface{}{
		"key": jsii.String("value"),
	}),
	timeout: awscdk.Duration.minutes(jsii.Number(30)),
	notifyDelayAfter: awscdk.Duration.minutes(jsii.Number(5)),
})
```

## Glue DataBrew

Step Functions supports [AWS Glue DataBrew](https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html) through the service integration pattern.

You can call the [`StartJobRun`](https://docs.aws.amazon.com/databrew/latest/dg/API_StartJobRun.html) API from a `Task` state.

```go
tasks.NewGlueDataBrewStartJobRun(this, jsii.String("Task"), &glueDataBrewStartJobRunProps{
	name: jsii.String("databrew-job"),
})
```

## Lambda

[Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html) a Lambda function.

You can specify the input to your Lambda function through the `payload` attribute.
By default, Step Functions invokes Lambda function with the state input (JSON path '$')
as the input.

The following snippet invokes a Lambda Function with the state input as the payload
by referencing the `$` path.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke with state input"), &lambdaInvokeProps{
	lambdaFunction: fn,
})
```

When a function is invoked, the Lambda service sends  [these response
elements](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_ResponseElements)
back.

⚠️ The response from the Lambda function is in an attribute called `Payload`

The following snippet invokes a Lambda Function by referencing the `$.Payload` path
to reference the output of a Lambda executed before it.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke with empty object as payload"), &lambdaInvokeProps{
	lambdaFunction: fn,
	payload: sfn.taskInput.fromObject(map[string]interface{}{
	}),
})

// use the output of fn as input
// use the output of fn as input
tasks.NewLambdaInvoke(this, jsii.String("Invoke with payload field in the state input"), &lambdaInvokeProps{
	lambdaFunction: fn,
	payload: sfn.*taskInput.fromJsonPathAt(jsii.String("$.Payload")),
})
```

The following snippet invokes a Lambda and sets the task output to only include
the Lambda function response.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke and set function response as task output"), &lambdaInvokeProps{
	lambdaFunction: fn,
	outputPath: jsii.String("$.Payload"),
})
```

If you want to combine the input and the Lambda function response you can use
the `payloadResponseOnly` property and specify the `resultPath`. This will put the
Lambda function ARN directly in the "Resource" string, but it conflicts with the
integrationPattern, invocationType, clientContext, and qualifier properties.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke and combine function response with task input"), &lambdaInvokeProps{
	lambdaFunction: fn,
	payloadResponseOnly: jsii.Boolean(true),
	resultPath: jsii.String("$.fn"),
})
```

You can have Step Functions pause a task, and wait for an external process to
return a task token. Read more about the [callback pattern](https://docs.aws.amazon.com/step-functions/latest/dg/callback-task-sample-sqs.html#call-back-lambda-example)

To use the callback pattern, set the `token` property on the task. Call the Step
Functions `SendTaskSuccess` or `SendTaskFailure` APIs with the token to
indicate that the task has completed and the state machine should resume execution.

The following snippet invokes a Lambda with the task token as part of the input
to the Lambda.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke with callback"), &lambdaInvokeProps{
	lambdaFunction: fn,
	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
	payload: sfn.taskInput.fromObject(map[string]interface{}{
		"token": sfn.JsonPath.taskToken,
		"input": sfn.JsonPath.stringAt(jsii.String("$.someField")),
	}),
})
```

⚠️ The task will pause until it receives that task token back with a `SendTaskSuccess` or `SendTaskFailure`
call. Learn more about [Callback with the Task
Token](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token).

AWS Lambda can occasionally experience transient service errors. In this case, invoking Lambda
results in a 500 error, such as `ServiceException`, `AWSLambdaException`, or `SdkClientException`.
As a best practice, the `LambdaInvoke` task will retry on those errors with an interval of 2 seconds,
a back-off rate of 2 and 6 maximum attempts. Set the `retryOnServiceExceptions` prop to `false` to
disable this behavior.

## SageMaker

Step Functions supports [AWS SageMaker](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html) through the service integration pattern.

If your training job or model uses resources from AWS Marketplace,
[network isolation is required](https://docs.aws.amazon.com/sagemaker/latest/dg/mkt-algo-model-internet-free.html).
To do so, set the `enableNetworkIsolation` property to `true` for `SageMakerCreateModel` or `SageMakerCreateTrainingJob`.

To set environment variables for the Docker container use the `environment` property.

### Create Training Job

You can call the [`CreateTrainingJob`](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateTrainingJob.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &sageMakerCreateTrainingJobProps{
	trainingJobName: sfn.jsonPath.stringAt(jsii.String("$.JobName")),
	algorithmSpecification: &algorithmSpecification{
		algorithmName: jsii.String("BlazingText"),
		trainingInputMode: tasks.inputMode_FILE,
	},
	inputDataConfig: []channel{
		&channel{
			channelName: jsii.String("train"),
			dataSource: &dataSource{
				s3DataSource: &s3DataSource{
					s3DataType: tasks.s3DataType_S3_PREFIX,
					s3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.S3Bucket")),
				},
			},
		},
	},
	outputDataConfig: &outputDataConfig{
		s3OutputLocation: tasks.*s3Location.fromBucket(s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
	},
	resourceConfig: &resourceConfig{
		instanceCount: jsii.Number(1),
		instanceType: ec2.NewInstanceType(sfn.*jsonPath.stringAt(jsii.String("$.InstanceType"))),
		volumeSize: awscdk.Size.gibibytes(jsii.Number(50)),
	},
	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
	stoppingCondition: &stoppingCondition{
		maxRuntime: awscdk.Duration.hours(jsii.Number(2)),
	},
})
```

### Create Transform Job

You can call the [`CreateTransformJob`](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateTransformJob.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &sageMakerCreateTransformJobProps{
	transformJobName: jsii.String("MyTransformJob"),
	modelName: jsii.String("MyModelName"),
	modelClientOptions: &modelClientOptions{
		invocationsMaxRetries: jsii.Number(3),
		 // default is 0
		invocationsTimeout: awscdk.Duration.minutes(jsii.Number(5)),
	},
	transformInput: &transformInput{
		transformDataSource: &transformDataSource{
			s3DataSource: &transformS3DataSource{
				s3Uri: jsii.String("s3://inputbucket/train"),
				s3DataType: tasks.s3DataType_S3_PREFIX,
			},
		},
	},
	transformOutput: &transformOutput{
		s3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
	},
	transformResources: &transformResources{
		instanceCount: jsii.Number(1),
		instanceType: ec2.instanceType.of(ec2.instanceClass_M4, ec2.instanceSize_XLARGE),
	},
})
```

### Create Endpoint

You can call the [`CreateEndpoint`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateEndpoint(this, jsii.String("SagemakerEndpoint"), &sageMakerCreateEndpointProps{
	endpointName: sfn.jsonPath.stringAt(jsii.String("$.EndpointName")),
	endpointConfigName: sfn.*jsonPath.stringAt(jsii.String("$.EndpointConfigName")),
})
```

### Create Endpoint Config

You can call the [`CreateEndpointConfig`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateEndpointConfig(this, jsii.String("SagemakerEndpointConfig"), &sageMakerCreateEndpointConfigProps{
	endpointConfigName: jsii.String("MyEndpointConfig"),
	productionVariants: []productionVariant{
		&productionVariant{
			initialInstanceCount: jsii.Number(2),
			instanceType: ec2.instanceType.of(ec2.instanceClass_M5, ec2.instanceSize_XLARGE),
			modelName: jsii.String("MyModel"),
			variantName: jsii.String("awesome-variant"),
		},
	},
})
```

### Create Model

You can call the [`CreateModel`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateModel(this, jsii.String("Sagemaker"), &sageMakerCreateModelProps{
	modelName: jsii.String("MyModel"),
	primaryContainer: tasks.NewContainerDefinition(&containerDefinitionOptions{
		image: tasks.dockerImage.fromJsonExpression(sfn.jsonPath.stringAt(jsii.String("$.Model.imageName"))),
		mode: tasks.mode_SINGLE_MODEL,
		modelS3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.TrainingJob.ModelArtifacts.S3ModelArtifacts")),
	}),
})
```

### Update Endpoint

You can call the [`UpdateEndpoint`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpoint.html) API from a `Task` state.

```go
tasks.NewSageMakerUpdateEndpoint(this, jsii.String("SagemakerEndpoint"), &sageMakerUpdateEndpointProps{
	endpointName: sfn.jsonPath.stringAt(jsii.String("$.Endpoint.Name")),
	endpointConfigName: sfn.*jsonPath.stringAt(jsii.String("$.Endpoint.EndpointConfig")),
})
```

## SNS

Step Functions supports [Amazon SNS](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sns.html) through the service integration pattern.

You can call the [`Publish`](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) API from a `Task` state to publish to an SNS topic.

```go
topic := sns.NewTopic(this, jsii.String("Topic"))

// Use a field from the execution data as message.
task1 := tasks.NewSnsPublish(this, jsii.String("Publish1"), &snsPublishProps{
	topic: topic,
	integrationPattern: sfn.integrationPattern_REQUEST_RESPONSE,
	message: sfn.taskInput.fromDataAt(jsii.String("$.state.message")),
	messageAttributes: map[string]messageAttribute{
		"place": &messageAttribute{
			"value": sfn.JsonPath.stringAt(jsii.String("$.place")),
		},
		"pic": &messageAttribute{
			// BINARY must be explicitly set
			"dataType": tasks.MessageAttributeDataType_BINARY,
			"value": sfn.JsonPath.stringAt(jsii.String("$.pic")),
		},
		"people": &messageAttribute{
			"value": jsii.Number(4),
		},
		"handles": &messageAttribute{
			"value": []interface{}{
				jsii.String("@kslater"),
				jsii.String("@jjf"),
				nil,
				jsii.String("@mfanning"),
			},
		},
	},
})

// Combine a field from the execution data with
// a literal object.
task2 := tasks.NewSnsPublish(this, jsii.String("Publish2"), &snsPublishProps{
	topic: topic,
	message: sfn.*taskInput.fromObject(map[string]interface{}{
		"field1": jsii.String("somedata"),
		"field2": sfn.JsonPath.stringAt(jsii.String("$.field2")),
	}),
})
```

## Step Functions

### Start Execution

You can manage [AWS Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/connect-stepfunctions.html) executions.

AWS Step Functions supports it's own [`StartExecution`](https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html) API as a service integration.

```go
// Define a state machine with one Pass state
child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &stateMachineProps{
	definition: sfn.chain.start(sfn.NewPass(this, jsii.String("PassState"))),
})

// Include the state machine in a Task state with callback pattern
task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &stepFunctionsStartExecutionProps{
	stateMachine: child,
	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
	input: sfn.taskInput.fromObject(map[string]interface{}{
		"token": sfn.JsonPath.taskToken,
		"foo": jsii.String("bar"),
	}),
	name: jsii.String("MyExecutionName"),
})

// Define a second state machine with the Task state above
// Define a second state machine with the Task state above
sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &stateMachineProps{
	definition: task,
})
```

You can utilize [Associate Workflow Executions](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-nested-workflows.html#nested-execution-startid)
via the `associateWithParent` property. This allows the Step Functions UI to link child
executions from parent executions, making it easier to trace execution flow across state machines.

```go
var child stateMachine

task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &stepFunctionsStartExecutionProps{
	stateMachine: child,
	associateWithParent: jsii.Boolean(true),
})
```

This will add the payload `AWS_STEP_FUNCTIONS_STARTED_BY_EXECUTION_ID.$: $$.Execution.Id` to the
`input`property for you, which will pass the execution ID from the context object to the
execution input. It requires `input` to be an object or not be set at all.

### Invoke Activity

You can invoke a [Step Functions Activity](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html) which enables you to have
a task in your state machine where the work is performed by a *worker* that can
be hosted on Amazon EC2, Amazon ECS, AWS Lambda, basically anywhere. Activities
are a way to associate code running somewhere (known as an activity worker) with
a specific task in a state machine.

When Step Functions reaches an activity task state, the workflow waits for an
activity worker to poll for a task. An activity worker polls Step Functions by
using GetActivityTask, and sending the ARN for the related activity.

After the activity worker completes its work, it can provide a report of its
success or failure by using `SendTaskSuccess` or `SendTaskFailure`. These two
calls use the taskToken provided by GetActivityTask to associate the result
with that task.

The following example creates an activity and creates a task that invokes the activity.

```go
submitJobActivity := sfn.NewActivity(this, jsii.String("SubmitJob"))

tasks.NewStepFunctionsInvokeActivity(this, jsii.String("Submit Job"), &stepFunctionsInvokeActivityProps{
	activity: submitJobActivity,
})
```

Use the [Parameters](https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters) field to create a collection of key-value pairs that are passed as input.
The values of each can either be static values that you include in your state machine definition, or selected from either the input or the context object with a path.

```go
submitJobActivity := sfn.NewActivity(this, jsii.String("SubmitJob"))

tasks.NewStepFunctionsInvokeActivity(this, jsii.String("Submit Job"), &stepFunctionsInvokeActivityProps{
	activity: submitJobActivity,
	parameters: map[string]interface{}{
		"comment": jsii.String("Selecting what I care about."),
		"MyDetails": map[string]interface{}{
			"size": sfn.JsonPath.stringAt(jsii.String("$.product.details.size")),
			"exists": sfn.JsonPath.stringAt(jsii.String("$.product.availability")),
			"StaticValue": jsii.String("foo"),
		},
	},
})
```

## SQS

Step Functions supports [Amazon SQS](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sqs.html)

You can call the [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) API from a `Task` state
to send a message to an SQS queue.

```go
queue := sqs.NewQueue(this, jsii.String("Queue"))

// Use a field from the execution data as message.
task1 := tasks.NewSqsSendMessage(this, jsii.String("Send1"), &sqsSendMessageProps{
	queue: queue,
	messageBody: sfn.taskInput.fromJsonPathAt(jsii.String("$.message")),
})

// Combine a field from the execution data with
// a literal object.
task2 := tasks.NewSqsSendMessage(this, jsii.String("Send2"), &sqsSendMessageProps{
	queue: queue,
	messageBody: sfn.*taskInput.fromObject(map[string]interface{}{
		"field1": jsii.String("somedata"),
		"field2": sfn.JsonPath.stringAt(jsii.String("$.field2")),
	}),
})
```
