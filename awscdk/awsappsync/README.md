# AWS AppSync Construct Library

The `aws-cdk-lib/aws-appsync` package contains constructs for building flexible
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
api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
	Name: jsii.String("demo"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphql"))),
	AuthorizationConfig: &AuthorizationConfig{
		DefaultAuthorization: &AuthorizationMode{
			AuthorizationType: appsync.AuthorizationType_IAM,
		},
	},
	XrayEnabled: jsii.Boolean(true),
})

demoTable := dynamodb.NewTable(this, jsii.String("DemoTable"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
})

demoDS := api.AddDynamoDbDataSource(jsii.String("demoDataSource"), demoTable)

// Resolver for the Query "getDemos" that scans the DynamoDb table and returns the entire list.
// Resolver Mapping Template Reference:
// https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-dynamodb.html
demoDS.CreateResolver(jsii.String("QueryGetDemosResolver"), &BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getDemos"),
	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbScanTable(),
	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultList(),
})

// Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
demoDS.CreateResolver(jsii.String("MutationAddDemoResolver"), &BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("addDemo"),
	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbPutItem(appsync.PrimaryKey_Partition(jsii.String("id")).Auto(), appsync.Values_Projecting(jsii.String("input"))),
	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultItem(),
})

//To enable DynamoDB read consistency with the `MappingTemplate`:
demoDS.CreateResolver(jsii.String("QueryGetDemosConsistentResolver"), &BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getDemosConsistent"),
	RequestMappingTemplate: appsync.MappingTemplate_*DynamoDbScanTable(jsii.Boolean(true)),
	ResponseMappingTemplate: appsync.MappingTemplate_*DynamoDbResultList(),
})
```

### Aurora Serverless

AppSync provides a data source for executing SQL commands against Amazon Aurora
Serverless clusters. You can use AppSync resolvers to execute SQL statements
against the Data API with GraphQL queries, mutations, and subscriptions.

#### Aurora Serverless V1 Cluster

```go
// Build a data source for AppSync to access the database.
var api graphqlApi
// Create username and password secret for DB Cluster
secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
	Username: jsii.String("clusteradmin"),
})

// The VPC to place the cluster in
vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))

// Create the serverless cluster, provide all values needed to customise the database.
cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &ServerlessClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
	Vpc: Vpc,
	Credentials: map[string]*string{
		"username": jsii.String("clusteradmin"),
	},
	ClusterIdentifier: jsii.String("db-endpoint-test"),
	DefaultDatabaseName: jsii.String("demos"),
})
rdsDS := api.AddRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))

// Set up a resolver for an RDS query.
rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getDemosRds"),
	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
	  {
	    "version": "2018-05-29",
	    "statements": [
	      "SELECT * FROM demos"
	    ]
	  }
	  `)),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
	  `)),
})

// Set up a resolver for an RDS mutation.
rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("addDemoRds"),
	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	  {
	    "version": "2018-05-29",
	    "statements": [
	      "INSERT INTO demos VALUES (:id, :version)",
	      "SELECT * WHERE id = :id"
	    ],
	    "variableMap": {
	      ":id": $util.toJson($util.autoId()),
	      ":version": $util.toJson($ctx.args.version)
	    }
	  }
	  `)),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
	  `)),
})
```

#### Aurora Serverless V2 Cluster

```go
// Build a data source for AppSync to access the database.
var api graphqlApi
// Create username and password secret for DB Cluster
secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
	Username: jsii.String("clusteradmin"),
})

// The VPC to place the cluster in
vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))

// Create the serverless cluster, provide all values needed to customise the database.
cluster := rds.NewDatabaseCluster(this, jsii.String("AuroraClusterV2"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
		Version: rds.AuroraPostgresEngineVersion_VER_15_5(),
	}),
	Credentials: map[string]*string{
		"username": jsii.String("clusteradmin"),
	},
	ClusterIdentifier: jsii.String("db-endpoint-test"),
	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
	ServerlessV2MinCapacity: jsii.Number(2),
	ServerlessV2MaxCapacity: jsii.Number(10),
	Vpc: Vpc,
	DefaultDatabaseName: jsii.String("demos"),
	EnableDataApi: jsii.Boolean(true),
})
rdsDS := api.AddRdsDataSourceV2(jsii.String("rds"), cluster, secret, jsii.String("demos"))

// Set up a resolver for an RDS query.
rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getDemosRds"),
	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
	  {
	    "version": "2018-05-29",
	    "statements": [
	      "SELECT * FROM demos"
	    ]
	  }
	  `)),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
	  `)),
})

// Set up a resolver for an RDS mutation.
rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("addDemoRds"),
	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	  {
	    "version": "2018-05-29",
	    "statements": [
	      "INSERT INTO demos VALUES (:id, :version)",
	      "SELECT * WHERE id = :id"
	    ],
	    "variableMap": {
	      ":id": $util.toJson($util.autoId()),
	      ":version": $util.toJson($ctx.args.version)
	    }
	  }
	  `)),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
	  `)),
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

type Query {
  _placeholder: String
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
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("api"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphql"))),
})

httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
	Name: jsii.String("httpDsWithStepF"),
	Description: jsii.String("from appsync to StepFunctions Workflow"),
	AuthorizationConfig: &AwsIamConfig{
		SigningRegion: jsii.String("us-east-1"),
		SigningServiceName: jsii.String("states"),
	},
})

httpDs.CreateResolver(jsii.String("MutationCallStepFunctionResolver"), &BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("callStepFunction"),
	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
})
```

### EventBridge

Integrating AppSync with EventBridge enables developers to use EventBridge rules to route commands for GraphQL mutations
that need to perform any one of a variety of asynchronous tasks. More broadly, it enables teams to expose an event bus
as a part of a GraphQL schema.

GraphQL schema file `schema.graphql`:

```gql
schema {
    query: Query
    mutation: Mutation
}

type Query {
    event(id:ID!): Event
}

type Mutation {
    emitEvent(id: ID!, name: String): PutEventsResult!
}

type Event {
    id: ID!
    name: String!
}

type Entry {
    ErrorCode: String
    ErrorMessage: String
    EventId: String
}

type PutEventsResult {
    Entries: [Entry!]
    FailedEntry: Int
}
```

GraphQL request mapping template `request.vtl`:

```
{
    "version" : "2018-05-29",
    "operation": "PutEvents",
    "events" : [
        {
            "source": "integ.appsync.eventbridge",
            "detailType": "Mutation.emitEvent",
            "detail": $util.toJson($context.arguments)
        }
    ]
}
```

GraphQL response mapping template `response.vtl`:

```
$util.toJson($ctx.result)'
```

This response mapping template simply converts the EventBridge PutEvents result to JSON.
For details about the response see the
[documentation](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html).
Additional logic can be added to the response template to map the response type, or to error in the event of failed
events. More information can be found
[here](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-eventbridge.html).

CDK stack file `app-stack.ts`:

```go
import events "github.com/aws/aws-cdk-go/awscdk"


api := appsync.NewGraphqlApi(this, jsii.String("EventBridgeApi"), &GraphqlApiProps{
	Name: jsii.String("EventBridgeApi"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
})

bus := events.NewEventBus(this, jsii.String("DestinationEventBus"), &EventBusProps{
})

dataSource := api.AddEventBridgeDataSource(jsii.String("NoneDS"), bus)

dataSource.CreateResolver(jsii.String("EventResolver"), &BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("emitEvent"),
	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
})
```

### Amazon OpenSearch Service

AppSync has builtin support for Amazon OpenSearch Service (successor to Amazon
Elasticsearch Service) from domains that are provisioned through your AWS account. You can
use AppSync resolvers to perform GraphQL operations such as queries, mutations, and
subscriptions.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var api graphqlApi


user := iam.NewUser(this, jsii.String("User"))
domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: opensearch.EngineVersion_OPENSEARCH_2_3(),
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	FineGrainedAccessControl: &AdvancedSecurityOptions{
		MasterUserArn: user.UserArn,
	},
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	NodeToNodeEncryption: jsii.Boolean(true),
	EnforceHttps: jsii.Boolean(true),
})
ds := api.AddOpenSearchDataSource(jsii.String("ds"), domain)

ds.CreateResolver(jsii.String("QueryGetTestsResolver"), &BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getTests"),
	RequestMappingTemplate: appsync.MappingTemplate_FromString(jSON.stringify(map[string]interface{}{
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
	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`[
	    #foreach($entry in $context.result.hits.hits)
	    #if( $velocityCount > 1 ) , #end
	    $utils.toJson($entry.get("_source"))
	    #end
	  ]`)),
})
```

## Merged APIs

AppSync supports [Merged APIs](https://docs.aws.amazon.com/appsync/latest/devguide/merged-api.html) which can be used to merge multiple source APIs into a single API.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"


// first source API
firstApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
	Name: jsii.String("FirstSourceAPI"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
})

// second source API
secondApi := appsync.NewGraphqlApi(this, jsii.String("SecondSourceAPI"), &GraphqlApiProps{
	Name: jsii.String("SecondSourceAPI"),
	Definition: appsync.Definition_*FromFile(path.join(__dirname, jsii.String("appsync.merged-api-2.graphql"))),
})

// Merged API
mergedApi := appsync.NewGraphqlApi(this, jsii.String("MergedAPI"), &GraphqlApiProps{
	Name: jsii.String("MergedAPI"),
	Definition: appsync.Definition_FromSourceApis(&SourceApiOptions{
		SourceApis: []sourceApi{
			&sourceApi{
				SourceApi: firstApi,
				MergeType: appsync.MergeType_MANUAL_MERGE,
			},
			&sourceApi{
				SourceApi: secondApi,
				MergeType: appsync.MergeType_AUTO_MERGE,
			},
		},
	}),
})
```

## Merged APIs Across Different Stacks

The SourceApiAssociation construct allows you to define a SourceApiAssociation to a Merged API in a different stack or account. This allows a source API owner the ability to associate it to an existing Merged API itself.

```go
sourceApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
	Name: jsii.String("FirstSourceAPI"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
})

importedMergedApi := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedMergedApi"), &GraphqlApiAttributes{
	GraphqlApiId: jsii.String("MyApiId"),
	GraphqlApiArn: jsii.String("MyApiArn"),
})

importedExecutionRole := iam.Role_FromRoleArn(this, jsii.String("ExecutionRole"), jsii.String("arn:aws:iam::ACCOUNT:role/MyExistingRole"))
appsync.NewSourceApiAssociation(this, jsii.String("SourceApiAssociation2"), &SourceApiAssociationProps{
	SourceApi: sourceApi,
	MergedApi: importedMergedApi,
	MergeType: appsync.MergeType_MANUAL_MERGE,
	MergedApiExecutionRole: importedExecutionRole,
})
```

## Merge Source API Update Within CDK Deployment

The SourceApiAssociationMergeOperation construct available in the [awscdk-appsync-utils](https://github.com/cdklabs/awscdk-appsync-utils) package provides the ability to merge a source API to a Merged API via a custom
resource. If the merge operation fails with a conflict, the stack update will fail and rollback the changes to the source API in the stack in order to prevent merge conflicts and ensure the source API changes are always propagated to the Merged API.

## Custom Domain Names

For many use cases you may want to associate a custom domain name with your
GraphQL API. This can be done during the API creation.

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// hosted zone and route53 features
var hostedZoneId string
zoneName := "example.com"


myDomainName := "api.example.com"
certificate := acm.NewCertificate(this, jsii.String("cert"), &CertificateProps{
	DomainName: myDomainName,
})
schema := appsync.NewSchemaFile(&SchemaProps{
	FilePath: jsii.String("mySchemaFile"),
})
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
	Definition: appsync.Definition_FromSchema(schema),
	DomainName: &DomainOptions{
		Certificate: *Certificate,
		DomainName: myDomainName,
	},
})

// hosted zone for adding appsync domain
zone := route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
	HostedZoneId: jsii.String(HostedZoneId),
	ZoneName: jsii.String(ZoneName),
})

// create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
// create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
route53.NewCnameRecord(this, jsii.String("CnameApiRecord"), &CnameRecordProps{
	RecordName: jsii.String("api"),
	Zone: Zone,
	DomainName: api.appSyncDomainName,
})
```

## Log Group

AppSync automatically create a log group with the name `/aws/appsync/apis/<graphql_api_id>` upon deployment with
log data set to never expire. If you want to set a different expiration period, use the `logConfig.retention` property.

Also you can choose the log level by setting the `logConfig.fieldLogLevel` property.

For more information, see [CloudWatch logs](https://docs.aws.amazon.com/en_us/appsync/latest/devguide/monitoring.html#cwl).

To obtain the GraphQL API's log group as a `logs.ILogGroup` use the `logGroup` property of the
`GraphqlApi` construct.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	AuthorizationConfig: &AuthorizationConfig{
	},
	Name: jsii.String("myApi"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("myApi.graphql"))),
	LogConfig: &LogConfig{
		FieldLogLevel: appsync.FieldLogLevel_INFO,
		Retention: logs.RetentionDays_ONE_WEEK,
	},
})
```

## Schema

You can define a schema using from a local file using `Definition.fromFile`

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphl"))),
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

importedApi := appsync.graphqlApi_FromGraphqlApiAttributes(this, jsii.String("IApi"), &GraphqlApiAttributes{
	GraphqlApiId: api.ApiId,
	GraphqlApiArn: api.Arn,
})
importedApi.AddDynamoDbDataSource(jsii.String("TableDataSource"), table)
```

If you don't specify `graphqlArn` in `fromXxxAttributes`, CDK will autogenerate
the expected `arn` for the imported api, given the `apiId`. For creating data
sources and resolvers, an `apiId` is sufficient.

## Private APIs

By default all AppSync GraphQL APIs are public and can be accessed from the internet.
For customers that want to limit access to be from their VPC, the optional API `visibility` property can be set to `Visibility.PRIVATE`
at creation time. To explicitly create a public API, the `visibility` property should be set to `Visibility.GLOBAL`.
If visibility is not set, the service will default to `GLOBAL`.

CDK stack file `app-stack.ts`:

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("MyPrivateAPI"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	Visibility: appsync.Visibility_PRIVATE,
})
```

See [documentation](https://docs.aws.amazon.com/appsync/latest/devguide/using-private-apis.html)
for more details about Private APIs

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


appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("api"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.test.graphql"))),
	AuthorizationConfig: &AuthorizationConfig{
		DefaultAuthorization: &AuthorizationMode{
			AuthorizationType: appsync.AuthorizationType_LAMBDA,
			LambdaAuthorizerConfig: &LambdaAuthorizerConfig{
				Handler: authFunction,
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
var api iGraphqlApi
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

api.Grant(role, appsync.IamResource_Custom(jsii.String("types/Mutation/fields/updateExample")), jsii.String("appsync:GraphQL"))
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
var api iGraphqlApi
var role role


// For generic types
api.GrantMutation(role, jsii.String("updateExample"))

// For custom types and granular design
api.Grant(role, appsync.IamResource_OfType(jsii.String("Mutation"), jsii.String("updateExample")), jsii.String("appsync:GraphQL"))
```

## Pipeline Resolvers and AppSync Functions

AppSync Functions are local functions that perform certain operations onto a
backend data source. Developers can compose operations (Functions) and execute
them in sequence with Pipeline Resolvers.

```go
var api graphqlApi


appsyncFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &AppsyncFunctionProps{
	Name: jsii.String("appsync_function"),
	Api: Api,
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
})
```

When using the `LambdaDataSource`, you can control the maximum number of resolver request
inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation
by setting the `maxBatchSize` property.

```go
var api graphqlApi
var lambdaDataSource lambdaDataSource


appsyncFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &AppsyncFunctionProps{
	Name: jsii.String("appsync_function"),
	Api: Api,
	DataSource: lambdaDataSource,
	MaxBatchSize: jsii.Number(10),
})
```

AppSync Functions are used in tandem with pipeline resolvers to compose multiple
operations.

```go
var api graphqlApi
var appsyncFunction appsyncFunction


pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &ResolverProps{
	Api: Api,
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	TypeName: jsii.String("typeName"),
	FieldName: jsii.String("fieldName"),
	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("beforeRequest.vtl")),
	PipelineConfig: []iAppsyncFunction{
		appsyncFunction,
	},
	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("afterResponse.vtl")),
})
```

### JS Functions and Resolvers

JS Functions and resolvers are also supported. You can use a `.js` file within your CDK project, or specify your function code inline.

```go
var api graphqlApi


myJsFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &AppsyncFunctionProps{
	Name: jsii.String("my_js_function"),
	Api: Api,
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	Code: appsync.Code_FromAsset(jsii.String("directory/function_code.js")),
	Runtime: appsync.FunctionRuntime_JS_1_0_0(),
})

appsync.NewResolver(this, jsii.String("PipelineResolver"), &ResolverProps{
	Api: Api,
	TypeName: jsii.String("typeName"),
	FieldName: jsii.String("fieldName"),
	Code: appsync.Code_FromInline(jsii.String(`
	    // The before step
	    export function request(...args) {
	      console.log(args);
	      return {}
	    }

	    // The after step
	    export function response(ctx) {
	      return ctx.prev.result
	    }
	  `)),
	Runtime: appsync.FunctionRuntime_JS_1_0_0(),
	PipelineConfig: []iAppsyncFunction{
		myJsFunction,
	},
})
```

Learn more about Pipeline Resolvers and AppSync Functions [here](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html).

## Introspection

By default, AppSync allows you to use introspection queries.

For customers that want to limit access to be introspection queries, the `introspectionConfig` property can be set to `IntrospectionConfig.DISABLED` at creation time.
If `introspectionConfig` is not set, the service will default to `ENABLED`.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("DisableIntrospectionApi"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	IntrospectionConfig: appsync.IntrospectionConfig_DISABLED,
})
```

## Query Depth Limits

By default, queries are able to process an unlimited amount of nested levels.
Limiting queries to a specified amount of nested levels has potential implications for the performance and flexibility of your project.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("LimitQueryDepths"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	QueryDepthLimit: jsii.Number(2),
})
```

## Resolver Count Limits

You can control how many resolvers each query can process.
By default, each query can process up to 10000 resolvers.
By setting a limit AppSync will not handle any resolvers past a certain number limit.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("LimitResolverCount"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	ResolverCountLimit: jsii.Number(2),
})
```

## Environment Variables

To use environment variables in resolvers, you can use the `environmentVariables` property and
the `addEnvironmentVariable` method.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("api"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	EnvironmentVariables: map[string]*string{
		"EnvKey1": jsii.String("non-empty-1"),
	},
})

api.AddEnvironmentVariable(jsii.String("EnvKey2"), jsii.String("non-empty-2"))
```

## Configure an EventBridge target that invokes an AppSync GraphQL API

Configuring the target relies on the `graphQLEndpointArn` property.

Use the `AppSync` event target to trigger an AppSync GraphQL API. You need to
create an `AppSync.GraphqlApi` configured with `AWS_IAM` authorization mode.

The code snippet below creates a AppSync GraphQL API target that is invoked, calling the `publish` mutation.

```go
import events "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var rule rule
var api graphqlApi


rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
	Variables: events.RuleTargetInput_FromObject(map[string]*string{
		"message": jsii.String("hello world"),
	}),
}))
```

## Owner Contact

You can set the owner contact information for an API resource.
This field accepts any string input with a length of 0 - 256 characters.

```go
api := appsync.NewGraphqlApi(this, jsii.String("OwnerContact"), &GraphqlApiProps{
	Name: jsii.String("OwnerContact"),
	Definition: appsync.Definition_FromSchema(appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.test.graphql")))),
	OwnerContact: jsii.String("test-owner-contact"),
})
```
