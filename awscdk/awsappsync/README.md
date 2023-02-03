# AWS AppSync Construct Library

The `@aws-cdk/aws-appsync` package contains constructs for building flexible
APIs that use GraphQL.

```go
import appsync "github.com/aws/aws-cdk-go/awscdk"
```

## Example

### DynamoDB

Example of a GraphQL API with `AWS_IAM` [authorization](#authorization) resolving into a DynamoDb
backend data source.

GraphQL schema file `schema.graphql`:

```gql
type demo {
  id: String!
  version: String!
}
type Query {
  getDemos: [ demo! ]
}
input DemoInput {
  version: String!
}
type Mutation {
  addDemo(input: DemoInput!): demo
}
```

CDK stack file `app-stack.ts`:

```go
api := appsync.NewGraphqlApi(this, jsii.String("Api"), &graphqlApiProps{
	name: jsii.String("demo"),
	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
	authorizationConfig: &authorizationConfig{
		defaultAuthorization: &authorizationMode{
			authorizationType: appsync.authorizationType_IAM,
		},
	},
	xrayEnabled: jsii.Boolean(true),
})

demoTable := dynamodb.NewTable(this, jsii.String("DemoTable"), &tableProps{
	partitionKey: &attribute{
		name: jsii.String("id"),
		type: dynamodb.attributeType_STRING,
	},
})

demoDS := api.addDynamoDbDataSource(jsii.String("demoDataSource"), demoTable)

// Resolver for the Query "getDemos" that scans the DynamoDb table and returns the entire list.
// Resolver Mapping Template Reference:
// https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-dynamodb.html
demoDS.createResolver(jsii.String("QueryGetDemosResolver"), &baseResolverProps{
	typeName: jsii.String("Query"),
	fieldName: jsii.String("getDemos"),
	requestMappingTemplate: appsync.mappingTemplate.dynamoDbScanTable(),
	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
})

// Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
demoDS.createResolver(jsii.String("MutationAddDemoResolver"), &baseResolverProps{
	typeName: jsii.String("Mutation"),
	fieldName: jsii.String("addDemo"),
	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbPutItem(appsync.primaryKey.partition(jsii.String("id")).auto(), appsync.values.projecting(jsii.String("input"))),
	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultItem(),
})

//To enable DynamoDB read consistency with the `MappingTemplate`:
demoDS.createResolver(jsii.String("QueryGetDemosConsistentResolver"), &baseResolverProps{
	typeName: jsii.String("Query"),
	fieldName: jsii.String("getDemosConsistent"),
	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbScanTable(jsii.Boolean(true)),
	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
})
```

### Aurora Serverless

AppSync provides a data source for executing SQL commands against Amazon Aurora
Serverless clusters. You can use AppSync resolvers to execute SQL statements
against the Data API with GraphQL queries, mutations, and subscriptions.

```go
// Build a data source for AppSync to access the database.
var api graphqlApi
// Create username and password secret for DB Cluster
secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &databaseSecretProps{
	username: jsii.String("clusteradmin"),
})

// The VPC to place the cluster in
vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))

// Create the serverless cluster, provide all values needed to customise the database.
cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &serverlessClusterProps{
	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
	vpc: vpc,
	credentials: map[string]*string{
		"username": jsii.String("clusteradmin"),
	},
	clusterIdentifier: jsii.String("db-endpoint-test"),
	defaultDatabaseName: jsii.String("demos"),
})
rdsDS := api.addRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))

// Set up a resolver for an RDS query.
rdsDS.createResolver(jsii.String("QueryGetDemosRdsResolver"), &baseResolverProps{
	typeName: jsii.String("Query"),
	fieldName: jsii.String("getDemosRds"),
	requestMappingTemplate: appsync.mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"SELECT * FROM demos\"\n    ]\n  }\n  ")),
	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])\n  ")),
})

// Set up a resolver for an RDS mutation.
rdsDS.createResolver(jsii.String("MutationAddDemoRdsResolver"), &baseResolverProps{
	typeName: jsii.String("Mutation"),
	fieldName: jsii.String("addDemoRds"),
	requestMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"INSERT INTO demos VALUES (:id, :version)\",\n      \"SELECT * WHERE id = :id\"\n    ],\n    \"variableMap\": {\n      \":id\": $util.toJson($util.autoId()),\n      \":version\": $util.toJson($ctx.args.version)\n    }\n  }\n  ")),
	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])\n  ")),
})
```

### HTTP Endpoints

GraphQL schema file `schema.graphql`:

```gql
type job {
  id: String!
  version: String!
}

input DemoInput {
  version: String!
}

type Mutation {
  callStepFunction(input: DemoInput!): job
}
```

GraphQL request mapping template `request.vtl`:

```json
{
  "version": "2018-05-29",
  "method": "POST",
  "resourcePath": "/",
  "params": {
    "headers": {
      "content-type": "application/x-amz-json-1.0",
      "x-amz-target":"AWSStepFunctions.StartExecution"
    },
    "body": {
      "stateMachineArn": "<your step functions arn>",
      "input": "{ \"id\": \"$context.arguments.id\" }"
    }
  }
}
```

GraphQL response mapping template `response.vtl`:

```json
{
  "id": "${context.result.id}"
}
```

CDK stack file `app-stack.ts`:

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
	name: jsii.String("api"),
	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
})

httpDs := api.addHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &httpDataSourceOptions{
	name: jsii.String("httpDsWithStepF"),
	description: jsii.String("from appsync to StepFunctions Workflow"),
	authorizationConfig: &awsIamConfig{
		signingRegion: jsii.String("us-east-1"),
		signingServiceName: jsii.String("states"),
	},
})

httpDs.createResolver(jsii.String("MutationCallStepFunctionResolver"), &baseResolverProps{
	typeName: jsii.String("Mutation"),
	fieldName: jsii.String("callStepFunction"),
	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
})
```

### Amazon OpenSearch Service

AppSync has builtin support for Amazon OpenSearch Service (successor to Amazon
Elasticsearch Service) from domains that are provisioned through your AWS account. You can
use AppSync resolvers to perform GraphQL operations such as queries, mutations, and
subscriptions.

```go
import opensearch "github.com/aws/aws-cdk-go/awscdk"

var api graphqlApi


user := iam.NewUser(this, jsii.String("User"))
domain := opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: opensearch.engineVersion_OPENSEARCH_2_3(),
	removalPolicy: awscdk.RemovalPolicy_DESTROY,
	fineGrainedAccessControl: &advancedSecurityOptions{
		masterUserArn: user.userArn,
	},
	encryptionAtRest: &encryptionAtRestOptions{
		enabled: jsii.Boolean(true),
	},
	nodeToNodeEncryption: jsii.Boolean(true),
	enforceHttps: jsii.Boolean(true),
})
ds := api.addOpenSearchDataSource(jsii.String("ds"), domain)

ds.createResolver(jsii.String("QueryGetTestsResolver"), &baseResolverProps{
	typeName: jsii.String("Query"),
	fieldName: jsii.String("getTests"),
	requestMappingTemplate: appsync.mappingTemplate.fromString(jSON.stringify(map[string]interface{}{
		"version": jsii.String("2017-02-28"),
		"operation": jsii.String("GET"),
		"path": jsii.String("/id/post/_search"),
		"params": map[string]map[string]interface{}{
			"headers": map[string]interface{}{
			},
			"queryString": map[string]interface{}{
			},
			"body": map[string]*f64{
				"from": jsii.Number(0),
				"size": jsii.Number(50),
			},
		},
	})),
	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("[\n    #foreach($entry in $context.result.hits.hits)\n    #if( $velocityCount > 1 ) , #end\n    $utils.toJson($entry.get(\"_source\"))\n    #end\n  ]")),
})
```

## Custom Domain Names

For many use cases you may want to associate a custom domain name with your
GraphQL API. This can be done during the API creation.

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
import route53 "github.com/aws/aws-cdk-go/awscdk"

// hosted zone and route53 features
var hostedZoneId string
zoneName := "example.com"


myDomainName := "api.example.com"
certificate := acm.NewCertificate(this, jsii.String("cert"), &certificateProps{
	domainName: myDomainName,
})
schema := appsync.NewSchemaFile(&schemaProps{
	filePath: jsii.String("mySchemaFile"),
})
api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
	name: jsii.String("myApi"),
	schema: schema,
	domainName: &domainOptions{
		certificate: certificate,
		domainName: myDomainName,
	},
})

// hosted zone for adding appsync domain
zone := route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("HostedZone"), &hostedZoneAttributes{
	hostedZoneId: jsii.String(hostedZoneId),
	zoneName: jsii.String(zoneName),
})

// create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
// create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
route53.NewCnameRecord(this, jsii.String("CnameApiRecord"), &cnameRecordProps{
	recordName: jsii.String("api"),
	zone: zone,
	domainName: api.appSyncDomainName,
})
```

## Log Group

AppSync automatically create a log group with the name `/aws/appsync/apis/<graphql_api_id>` upon deployment with
log data set to never expire. If you want to set a different expiration period, use the `logConfig.retention` property.

To obtain the GraphQL API's log group as a `logs.ILogGroup` use the `logGroup` property of the
`GraphqlApi` construct.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logConfig := &logConfig{
	retention: logs.retentionDays_ONE_WEEK,
}

appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
	authorizationConfig: &authorizationConfig{
	},
	name: jsii.String("myApi"),
	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("myApi.graphql"))),
	logConfig: logConfig,
})
```

## Schema

You can define a schema using from a local file using `SchemaFile.fromAsset`

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
	name: jsii.String("myApi"),
	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphl"))),
})
```

### ISchema

Alternative schema sources can be defined by implementing the `ISchema`
interface. An example of this is the `CodeFirstSchema` class provided in
[awscdk-appsync-utils](https://github.com/cdklabs/awscdk-appsync-utils)

## Imports

Any GraphQL Api that has been created outside the stack can be imported from
another stack into your CDK app. Utilizing the `fromXxx` function, you have
the ability to add data sources and resolvers through a `IGraphqlApi` interface.

```go
var api graphqlApi
var table table

importedApi := appsync.graphqlApi.fromGraphqlApiAttributes(this, jsii.String("IApi"), &graphqlApiAttributes{
	graphqlApiId: api.apiId,
	graphqlApiArn: api.arn,
})
importedApi.addDynamoDbDataSource(jsii.String("TableDataSource"), table)
```

If you don't specify `graphqlArn` in `fromXxxAttributes`, CDK will autogenerate
the expected `arn` for the imported api, given the `apiId`. For creating data
sources and resolvers, an `apiId` is sufficient.

## Authorization

There are multiple authorization types available for GraphQL API to cater to different
access use cases. They are:

* API Keys (`AuthorizationType.API_KEY`)
* Amazon Cognito User Pools (`AuthorizationType.USER_POOL`)
* OpenID Connect (`AuthorizationType.OPENID_CONNECT`)
* AWS Identity and Access Management (`AuthorizationType.AWS_IAM`)
* AWS Lambda (`AuthorizationType.AWS_LAMBDA`)

These types can be used simultaneously in a single API, allowing different types of clients to
access data. When you specify an authorization type, you can also specify the corresponding
authorization mode to finish defining your authorization. For example, this is a GraphQL API
with AWS Lambda Authorization.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var authFunction function


appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
	name: jsii.String("api"),
	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
	authorizationConfig: &authorizationConfig{
		defaultAuthorization: &authorizationMode{
			authorizationType: appsync.authorizationType_LAMBDA,
			lambdaAuthorizerConfig: &lambdaAuthorizerConfig{
				handler: authFunction,
			},
		},
	},
})
```

## Permissions

When using `AWS_IAM` as the authorization type for GraphQL API, an IAM Role
with correct permissions must be used for access to API.

When configuring permissions, you can specify specific resources to only be
accessible by `IAM` authorization. For example, if you want to only allow mutability
for `IAM` authorized access you would configure the following.

In `schema.graphql`:

```gql
type Mutation {
  updateExample(...): ...
    @aws_iam
}
```

In `IAM`:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "appsync:GraphQL"
      ],
      "Resource": [
        "arn:aws:appsync:REGION:ACCOUNT_ID:apis/GRAPHQL_ID/types/Mutation/fields/updateExample"
      ]
    }
  ]
}
```

See [documentation](https://docs.aws.amazon.com/appsync/latest/devguide/security.html#aws-iam-authorization) for more details.

To make this easier, CDK provides `grant` API.

Use the `grant` function for more granular authorization.

```go
var api graphqlApi
role := iam.NewRole(this, jsii.String("Role"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

api.grant(role, appsync.iamResource.custom(jsii.String("types/Mutation/fields/updateExample")), jsii.String("appsync:GraphQL"))
```

### IamResource

In order to use the `grant` functions, you need to use the class `IamResource`.

* `IamResource.custom(...arns)` permits custom ARNs and requires an argument.
* `IamResouce.ofType(type, ...fields)` permits ARNs for types and their fields.
* `IamResource.all()` permits ALL resources.

### Generic Permissions

Alternatively, you can use more generic `grant` functions to accomplish the same usage.

These include:

* grantMutation (use to grant access to Mutation fields)
* grantQuery (use to grant access to Query fields)
* grantSubscription (use to grant access to Subscription fields)

```go
var api graphqlApi
var role role


// For generic types
api.grantMutation(role, jsii.String("updateExample"))

// For custom types and granular design
api.grant(role, appsync.iamResource.ofType(jsii.String("Mutation"), jsii.String("updateExample")), jsii.String("appsync:GraphQL"))
```

## Pipeline Resolvers and AppSync Functions

AppSync Functions are local functions that perform certain operations onto a
backend data source. Developers can compose operations (Functions) and execute
them in sequence with Pipeline Resolvers.

```go
var api graphqlApi


appsyncFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &appsyncFunctionProps{
	name: jsii.String("appsync_function"),
	api: api,
	dataSource: api.addNoneDataSource(jsii.String("none")),
	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
})
```

AppSync Functions are used in tandem with pipeline resolvers to compose multiple
operations.

```go
var api graphqlApi
var appsyncFunction appsyncFunction


pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &resolverProps{
	api: api,
	dataSource: api.addNoneDataSource(jsii.String("none")),
	typeName: jsii.String("typeName"),
	fieldName: jsii.String("fieldName"),
	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("beforeRequest.vtl")),
	pipelineConfig: []iAppsyncFunction{
		appsyncFunction,
	},
	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("afterResponse.vtl")),
})
```

### JS Functions and Resolvers

JS Functions and resolvers are also supported. You can use a `.js` file within your CDK project, or specify your function code inline.

```go
var api graphqlApi


myJsFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &appsyncFunctionProps{
	name: jsii.String("my_js_function"),
	api: api,
	dataSource: api.addNoneDataSource(jsii.String("none")),
	code: appsync.code.fromAsset(jsii.String("directory/function_code.js")),
	runtime: appsync.functionRuntime_JS_1_0_0(),
})

appsync.NewResolver(this, jsii.String("PipelineResolver"), &resolverProps{
	api: api,
	typeName: jsii.String("typeName"),
	fieldName: jsii.String("fieldName"),
	code: appsync.*code.fromInline(jsii.String("\n    // The before step\n    export function request(...args) {\n      console.log(args);\n      return {}\n    }\n\n    // The after step\n    export function response(ctx) {\n      return ctx.prev.result\n    }\n  ")),
	runtime: appsync.*functionRuntime_JS_1_0_0(),
	pipelineConfig: []iAppsyncFunction{
		myJsFunction,
	},
})
```

Learn more about Pipeline Resolvers and AppSync Functions [here](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html).
