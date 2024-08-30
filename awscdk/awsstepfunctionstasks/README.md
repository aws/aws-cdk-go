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
  * [Paths](#paths)
  * [Evaluate Expression](#evaluate-expression)
  * [API Gateway](#api-gateway)

    * [Call REST API Endpoint](#call-rest-api-endpoint)
    * [Call HTTP API Endpoint](#call-http-api-endpoint)
  * [AWS SDK](#aws-sdk)

    * [Cross-region AWS API call](#cross-region-aws-api-call)
  * [Athena](#athena)

    * [StartQueryExecution](#startqueryexecution)
    * [GetQueryExecution](#getqueryexecution)
    * [GetQueryResults](#getqueryresults)
    * [StopQueryExecution](#stopqueryexecution)
  * [Batch](#batch)

    * [SubmitJob](#submitjob)
  * [Bedrock](#bedrock)

    * [InvokeModel](#invokemodel)
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

    * [Start Job Run](#start-job-run)
    * [Start Crawler Run](#startcrawlerrun)
  * [Glue DataBrew](#glue-databrew)

    * [Start Job Run](#start-job-run-1)
  * [Lambda](#lambda)

    * [Invoke](#invoke)
  * [MediaConvert](#mediaconvert)

    * [Create Job](#create-job)
  * [SageMaker](#sagemaker)

    * [Create Training Job](#create-training-job)
    * [Create Transform Job](#create-transform-job)
    * [Create Endpoint](#create-endpoint)
    * [Create Endpoint Config](#create-endpoint-config)
    * [Create Model](#create-model)
    * [Update Endpoint](#update-endpoint)
  * [SNS](#sns)

    * [Publish](#publish)
  * [Step Functions](#step-functions)

    * [Start Execution](#start-execution)
    * [Invoke Activity](#invoke-activity)
  * [SQS](#sqs)

    * [Send Message](#send-message)

## Paths

Learn more about input and output processing in Step Functions [here](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html)

## Evaluate Expression

Use the `EvaluateExpression` to perform simple operations referencing state paths. The
`expression` referenced in the task will be evaluated in a Lambda function
(`eval()`). This allows you to not have to write Lambda code for simple operations.

Example: convert a wait time from milliseconds to seconds, concat this in a message and wait:

```go
convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &EvaluateExpressionProps{
	Expression: jsii.String("$.waitMilliseconds / 1000"),
	ResultPath: jsii.String("$.waitSeconds"),
})

createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &EvaluateExpressionProps{
	// Note: this is a string inside a string.
	Expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	ResultPath: jsii.String("$.message"),
})

publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &SnsPublishProps{
	Topic: sns.NewTopic(this, jsii.String("cool-topic")),
	Message: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
	ResultPath: jsii.String("$.sns"),
})

wait := sfn.NewWait(this, jsii.String("Wait"), &WaitProps{
	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
})

sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	Definition: convertToSeconds.Next(createMessage).Next(publishMessage).*Next(wait),
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

invokeTask := tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Call REST API"), &CallApiGatewayRestApiEndpointProps{
	Api: restApi,
	StageName: jsii.String("prod"),
	Method: tasks.HttpMethod_GET,
})
```

Be aware that the header values must be arrays. When passing the Task Token
in the headers field `WAIT_FOR_TASK_TOKEN` integration, use
`JsonPath.array()` to wrap the token in an array:

```go
import apigateway "github.com/aws/aws-cdk-go/awscdk"
var api restApi


tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Endpoint"), &CallApiGatewayRestApiEndpointProps{
	Api: Api,
	StageName: jsii.String("Stage"),
	Method: tasks.HttpMethod_PUT,
	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
	Headers: sfn.TaskInput_FromObject(map[string]interface{}{
		"TaskToken": sfn.JsonPath_array(sfn.JsonPath_taskToken()),
	}),
})
```

### Call HTTP API Endpoint

The `CallApiGatewayHttpApiEndpoint` calls the HTTP API endpoint.

```go
import apigatewayv2 "github.com/aws/aws-cdk-go/awscdk"

httpApi := apigatewayv2.NewHttpApi(this, jsii.String("MyHttpApi"))

invokeTask := tasks.NewCallApiGatewayHttpApiEndpoint(this, jsii.String("Call HTTP API"), &CallApiGatewayHttpApiEndpointProps{
	ApiId: httpApi.ApiId,
	ApiStack: awscdk.*stack_Of(httpApi),
	Method: tasks.HttpMethod_GET,
})
```

## AWS SDK

Step Functions supports calling [AWS service's API actions](https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html)
through the service integration pattern.

You can use Step Functions' AWS SDK integrations to call any of the over two hundred AWS services
directly from your state machine, giving you access to over nine thousand API actions.

```go
var myBucket bucket

getObject := tasks.NewCallAwsService(this, jsii.String("GetObject"), &CallAwsServiceProps{
	Service: jsii.String("s3"),
	Action: jsii.String("getObject"),
	Parameters: map[string]interface{}{
		"Bucket": myBucket.bucketName,
		"Key": sfn.JsonPath_stringAt(jsii.String("$.key")),
	},
	IamResources: []*string{
		myBucket.ArnForObjects(jsii.String("*")),
	},
})
```

Use camelCase for actions and PascalCase for parameter names.

The task automatically adds an IAM statement to the state machine role's policy based on the
service and action called. The resources for this statement must be specified in `iamResources`.

Use the `iamAction` prop to manually specify the IAM action name in the case where the IAM
action name does not match with the API service/action name:

```go
listBuckets := tasks.NewCallAwsService(this, jsii.String("ListBuckets"), &CallAwsServiceProps{
	Service: jsii.String("s3"),
	Action: jsii.String("listBuckets"),
	IamResources: []*string{
		jsii.String("*"),
	},
	IamAction: jsii.String("s3:ListAllMyBuckets"),
})
```

Use the `additionalIamStatements` prop to pass additional IAM statements that will be added to the
state machine role's policy. Use it in the case where the call requires more than a single statement
to be executed:

```go
detectLabels := tasks.NewCallAwsService(this, jsii.String("DetectLabels"), &CallAwsServiceProps{
	Service: jsii.String("rekognition"),
	Action: jsii.String("detectLabels"),
	IamResources: []*string{
		jsii.String("*"),
	},
	AdditionalIamStatements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("s3:getObject"),
			},
			Resources: []*string{
				jsii.String("arn:aws:s3:::my-bucket/*"),
			},
		}),
	},
})
```

### Cross-region AWS API call

You can call AWS API in a different region from your state machine by using the `CallAwsServiceCrossRegion` construct. In addition to the properties for `CallAwsService` construct, you have to set `region` property to call the API.

```go
var myBucket bucket

getObject := tasks.NewCallAwsServiceCrossRegion(this, jsii.String("GetObject"), &CallAwsServiceCrossRegionProps{
	Region: jsii.String("ap-northeast-1"),
	Service: jsii.String("s3"),
	Action: jsii.String("getObject"),
	Parameters: map[string]interface{}{
		"Bucket": myBucket.bucketName,
		"Key": sfn.JsonPath_stringAt(jsii.String("$.key")),
	},
	IamResources: []*string{
		myBucket.ArnForObjects(jsii.String("*")),
	},
})
```

Other properties such as `additionalIamStatements` can be used in the same way as the `CallAwsService` task.

## Athena

Step Functions supports [Athena](https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html) through the service integration pattern.

### StartQueryExecution

The [StartQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_StartQueryExecution.html) API runs the SQL query statement.

```go
startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
	QueryExecutionContext: &QueryExecutionContext{
		DatabaseName: jsii.String("mydatabase"),
	},
	ResultConfiguration: &ResultConfiguration{
		EncryptionConfiguration: &EncryptionConfiguration{
			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
		},
		OutputLocation: &Location{
			BucketName: jsii.String("query-results-bucket"),
			ObjectKey: jsii.String("folder"),
		},
	},
	ExecutionParameters: []*string{
		jsii.String("param1"),
		jsii.String("param2"),
	},
})
```

You can reuse the query results by setting the `resultReuseConfigurationMaxAge` property.

```go
startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
	QueryExecutionContext: &QueryExecutionContext{
		DatabaseName: jsii.String("mydatabase"),
	},
	ResultConfiguration: &ResultConfiguration{
		EncryptionConfiguration: &EncryptionConfiguration{
			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
		},
		OutputLocation: &Location{
			BucketName: jsii.String("query-results-bucket"),
			ObjectKey: jsii.String("folder"),
		},
	},
	ExecutionParameters: []*string{
		jsii.String("param1"),
		jsii.String("param2"),
	},
	ResultReuseConfigurationMaxAge: awscdk.Duration_Minutes(jsii.Number(100)),
})
```

### GetQueryExecution

The [GetQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetQueryExecution.html) API gets information about a single execution of a query.

```go
getQueryExecutionJob := tasks.NewAthenaGetQueryExecution(this, jsii.String("Get Query Execution"), &AthenaGetQueryExecutionProps{
	QueryExecutionId: sfn.JsonPath_StringAt(jsii.String("$.QueryExecutionId")),
})
```

### GetQueryResults

The [GetQueryResults](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetQueryResults.html) API that streams the results of a single query execution specified by QueryExecutionId from S3.

```go
getQueryResultsJob := tasks.NewAthenaGetQueryResults(this, jsii.String("Get Query Results"), &AthenaGetQueryResultsProps{
	QueryExecutionId: sfn.JsonPath_StringAt(jsii.String("$.QueryExecutionId")),
})
```

### StopQueryExecution

The [StopQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_StopQueryExecution.html) API that stops a query execution.

```go
stopQueryExecutionJob := tasks.NewAthenaStopQueryExecution(this, jsii.String("Stop Query Execution"), &AthenaStopQueryExecutionProps{
	QueryExecutionId: sfn.JsonPath_StringAt(jsii.String("$.QueryExecutionId")),
})
```

## Batch

Step Functions supports [Batch](https://docs.aws.amazon.com/step-functions/latest/dg/connect-batch.html) through the service integration pattern.

### SubmitJob

The [SubmitJob](https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html) API submits an AWS Batch job from a job definition.

```go
import batch "github.com/aws/aws-cdk-go/awscdk"
var batchJobDefinition ecsJobDefinition
var batchQueue jobQueue


task := tasks.NewBatchSubmitJob(this, jsii.String("Submit Job"), &BatchSubmitJobProps{
	JobDefinitionArn: batchJobDefinition.JobDefinitionArn,
	JobName: jsii.String("MyJob"),
	JobQueueArn: batchQueue.JobQueueArn,
})
```

## Bedrock

Step Functions supports [Bedrock](https://docs.aws.amazon.com/step-functions/latest/dg/connect-bedrock.html) through the service integration pattern.

### InvokeModel

The [InvokeModel](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModel.html) API
invokes the specified Bedrock model to run inference using the input provided.
The format of the input body and the response body depend on the model selected.

```go
import "github.com/aws/aws-cdk-go/awscdk"


model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())

task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model"), &BedrockInvokeModelProps{
	Model: Model,
	Body: sfn.TaskInput_FromObject(map[string]interface{}{
		"inputText": jsii.String("Generate a list of five first names."),
		"textGenerationConfig": map[string]*f64{
			"maxTokenCount": jsii.Number(100),
			"temperature": jsii.Number(1),
		},
	}),
	ResultSelector: map[string]interface{}{
		"names": sfn.JsonPath_stringAt(jsii.String("$.Body.results[0].outputText")),
	},
})
```

### Using Input Path

Provide S3 URI as an input or output path to invoke a model

```go
import "github.com/aws/aws-cdk-go/awscdk"


model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())

task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model"), &BedrockInvokeModelProps{
	Model: Model,
	InputPath: sfn.JsonPath_StringAt(jsii.String("$.prompt")),
	OutputPath: sfn.JsonPath_*StringAt(jsii.String("$.prompt")),
})
```

You can apply a guardrail to the invocation by setting `guardrail`.

```go
import "github.com/aws/aws-cdk-go/awscdk"


model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())

task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model with guardrail"), &BedrockInvokeModelProps{
	Model: Model,
	Body: sfn.TaskInput_FromObject(map[string]interface{}{
		"inputText": jsii.String("Generate a list of five first names."),
		"textGenerationConfig": map[string]*f64{
			"maxTokenCount": jsii.Number(100),
			"temperature": jsii.Number(1),
		},
	}),
	Guardrail: tasks.Guardrail_Enable(jsii.String("guardrailId"), jsii.Number(1)),
	ResultSelector: map[string]interface{}{
		"names": sfn.JsonPath_stringAt(jsii.String("$.Body.results[0].outputText")),
	},
})
```

## CodeBuild

Step Functions supports [CodeBuild](https://docs.aws.amazon.com/step-functions/latest/dg/connect-codebuild.html) through the service integration pattern.

### StartBuild

[StartBuild](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuild.html) starts a CodeBuild Project by Project Name.

```go
import "github.com/aws/aws-cdk-go/awscdk"


codebuildProject := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	ProjectName: jsii.String("MyTestProject"),
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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

task := tasks.NewCodeBuildStartBuild(this, jsii.String("Task"), &CodeBuildStartBuildProps{
	Project: codebuildProject,
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
	EnvironmentVariablesOverride: map[string]buildEnvironmentVariable{
		"ZONE": &buildEnvironmentVariable{
			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
			"value": sfn.JsonPath_stringAt(jsii.String("$.envVariables.zone")),
		},
	},
})
```

### StartBuildBatch

[StartBuildBatch](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuildBatch.html) starts a batch CodeBuild for a project by project name.
It is necessary to enable the batch build feature in the CodeBuild project.

```go
import "github.com/aws/aws-cdk-go/awscdk"


project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	ProjectName: jsii.String("MyTestProject"),
	BuildSpec: codebuild.BuildSpec_FromObjectToYaml(map[string]interface{}{
		"version": jsii.Number(0.2),
		"batch": map[string][]map[string]*string{
			"build-list": []map[string]*string{
				map[string]*string{
					"identifier": jsii.String("id"),
					"buildspec": jsii.String("version: 0.2\nphases:\n  build:\n    commands:\n      - echo \"Hello, from small!\""),
				},
			},
		},
	}),
})
project.EnableBatchBuilds()

task := tasks.NewCodeBuildStartBuildBatch(this, jsii.String("buildBatchTask"), &CodeBuildStartBuildBatchProps{
	Project: Project,
	IntegrationPattern: sfn.IntegrationPattern_REQUEST_RESPONSE,
	EnvironmentVariablesOverride: map[string]buildEnvironmentVariable{
		"test": &buildEnvironmentVariable{
			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
			"value": jsii.String("testValue"),
		},
	},
})
```

**Note**: `enableBatchBuilds()` will do nothing for imported projects.
If you are using an imported project, you must ensure that the project is already configured for batch builds.

## DynamoDB

You can call DynamoDB APIs from a `Task` state.
Read more about calling DynamoDB APIs [here](https://docs.aws.amazon.com/step-functions/latest/dg/connect-ddb.html)

### GetItem

The [GetItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html) operation returns a set of attributes for the item with the given primary key.

```go
var myTable table

tasks.NewDynamoGetItem(this, jsii.String("Get Item"), &DynamoGetItemProps{
	Key: map[string]dynamoAttributeValue{
		"messageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
	},
	Table: myTable,
})
```

### PutItem

The [PutItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html) operation creates a new item, or replaces an old item with a new item.

```go
var myTable table

tasks.NewDynamoPutItem(this, jsii.String("PutItem"), &DynamoPutItemProps{
	Item: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
		"Text": tasks.*dynamoAttributeValue_fromString(sfn.JsonPath_stringAt(jsii.String("$.bar"))),
		"TotalCount": tasks.*dynamoAttributeValue_fromNumber(jsii.Number(10)),
	},
	Table: myTable,
})
```

### DeleteItem

The [DeleteItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html) operation deletes a single item in a table by primary key.

```go
var myTable table

tasks.NewDynamoDeleteItem(this, jsii.String("DeleteItem"), &DynamoDeleteItemProps{
	Key: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
	},
	Table: myTable,
	ResultPath: sfn.JsonPath_DISCARD(),
})
```

### UpdateItem

The [UpdateItem](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html) operation edits an existing item's attributes, or adds a new item
to the table if it does not already exist.

```go
var myTable table

tasks.NewDynamoUpdateItem(this, jsii.String("UpdateItem"), &DynamoUpdateItemProps{
	Key: map[string]dynamoAttributeValue{
		"MessageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
	},
	Table: myTable,
	ExpressionAttributeValues: map[string]*dynamoAttributeValue{
		":val": tasks.*dynamoAttributeValue_numberFromString(sfn.JsonPath_stringAt(jsii.String("$.Item.TotalCount.N"))),
		":rand": tasks.*dynamoAttributeValue_fromNumber(jsii.Number(20)),
	},
	UpdateExpression: jsii.String("SET TotalCount = :val + :rand"),
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
vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})

cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &ClusterProps{
	Vpc: Vpc,
})
cluster.AddCapacity(jsii.String("DefaultAutoScalingGroup"), &AddCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
	Compatibility: ecs.Compatibility_EC2,
})

taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
	MemoryLimitMiB: jsii.Number(256),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	LaunchTarget: tasks.NewEcsEc2LaunchTarget(&EcsEc2LaunchTargetOptions{
		PlacementStrategies: []placementStrategy{
			ecs.*placementStrategy_SpreadAcrossInstances(),
			ecs.*placementStrategy_PackedByCpu(),
			ecs.*placementStrategy_Randomly(),
		},
		PlacementConstraints: []placementConstraint{
			ecs.*placementConstraint_MemberOf(jsii.String("blieptuut")),
		},
	}),
	PropagatedTagSource: ecs.PropagatedTagSource_TASK_DEFINITION,
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
vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})

cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &ClusterProps{
	Vpc: Vpc,
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
	MemoryMiB: jsii.String("512"),
	Cpu: jsii.String("256"),
	Compatibility: ecs.Compatibility_FARGATE,
})

containerDefinition := taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
	MemoryLimitMiB: jsii.Number(256),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &EcsRunTaskProps{
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	AssignPublicIp: jsii.Boolean(true),
	ContainerOverrides: []containerOverride{
		&containerOverride{
			ContainerDefinition: *ContainerDefinition,
			Environment: []taskEnvironmentVariable{
				&taskEnvironmentVariable{
					Name: jsii.String("SOME_KEY"),
					Value: sfn.JsonPath_StringAt(jsii.String("$.SomeKey")),
				},
			},
		},
	},
	LaunchTarget: tasks.NewEcsFargateLaunchTarget(),
	PropagatedTagSource: ecs.PropagatedTagSource_TASK_DEFINITION,
})
```

#### Override CPU and Memory Parameter

By setting the property cpu or memoryMiB, you can override the Fargate or EC2 task instance size at runtime.

see: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskOverride.html

```go
vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})
cluster := ecs.NewCluster(this, jsii.String("ECSCluster"), &ClusterProps{
	Vpc: Vpc,
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
	Compatibility: ecs.Compatibility_FARGATE,
	Cpu: jsii.String("256"),
	MemoryMiB: jsii.String("512"),
})

taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	LaunchTarget: tasks.NewEcsFargateLaunchTarget(),
	Cpu: jsii.String("1024"),
	MemoryMiB: jsii.String("1048"),
})
```

#### ECS enable Exec

By setting the property [`enableExecuteCommand`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-enableExecuteCommand) to `true`, you can enable the [ECS Exec feature](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-exec.html) for the task for either Fargate or EC2 launch types.

```go
vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})
cluster := ecs.NewCluster(this, jsii.String("ECSCluster"), &ClusterProps{
	Vpc: Vpc,
})

taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
	Compatibility: ecs.Compatibility_EC2,
})

taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
	MemoryLimitMiB: jsii.Number(256),
})

runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	LaunchTarget: tasks.NewEcsEc2LaunchTarget(),
	EnableExecuteCommand: jsii.Boolean(true),
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
clusterRole := iam.NewRole(this, jsii.String("ClusterRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
})

serviceRole := iam.NewRole(this, jsii.String("ServiceRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
})

autoScalingRole := iam.NewRole(this, jsii.String("AutoScalingRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
})

autoScalingRole.AssumeRolePolicy.AddStatements(
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Principals: []iPrincipal{
		iam.NewServicePrincipal(jsii.String("application-autoscaling.amazonaws.com")),
	},
	Actions: []*string{
		jsii.String("sts:AssumeRole"),
	},
}))

tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &EmrCreateClusterProps{
	Instances: &InstancesConfigProperty{
	},
	ClusterRole: ClusterRole,
	Name: sfn.TaskInput_FromJsonPathAt(jsii.String("$.ClusterName")).value,
	ServiceRole: ServiceRole,
	AutoScalingRole: AutoScalingRole,
})
```

You can use the launch specification for On-Demand and Spot instances in the fleet.

```go
tasks.NewEmrCreateCluster(this, jsii.String("OnDemandSpecification"), &EmrCreateClusterProps{
	Instances: &InstancesConfigProperty{
		InstanceFleets: []instanceFleetConfigProperty{
			&instanceFleetConfigProperty{
				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
					OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
						AllocationStrategy: tasks.EmrCreateCluster.OnDemandAllocationStrategy_LOWEST_PRICE,
					},
				},
			},
		},
	},
	Name: jsii.String("OnDemandCluster"),
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
})

tasks.NewEmrCreateCluster(this, jsii.String("SpotSpecification"), &EmrCreateClusterProps{
	Instances: &InstancesConfigProperty{
		InstanceFleets: []*instanceFleetConfigProperty{
			&instanceFleetConfigProperty{
				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
					SpotSpecification: &SpotProvisioningSpecificationProperty{
						AllocationStrategy: tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
						TimeoutAction: tasks.EmrCreateCluster.SpotTimeoutAction_TERMINATE_CLUSTER,
						Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
					},
				},
			},
		},
	},
	Name: jsii.String("SpotCluster"),
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
})
```

If you want to run multiple steps in [parallel](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-concurrent-steps.html),
you can specify the `stepConcurrencyLevel` property. The concurrency range is between 1
and 256 inclusive, where the default concurrency of 1 means no step concurrency is allowed.
`stepConcurrencyLevel` requires the EMR release label to be 5.28.0 or above.

```go
tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &EmrCreateClusterProps{
	Instances: &InstancesConfigProperty{
	},
	Name: sfn.TaskInput_FromJsonPathAt(jsii.String("$.ClusterName")).value,
	StepConcurrencyLevel: jsii.Number(10),
})
```

### Termination Protection

Locks a cluster (job flow) so the EC2 instances in the cluster cannot be
terminated by user intervention, an API call, or a job-flow error.

Corresponds to the [`setTerminationProtection`](https://docs.aws.amazon.com/step-functions/latest/dg/connect-emr.html) API in EMR.

```go
tasks.NewEmrSetClusterTerminationProtection(this, jsii.String("Task"), &EmrSetClusterTerminationProtectionProps{
	ClusterId: jsii.String("ClusterId"),
	TerminationProtected: jsii.Boolean(false),
})
```

### Terminate Cluster

Shuts down a cluster (job flow).
Corresponds to the [`terminateJobFlows`](https://docs.aws.amazon.com/emr/latest/APIReference/API_TerminateJobFlows.html) API in EMR.

```go
tasks.NewEmrTerminateCluster(this, jsii.String("Task"), &EmrTerminateClusterProps{
	ClusterId: jsii.String("ClusterId"),
})
```

### Add Step

Adds a new step to a running cluster.
Corresponds to the [`addJobFlowSteps`](https://docs.aws.amazon.com/emr/latest/APIReference/API_AddJobFlowSteps.html) API in EMR.

```go
tasks.NewEmrAddStep(this, jsii.String("Task"), &EmrAddStepProps{
	ClusterId: jsii.String("ClusterId"),
	Name: jsii.String("StepName"),
	Jar: jsii.String("Jar"),
	ActionOnFailure: tasks.ActionOnFailure_CONTINUE,
})
```

To specify a custom runtime role use the `executionRoleArn` property.

**Note:** The EMR cluster must be created with a security configuration and the runtime role must have a specific trust policy.
See this [blog post](https://aws.amazon.com/blogs/big-data/introducing-runtime-roles-for-amazon-emr-steps-use-iam-roles-and-aws-lake-formation-for-access-control-with-amazon-emr/) for more details.

```go
import emr "github.com/aws/aws-cdk-go/awscdk"


cfnSecurityConfiguration := emr.NewCfnSecurityConfiguration(this, jsii.String("EmrSecurityConfiguration"), &CfnSecurityConfigurationProps{
	Name: jsii.String("AddStepRuntimeRoleSecConfig"),
	SecurityConfiguration: jSON.parse(jsii.String(`
	    {
	      "AuthorizationConfiguration": {
	          "IAMConfiguration": {
	              "EnableApplicationScopedIAMRole": true,
	              "ApplicationScopedIAMRoleConfiguration":
	                  {
	                      "PropagateSourceIdentity": true
	                  }
	          },
	          "LakeFormationConfiguration": {
	              "AuthorizedSessionTagValue": "Amazon EMR"
	          }
	      }
	    }`)),
})

task := tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &EmrCreateClusterProps{
	Instances: &InstancesConfigProperty{
	},
	Name: sfn.TaskInput_FromJsonPathAt(jsii.String("$.ClusterName")).value,
	SecurityConfiguration: cfnSecurityConfiguration.Name,
})

executionRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewArnPrincipal(task.clusterRole.RoleArn),
})

executionRole.AssumeRolePolicy.AddStatements(
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Principals: []iPrincipal{
		task.clusterRole,
	},
	Actions: []*string{
		jsii.String("sts:SetSourceIdentity"),
	},
}),
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Principals: []*iPrincipal{
		task.clusterRole,
	},
	Actions: []*string{
		jsii.String("sts:TagSession"),
	},
	Conditions: map[string]interface{}{
		"StringEquals": map[string]*string{
			"aws:RequestTag/LakeFormationAuthorizedCaller": jsii.String("Amazon EMR"),
		},
	},
}))

tasks.NewEmrAddStep(this, jsii.String("Task"), &EmrAddStepProps{
	ClusterId: jsii.String("ClusterId"),
	ExecutionRoleArn: executionRole.RoleArn,
	Name: jsii.String("StepName"),
	Jar: jsii.String("Jar"),
	ActionOnFailure: tasks.ActionOnFailure_CONTINUE,
})
```

### Cancel Step

Cancels a pending step in a running cluster.
Corresponds to the [`cancelSteps`](https://docs.aws.amazon.com/emr/latest/APIReference/API_CancelSteps.html) API in EMR.

```go
tasks.NewEmrCancelStep(this, jsii.String("Task"), &EmrCancelStepProps{
	ClusterId: jsii.String("ClusterId"),
	StepId: jsii.String("StepId"),
})
```

### Modify Instance Fleet

Modifies the target On-Demand and target Spot capacities for the instance
fleet with the specified InstanceFleetName.

Corresponds to the [`modifyInstanceFleet`](https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceFleet.html) API in EMR.

```go
tasks.NewEmrModifyInstanceFleetByName(this, jsii.String("Task"), &EmrModifyInstanceFleetByNameProps{
	ClusterId: jsii.String("ClusterId"),
	InstanceFleetName: jsii.String("InstanceFleetName"),
	TargetOnDemandCapacity: jsii.Number(2),
	TargetSpotCapacity: jsii.Number(0),
})
```

### Modify Instance Group

Modifies the number of nodes and configuration settings of an instance group.

Corresponds to the [`modifyInstanceGroups`](https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceGroups.html) API in EMR.

```go
tasks.NewEmrModifyInstanceGroupByName(this, jsii.String("Task"), &EmrModifyInstanceGroupByNameProps{
	ClusterId: jsii.String("ClusterId"),
	InstanceGroupName: sfn.JsonPath_StringAt(jsii.String("$.InstanceGroupName")),
	InstanceGroup: &InstanceGroupModifyConfigProperty{
		InstanceCount: jsii.Number(1),
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
tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &EmrContainersCreateVirtualClusterProps{
	EksCluster: tasks.EksClusterInput_FromTaskInput(sfn.TaskInput_FromText(jsii.String("clusterId"))),
})
```

The EKS cluster can also be passed in directly.

```go
import eks "github.com/aws/aws-cdk-go/awscdk"

var eksCluster cluster


tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &EmrContainersCreateVirtualClusterProps{
	EksCluster: tasks.EksClusterInput_FromCluster(eksCluster),
})
```

By default, the Kubernetes namespace that a virtual cluster maps to is "default", but a specific namespace within an EKS cluster can be selected.

```go
tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &EmrContainersCreateVirtualClusterProps{
	EksCluster: tasks.EksClusterInput_FromTaskInput(sfn.TaskInput_FromText(jsii.String("clusterId"))),
	EksNamespace: jsii.String("specified-namespace"),
})
```

### Delete Virtual Cluster

The [DeleteVirtualCluster](https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_DeleteVirtualCluster.html) API deletes a virtual cluster.

```go
tasks.NewEmrContainersDeleteVirtualCluster(this, jsii.String("Delete a Virtual Cluster"), &EmrContainersDeleteVirtualClusterProps{
	VirtualClusterId: sfn.TaskInput_FromJsonPathAt(jsii.String("$.virtualCluster")),
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
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
	JobDriver: &JobDriver{
		SparkSubmitJobDriver: &SparkSubmitJobDriver{
			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			SparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
})
```

Configuring the job can also be done via application configuration:

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
	JobName: jsii.String("EMR-Containers-Job"),
	JobDriver: &JobDriver{
		SparkSubmitJobDriver: &SparkSubmitJobDriver{
			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
		},
	},
	ApplicationConfig: []applicationConfiguration{
		&applicationConfiguration{
			Classification: tasks.Classification_SPARK_DEFAULTS(),
			Properties: map[string]*string{
				"spark.executor.instances": jsii.String("1"),
				"spark.executor.memory": jsii.String("512M"),
			},
		},
	},
})
```

Job monitoring can be enabled if `monitoring.logging` is set true. This automatically generates an S3 bucket and CloudWatch logs.

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
	JobDriver: &JobDriver{
		SparkSubmitJobDriver: &SparkSubmitJobDriver{
			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			SparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
	Monitoring: &Monitoring{
		Logging: jsii.Boolean(true),
	},
})
```

Otherwise, providing monitoring for jobs with existing log groups and log buckets is also available.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
logBucket := s3.NewBucket(this, jsii.String("S3 Bucket"))

tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
	JobDriver: &JobDriver{
		SparkSubmitJobDriver: &SparkSubmitJobDriver{
			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			SparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
		},
	},
	Monitoring: &Monitoring{
		LogGroup: logGroup,
		LogBucket: logBucket,
	},
})
```

Users can provide their own existing Job Execution Role.

```go
tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
	VirtualCluster: tasks.VirtualClusterInput_FromTaskInput(sfn.TaskInput_FromJsonPathAt(jsii.String("$.VirtualClusterId"))),
	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
	JobName: jsii.String("EMR-Containers-Job"),
	ExecutionRole: iam.Role_FromRoleArn(this, jsii.String("Job-Execution-Role"), jsii.String("arn:aws:iam::xxxxxxxxxxxx:role/JobExecutionRole")),
	JobDriver: &JobDriver{
		SparkSubmitJobDriver: &SparkSubmitJobDriver{
			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
			SparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
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
import "github.com/aws/aws-cdk-go/awscdk"


myEksCluster := eks.NewCluster(this, jsii.String("my sample cluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_18(),
	ClusterName: jsii.String("myEksCluster"),
})

tasks.NewEksCall(this, jsii.String("Call a EKS Endpoint"), &EksCallProps{
	Cluster: myEksCluster,
	HttpMethod: tasks.HttpMethods_GET,
	HttpPath: jsii.String("/api/v1/namespaces/default/pods"),
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


myEventBus := events.NewEventBus(this, jsii.String("EventBus"), &EventBusProps{
	EventBusName: jsii.String("MyEventBus1"),
})

tasks.NewEventBridgePutEvents(this, jsii.String("Send an event to EventBridge"), &EventBridgePutEventsProps{
	Entries: []eventBridgePutEventsEntry{
		&eventBridgePutEventsEntry{
			Detail: sfn.TaskInput_FromObject(map[string]interface{}{
				"Message": jsii.String("Hello from Step Functions!"),
			}),
			EventBus: myEventBus,
			DetailType: jsii.String("MessageFromStepFunctions"),
			Source: jsii.String("step.functions"),
		},
	},
})
```

## Glue

Step Functions supports [AWS Glue](https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html) through the service integration pattern.

### StartJobRun

You can call the [`StartJobRun`](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-StartJobRun) API from a `Task` state.

```go
tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
	GlueJobName: jsii.String("my-glue-job"),
	Arguments: sfn.TaskInput_FromObject(map[string]interface{}{
		"key": jsii.String("value"),
	}),
	TaskTimeout: sfn.Timeout_Duration(awscdk.Duration_Minutes(jsii.Number(30))),
	NotifyDelayAfter: awscdk.Duration_*Minutes(jsii.Number(5)),
})
```

You can configure workers by setting the `workerType` and `numberOfWorkers` properties.

```go
tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
	GlueJobName: jsii.String("my-glue-job"),
	WorkerConfiguration: &WorkerConfigurationProperty{
		WorkerType: tasks.WorkerType_G_1X,
		 // Worker type
		NumberOfWorkers: jsii.Number(2),
	},
})
```

You can choose the execution class by setting the `executionClass` property.

```go
tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
	GlueJobName: jsii.String("my-glue-job"),
	ExecutionClass: tasks.ExecutionClass_FLEX,
})
```

### StartCrawlerRun

You can call the [`StartCrawler`](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-crawling.html#aws-glue-api-crawler-crawling-StartCrawler) API from a `Task` state through AWS SDK service integrations.

```go
import glue "github.com/aws/aws-cdk-go/awscdk"

var myCrawler cfnCrawler


// You can get the crawler name from `crawler.ref`
// You can get the crawler name from `crawler.ref`
tasks.NewGlueStartCrawlerRun(this, jsii.String("Task1"), &GlueStartCrawlerRunProps{
	CrawlerName: myCrawler.ref,
})

// Of course, you can also specify the crawler name directly.
// Of course, you can also specify the crawler name directly.
tasks.NewGlueStartCrawlerRun(this, jsii.String("Task2"), &GlueStartCrawlerRunProps{
	CrawlerName: jsii.String("my-crawler-job"),
})
```

## Glue DataBrew

Step Functions supports [AWS Glue DataBrew](https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html) through the service integration pattern.

### Start Job Run

You can call the [`StartJobRun`](https://docs.aws.amazon.com/databrew/latest/dg/API_StartJobRun.html) API from a `Task` state.

```go
tasks.NewGlueDataBrewStartJobRun(this, jsii.String("Task"), &GlueDataBrewStartJobRunProps{
	Name: jsii.String("databrew-job"),
})
```

## Invoke HTTP API

Step Functions supports [calling third-party APIs](https://docs.aws.amazon.com/step-functions/latest/dg/connect-third-party-apis.html) with credentials managed by Amazon EventBridge [Connections](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Connection.html).

The following snippet creates a new API destination connection, and uses it to make a POST request to the specified URL. The endpoint response is available at the `$.ResponseBody` path.

```go
import "github.com/aws/aws-cdk-go/awscdk"


connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
	Authorization: events.Authorization_Basic(jsii.String("username"), awscdk.SecretValue_UnsafePlainText(jsii.String("password"))),
})

tasks.NewHttpInvoke(this, jsii.String("Invoke HTTP API"), &HttpInvokeProps{
	ApiRoot: jsii.String("https://api.example.com"),
	ApiEndpoint: sfn.TaskInput_FromText(jsii.String("https://api.example.com/path/to/resource")),
	Body: sfn.TaskInput_FromObject(map[string]interface{}{
		"foo": jsii.String("bar"),
	}),
	Connection: Connection,
	Headers: sfn.TaskInput_*FromObject(map[string]interface{}{
		"Content-Type": jsii.String("application/json"),
	}),
	Method: sfn.TaskInput_*FromText(jsii.String("POST")),
	QueryStringParameters: sfn.TaskInput_*FromObject(map[string]interface{}{
		"id": jsii.String("123"),
	}),
	UrlEncodingFormat: tasks.URLEncodingFormat_BRACKETS,
})
```

## Lambda

Step Functions supports [AWS Lambda](https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html) through the service integration pattern.

### Invoke

[Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html) a Lambda function.

You can specify the input to your Lambda function through the `payload` attribute.
By default, Step Functions invokes Lambda function with the state input (JSON path '$')
as the input.

The following snippet invokes a Lambda Function with the state input as the payload
by referencing the `$` path.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke with state input"), &LambdaInvokeProps{
	LambdaFunction: fn,
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

tasks.NewLambdaInvoke(this, jsii.String("Invoke with empty object as payload"), &LambdaInvokeProps{
	LambdaFunction: fn,
	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
	}),
})

// use the output of fn as input
// use the output of fn as input
tasks.NewLambdaInvoke(this, jsii.String("Invoke with payload field in the state input"), &LambdaInvokeProps{
	LambdaFunction: fn,
	Payload: sfn.TaskInput_FromJsonPathAt(jsii.String("$.Payload")),
})
```

The following snippet invokes a Lambda and sets the task output to only include
the Lambda function response.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke and set function response as task output"), &LambdaInvokeProps{
	LambdaFunction: fn,
	OutputPath: jsii.String("$.Payload"),
})
```

If you want to combine the input and the Lambda function response you can use
the `payloadResponseOnly` property and specify the `resultPath`. This will put the
Lambda function ARN directly in the "Resource" string, but it conflicts with the
integrationPattern, invocationType, clientContext, and qualifier properties.

```go
var fn function

tasks.NewLambdaInvoke(this, jsii.String("Invoke and combine function response with task input"), &LambdaInvokeProps{
	LambdaFunction: fn,
	PayloadResponseOnly: jsii.Boolean(true),
	ResultPath: jsii.String("$.fn"),
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

tasks.NewLambdaInvoke(this, jsii.String("Invoke with callback"), &LambdaInvokeProps{
	LambdaFunction: fn,
	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
		"token": sfn.JsonPath_taskToken(),
		"input": sfn.JsonPath_stringAt(jsii.String("$.someField")),
	}),
})
```

⚠️ The task will pause until it receives that task token back with a `SendTaskSuccess` or `SendTaskFailure`
call. Learn more about [Callback with the Task
Token](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token).

AWS Lambda can occasionally experience transient service errors. In this case, invoking Lambda
results in a 500 error, such as `ClientExecutionTimeoutException`, `ServiceException`, `AWSLambdaException`, or `SdkClientException`.
As a best practice, the `LambdaInvoke` task will retry on those errors with an interval of 2 seconds,
a back-off rate of 2 and 6 maximum attempts. Set the `retryOnServiceExceptions` prop to `false` to
disable this behavior.

## MediaConvert

Step Functions supports [AWS MediaConvert](https://docs.aws.amazon.com/step-functions/latest/dg/connect-mediaconvert.html) through the Optimized integration pattern.

### CreateJob

The [CreateJob](https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobspost) API creates a new transcoding job.
For information about jobs and job settings, see the User Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html

You can call the `CreateJob` API from a `Task` state. Optionally you can specify the `integrationPattern`.

Make sure you update the required fields - [Role](https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobs-prop-createjobrequest-role) &
[Settings](https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobs-prop-createjobrequest-settings) and refer
[CreateJobRequest](https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobs-model-createjobrequest) for all other optional parameters.

```go
tasks.NewMediaConvertCreateJob(this, jsii.String("CreateJob"), &MediaConvertCreateJobProps{
	CreateJobRequest: map[string]interface{}{
		"Role": jsii.String("arn:aws:iam::123456789012:role/MediaConvertRole"),
		"Settings": map[string][]map[string]interface{}{
			"OutputGroups": []map[string]interface{}{
				map[string]interface{}{
					"Outputs": []map[string]interface{}{
						map[string]interface{}{
							"ContainerSettings": map[string]*string{
								"Container": jsii.String("MP4"),
							},
							"VideoDescription": map[string]map[string]interface{}{
								"CodecSettings": map[string]interface{}{
									"Codec": jsii.String("H_264"),
									"H264Settings": map[string]interface{}{
										"MaxBitrate": jsii.Number(1000),
										"RateControlMode": jsii.String("QVBR"),
										"SceneChangeDetect": jsii.String("TRANSITION_DETECTION"),
									},
								},
							},
							"AudioDescriptions": []map[string]map[string]interface{}{
								map[string]map[string]interface{}{
									"CodecSettings": map[string]interface{}{
										"Codec": jsii.String("AAC"),
										"AacSettings": map[string]interface{}{
											"Bitrate": jsii.Number(96000),
											"CodingMode": jsii.String("CODING_MODE_2_0"),
											"SampleRate": jsii.Number(48000),
										},
									},
								},
							},
						},
					},
					"OutputGroupSettings": map[string]interface{}{
						"Type": jsii.String("FILE_GROUP_SETTINGS"),
						"FileGroupSettings": map[string]*string{
							"Destination": jsii.String("s3://EXAMPLE-DESTINATION-BUCKET/"),
						},
					},
				},
			},
			"Inputs": []map[string]interface{}{
				map[string]interface{}{
					"AudioSelectors": map[string]map[string]*string{
						"Audio Selector 1": map[string]*string{
							"DefaultSelection": jsii.String("DEFAULT"),
						},
					},
					"FileInput": jsii.String("s3://EXAMPLE-SOURCE-BUCKET/EXAMPLE-SOURCE_FILE"),
				},
			},
		},
	},
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
})
```

## SageMaker

Step Functions supports [AWS SageMaker](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html) through the service integration pattern.

If your training job or model uses resources from AWS Marketplace,
[network isolation is required](https://docs.aws.amazon.com/sagemaker/latest/dg/mkt-algo-model-internet-free.html).
To do so, set the `enableNetworkIsolation` property to `true` for `SageMakerCreateModel` or `SageMakerCreateTrainingJob`.

To set environment variables for the Docker container use the `environment` property.

### Create Training Job

You can call the [`CreateTrainingJob`](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateTrainingJob.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
	AlgorithmSpecification: &AlgorithmSpecification{
		AlgorithmName: jsii.String("BlazingText"),
		TrainingInputMode: tasks.InputMode_FILE,
	},
	InputDataConfig: []channel{
		&channel{
			ChannelName: jsii.String("train"),
			DataSource: &DataSource{
				S3DataSource: &S3DataSource{
					S3DataType: tasks.S3DataType_S3_PREFIX,
					S3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.S3Bucket")),
				},
			},
		},
	},
	OutputDataConfig: &OutputDataConfig{
		S3OutputLocation: tasks.S3Location_FromBucket(s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
	},
	ResourceConfig: &ResourceConfig{
		InstanceCount: jsii.Number(1),
		InstanceType: ec2.NewInstanceType(sfn.JsonPath_*StringAt(jsii.String("$.InstanceType"))),
		VolumeSize: awscdk.Size_Gibibytes(jsii.Number(50)),
	},
	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
	StoppingCondition: &StoppingCondition{
		MaxRuntime: awscdk.Duration_Hours(jsii.Number(2)),
	},
})
```

You can specify [TrainingInputMode](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html#API_AlgorithmSpecification_Contents) via the trainingInputMode property.

* To download the data from Amazon Simple Storage Service (Amazon S3) to the provisioned ML storage volume, and mount the directory to a Docker volume, choose `InputMode.FILE` if an algorithm supports it.
* To stream data directly from Amazon S3 to the container, choose `InputMode.PIPE` if an algorithm supports it.
* To stream data directly from Amazon S3 to the container with no code changes and to provide file system access to the data, choose `InputMode.FAST_FILE` if an algorithm supports it.

### Create Transform Job

You can call the [`CreateTransformJob`](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateTransformJob.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &SageMakerCreateTransformJobProps{
	TransformJobName: jsii.String("MyTransformJob"),
	ModelName: jsii.String("MyModelName"),
	ModelClientOptions: &ModelClientOptions{
		InvocationsMaxRetries: jsii.Number(3),
		 // default is 0
		InvocationsTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
	},
	TransformInput: &TransformInput{
		TransformDataSource: &TransformDataSource{
			S3DataSource: &TransformS3DataSource{
				S3Uri: jsii.String("s3://inputbucket/train"),
				S3DataType: tasks.S3DataType_S3_PREFIX,
			},
		},
	},
	TransformOutput: &TransformOutput{
		S3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
	},
	TransformResources: &TransformResources{
		InstanceCount: jsii.Number(1),
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M4, ec2.InstanceSize_XLARGE),
	},
})
```

### Create Endpoint

You can call the [`CreateEndpoint`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateEndpoint(this, jsii.String("SagemakerEndpoint"), &SageMakerCreateEndpointProps{
	EndpointName: sfn.JsonPath_StringAt(jsii.String("$.EndpointName")),
	EndpointConfigName: sfn.JsonPath_*StringAt(jsii.String("$.EndpointConfigName")),
})
```

### Create Endpoint Config

You can call the [`CreateEndpointConfig`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateEndpointConfig(this, jsii.String("SagemakerEndpointConfig"), &SageMakerCreateEndpointConfigProps{
	EndpointConfigName: jsii.String("MyEndpointConfig"),
	ProductionVariants: []productionVariant{
		&productionVariant{
			InitialInstanceCount: jsii.Number(2),
			InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M5, ec2.InstanceSize_XLARGE),
			ModelName: jsii.String("MyModel"),
			VariantName: jsii.String("awesome-variant"),
		},
	},
})
```

### Create Model

You can call the [`CreateModel`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html) API from a `Task` state.

```go
tasks.NewSageMakerCreateModel(this, jsii.String("Sagemaker"), &SageMakerCreateModelProps{
	ModelName: jsii.String("MyModel"),
	PrimaryContainer: tasks.NewContainerDefinition(&ContainerDefinitionOptions{
		Image: tasks.DockerImage_FromJsonExpression(sfn.JsonPath_StringAt(jsii.String("$.Model.imageName"))),
		Mode: tasks.Mode_SINGLE_MODEL,
		ModelS3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.TrainingJob.ModelArtifacts.S3ModelArtifacts")),
	}),
})
```

### Update Endpoint

You can call the [`UpdateEndpoint`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpoint.html) API from a `Task` state.

```go
tasks.NewSageMakerUpdateEndpoint(this, jsii.String("SagemakerEndpoint"), &SageMakerUpdateEndpointProps{
	EndpointName: sfn.JsonPath_StringAt(jsii.String("$.Endpoint.Name")),
	EndpointConfigName: sfn.JsonPath_*StringAt(jsii.String("$.Endpoint.EndpointConfig")),
})
```

## SNS

Step Functions supports [Amazon SNS](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sns.html) through the service integration pattern.

### Publish

You can call the [`Publish`](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) API from a `Task` state to publish to an SNS topic.

```go
topic := sns.NewTopic(this, jsii.String("Topic"))

// Use a field from the execution data as message.
task1 := tasks.NewSnsPublish(this, jsii.String("Publish1"), &SnsPublishProps{
	Topic: Topic,
	IntegrationPattern: sfn.IntegrationPattern_REQUEST_RESPONSE,
	Message: sfn.TaskInput_FromDataAt(jsii.String("$.state.message")),
	MessageAttributes: map[string]messageAttribute{
		"place": &messageAttribute{
			"value": sfn.JsonPath_stringAt(jsii.String("$.place")),
		},
		"pic": &messageAttribute{
			// BINARY must be explicitly set
			"dataType": tasks.MessageAttributeDataType_BINARY,
			"value": sfn.JsonPath_stringAt(jsii.String("$.pic")),
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
task2 := tasks.NewSnsPublish(this, jsii.String("Publish2"), &SnsPublishProps{
	Topic: Topic,
	Message: sfn.TaskInput_FromObject(map[string]interface{}{
		"field1": jsii.String("somedata"),
		"field2": sfn.JsonPath_stringAt(jsii.String("$.field2")),
	}),
})
```

## Step Functions

Step Functions supports [AWS Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/connect-stepfunctions.html) through the service integration pattern.

### Start Execution

You can manage [AWS Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/connect-stepfunctions.html) executions.

AWS Step Functions supports it's own [`StartExecution`](https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html) API as a service integration.

```go
// Define a state machine with one Pass state
child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &StateMachineProps{
	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("PassState"))),
})

// Include the state machine in a Task state with callback pattern
task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &StepFunctionsStartExecutionProps{
	StateMachine: child,
	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
	Input: sfn.TaskInput_FromObject(map[string]interface{}{
		"token": sfn.JsonPath_taskToken(),
		"foo": jsii.String("bar"),
	}),
	Name: jsii.String("MyExecutionName"),
})

// Define a second state machine with the Task state above
// Define a second state machine with the Task state above
sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &StateMachineProps{
	Definition: task,
})
```

You can utilize [Associate Workflow Executions](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-nested-workflows.html#nested-execution-startid)
via the `associateWithParent` property. This allows the Step Functions UI to link child
executions from parent executions, making it easier to trace execution flow across state machines.

```go
var child stateMachine

task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &StepFunctionsStartExecutionProps{
	StateMachine: child,
	AssociateWithParent: jsii.Boolean(true),
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

tasks.NewStepFunctionsInvokeActivity(this, jsii.String("Submit Job"), &StepFunctionsInvokeActivityProps{
	Activity: submitJobActivity,
})
```

Use the [Parameters](https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters) field to create a collection of key-value pairs that are passed as input.
The values of each can either be static values that you include in your state machine definition, or selected from either the input or the context object with a path.

```go
submitJobActivity := sfn.NewActivity(this, jsii.String("SubmitJob"))

tasks.NewStepFunctionsInvokeActivity(this, jsii.String("Submit Job"), &StepFunctionsInvokeActivityProps{
	Activity: submitJobActivity,
	Parameters: map[string]interface{}{
		"comment": jsii.String("Selecting what I care about."),
		"MyDetails": map[string]interface{}{
			"size": sfn.JsonPath_stringAt(jsii.String("$.product.details.size")),
			"exists": sfn.JsonPath_stringAt(jsii.String("$.product.availability")),
			"StaticValue": jsii.String("foo"),
		},
	},
})
```

## SQS

Step Functions supports [Amazon SQS](https://docs.aws.amazon.com/step-functions/latest/dg/connect-sqs.html)

### Send Message

You can call the [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) API from a `Task` state
to send a message to an SQS queue.

```go
queue := sqs.NewQueue(this, jsii.String("Queue"))

// Use a field from the execution data as message.
task1 := tasks.NewSqsSendMessage(this, jsii.String("Send1"), &SqsSendMessageProps{
	Queue: Queue,
	MessageBody: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
})

// Combine a field from the execution data with
// a literal object.
task2 := tasks.NewSqsSendMessage(this, jsii.String("Send2"), &SqsSendMessageProps{
	Queue: Queue,
	MessageBody: sfn.TaskInput_FromObject(map[string]interface{}{
		"field1": jsii.String("somedata"),
		"field2": sfn.JsonPath_stringAt(jsii.String("$.field2")),
	}),
})
```
