# AWS AppSync Construct Library

The `aws-cdk-lib/aws-appsync` package contains constructs for building flexible
APIs that use [GraphQL](https://docs.aws.amazon.com/appsync/latest/devguide/what-is-appsync.html) and [Events](https://docs.aws.amazon.com/appsync/latest/eventapi/event-api-welcome.html).

```go
import appsync "github.com/aws/aws-cdk-go/awscdk"
```

## GraphQL

### Example

#### DynamoDB

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
  getDemosConsistent: [demo!]
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

#### Aurora Serverless

AppSync provides a data source for executing SQL commands against Amazon Aurora
Serverless clusters. You can use AppSync resolvers to execute SQL statements
against the Data API with GraphQL queries, mutations, and subscriptions.

##### Aurora Serverless V1 Cluster

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

##### Aurora Serverless V2 Cluster

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

#### HTTP Endpoints

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

#### EventBridge

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

#### Amazon OpenSearch Service

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

### Merged APIs

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

### Merged APIs Across Different Stacks

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

### Merge Source API Update Within CDK Deployment

The SourceApiAssociationMergeOperation construct available in the [awscdk-appsync-utils](https://github.com/cdklabs/awscdk-appsync-utils) package provides the ability to merge a source API to a Merged API via a custom
resource. If the merge operation fails with a conflict, the stack update will fail and rollback the changes to the source API in the stack in order to prevent merge conflicts and ensure the source API changes are always propagated to the Merged API.

### Custom Domain Names

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

### Log Group

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

### Schema

You can define a schema using from a local file using `Definition.fromFile`

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphl"))),
})
```

#### ISchema

Alternative schema sources can be defined by implementing the `ISchema`
interface. An example of this is the `CodeFirstSchema` class provided in
[awscdk-appsync-utils](https://github.com/cdklabs/awscdk-appsync-utils)

### Imports

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

### Private APIs

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

### Authorization

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

### Permissions

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

#### IamResource

In order to use the `grant` functions, you need to use the class `IamResource`.

* `IamResource.custom(...arns)` permits custom ARNs and requires an argument.
* `IamResouce.ofType(type, ...fields)` permits ARNs for types and their fields.
* `IamResource.all()` permits ALL resources.

#### Generic Permissions

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

### Pipeline Resolvers and AppSync Functions

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

#### JS Functions and Resolvers

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

### Introspection

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

### Query Depth Limits

By default, queries are able to process an unlimited amount of nested levels.
Limiting queries to a specified amount of nested levels has potential implications for the performance and flexibility of your project.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("LimitQueryDepths"),
	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
	QueryDepthLimit: jsii.Number(2),
})
```

### Resolver Count Limits

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

### Environment Variables

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

### Configure an EventBridge target that invokes an AppSync GraphQL API

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

### Owner Contact

You can set the owner contact information for an API resource.
This field accepts any string input with a length of 0 - 256 characters.

```go
api := appsync.NewGraphqlApi(this, jsii.String("OwnerContact"), &GraphqlApiProps{
	Name: jsii.String("OwnerContact"),
	Definition: appsync.Definition_FromSchema(appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.test.graphql")))),
	OwnerContact: jsii.String("test-owner-contact"),
})
```

## Events

### Example

AWS AppSync Events lets you create secure and performant serverless WebSocket APIs that can broadcast real-time event data to millions of subscribers, without you having to manage connections or resource scaling.

```go
apiKeyProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
}

api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
	ApiName: jsii.String("Api"),
	OwnerContact: jsii.String("OwnerContact"),
	AuthorizationConfig: &EventApiAuthConfig{
		AuthProviders: []appSyncAuthProvider{
			apiKeyProvider,
		},
		ConnectionAuthModeTypes: []appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
	},
})

api.AddChannelNamespace(jsii.String("default"))
```

### Authorization

AWS AppSync Events offers the following authorization types to secure Event APIs: API keys, Lambda, IAM, OpenID Connect, and Amazon Cognito user pools.
Each option provides a different method of security:

* API Keys (`AppSyncAuthorizationType.API_KEY`)
* Amazon Cognito User Pools (`AppSyncAuthorizationType.USER_POOL`)
* OpenID Connect (`AppSyncAuthorizationType.OIDC`)
* AWS Identity and Access Management (`AppSyncAuthorizationType.IAM`)
* AWS Lambda (`AppSyncAuthorizationType.LAMBDA`)

When you define your API, you configure the authorization mode to connect to your Event API WebSocket.
You also configure the default authorization modes to use when publishing and subscribing to messages.
If you don't specify any authorization providers, an API key will be created for you as the authorization mode for the API.

For mor information, see [Configuring authorization and authentication to secure Event APIs](https://docs.aws.amazon.com/appsync/latest/eventapi/configure-event-api-auth.html).

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var handler function


iamProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_IAM,
}

apiKeyProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
}

lambdaProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_LAMBDA,
	LambdaAuthorizerConfig: &AppSyncLambdaAuthorizerConfig{
		Handler: *Handler,
		ResultsCacheTtl: awscdk.Duration_Minutes(jsii.Number(6)),
		ValidationRegex: jsii.String("test"),
	},
}

api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
	ApiName: jsii.String("api"),
	AuthorizationConfig: &EventApiAuthConfig{
		// set auth providers
		AuthProviders: []appSyncAuthProvider{
			iamProvider,
			apiKeyProvider,
			lambdaProvider,
		},
		ConnectionAuthModeTypes: []appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_IAM,
		},
		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_LAMBDA,
		},
	},
})

api.AddChannelNamespace(jsii.String("default"))
```

If you don't specify any overrides for the `connectionAuthModeTypes`, `defaultPublishAuthModeTypes`, and `defaultSubscribeAuthModeTypes` parameters then all `authProviders` defined are included as default authorization mode types for connection, publish, and subscribe.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var handler function


iamProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_IAM,
}

apiKeyProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
}

/* API with IAM and API Key providers.
 * Connection, default publish and default subscribe
 * can be done with either IAM and API Key.
 */
api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
	ApiName: jsii.String("api"),
	AuthorizationConfig: &EventApiAuthConfig{
		// set auth providers
		AuthProviders: []appSyncAuthProvider{
			iamProvider,
			apiKeyProvider,
		},
	},
})

api.AddChannelNamespace(jsii.String("default"))
```

### Data Sources

With AWS AppSync Events, you can configure data source integrations with Amazon DynamoDB, Amazon Aurora Serverless, Amazon EventBridge, Amazon Bedrock Runtime, AWS Lambda, Amazon OpenSearch Service, and HTTP endpoints. The Event API can be associated with the data source and you can use the data source as an integration in your channel namespace event handlers for `onPublish` and `onSubscribe` operations.

Below are examples for how you add the various data sources to you Event API.

#### Amazon DynamoDB

```go
api := appsync.NewEventApi(this, jsii.String("EventApiDynamoDB"), &EventApiProps{
	ApiName: jsii.String("DynamoDBEventApi"),
})

table := dynamodb.NewTable(this, jsii.String("table"), &TableProps{
	TableName: jsii.String("event-messages"),
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
})

dataSource := api.AddDynamoDbDataSource(jsii.String("ddbsource"), table)
```

#### Amazon Aurora Serverless

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

databaseName := "mydb"
cluster := rds.NewDatabaseCluster(this, jsii.String("Cluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
		Version: rds.AuroraPostgresEngineVersion_VER_16_6(),
	}),
	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
	Vpc: vpc,
	Credentials: map[string]*string{
		"username": jsii.String("clusteradmin"),
	},
	DefaultDatabaseName: databaseName,
	EnableDataApi: jsii.Boolean(true),
})

secret := secretsmanager.Secret_FromSecretNameV2(this, jsii.String("Secret"), jsii.String("db-secretName"))

api := appsync.NewEventApi(this, jsii.String("EventApiRds"), &EventApiProps{
	ApiName: jsii.String("RdsEventApi"),
})

dataSource := api.AddRdsDataSource(jsii.String("rdsds"), cluster, secret, databaseName)
```

#### Amazon EventBridge

```go
import events "github.com/aws/aws-cdk-go/awscdk"


api := appsync.NewEventApi(this, jsii.String("EventApiEventBridge"), &EventApiProps{
	ApiName: jsii.String("EventBridgeEventApi"),
})

eventBus := events.NewEventBus(this, jsii.String("test-bus"))

dataSource := api.AddEventBridgeDataSource(jsii.String("eventbridgeds"), eventBus)
```

#### AWS Lambda

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var lambdaDs function


api := appsync.NewEventApi(this, jsii.String("EventApiLambda"), &EventApiProps{
	ApiName: jsii.String("LambdaEventApi"),
})

dataSource := api.AddLambdaDataSource(jsii.String("lambdads"), lambdaDs)
```

#### Amazon OpenSearch Service

```go
import "github.com/aws/aws-cdk-go/awscdk"


domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: opensearch.EngineVersion_OPENSEARCH_2_17(),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	NodeToNodeEncryption: jsii.Boolean(true),
	EnforceHttps: jsii.Boolean(true),
	Capacity: &CapacityConfig{
		MultiAzWithStandbyEnabled: jsii.Boolean(false),
	},
	Ebs: &EbsOptions{
		Enabled: jsii.Boolean(true),
		VolumeSize: jsii.Number(10),
	},
})
api := appsync.NewEventApi(this, jsii.String("EventApiOpenSearch"), &EventApiProps{
	ApiName: jsii.String("OpenSearchEventApi"),
})

dataSource := api.AddOpenSearchDataSource(jsii.String("opensearchds"), domain)
```

#### HTTP Endpoints

```go
import "github.com/aws/aws-cdk-go/awscdk"


api := appsync.NewEventApi(this, jsii.String("EventApiHttp"), &EventApiProps{
	ApiName: jsii.String("HttpEventApi"),
})

randomApi := apigw.NewRestApi(this, jsii.String("RandomApi"))
randomRoute := randomApi.Root.AddResource(jsii.String("random"))
randomRoute.AddMethod(jsii.String("GET"), apigw.NewMockIntegration(&IntegrationOptions{
	IntegrationResponses: []integrationResponse{
		&integrationResponse{
			StatusCode: jsii.String("200"),
			ResponseTemplates: map[string]*string{
				"application/json": jsii.String("my-random-value"),
			},
		},
	},
	PassthroughBehavior: apigw.PassthroughBehavior_NEVER,
	RequestTemplates: map[string]*string{
		"application/json": jsii.String("{ \"statusCode\": 200 }"),
	},
}), &MethodOptions{
	MethodResponses: []methodResponse{
		&methodResponse{
			StatusCode: jsii.String("200"),
		},
	},
})

dataSource := api.AddHttpDataSource(jsii.String("httpsource"), fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com", randomApi.RestApiId, this.Region))
```

### Custom Domain Names

With AWS AppSync, you can use custom domain names to configure a single, memorable domain that works for your Event APIs.
You can set custom domain by setting `domainName`. Also you can get custom HTTP/Realtime endpoint by `customHttpEndpoint`, `customRealtimeEndpoint`.

For more information, see [Configuring custom domain names for Event APIs](https://docs.aws.amazon.com/appsync/latest/eventapi/event-api-custom-domains.html).

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
import route53 "github.com/aws/aws-cdk-go/awscdk"


myDomainName := "api.example.com"
certificate := acm.NewCertificate(this, jsii.String("cert"), &CertificateProps{
	DomainName: myDomainName,
})

apiKeyProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
}

api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
	ApiName: jsii.String("Api"),
	OwnerContact: jsii.String("OwnerContact"),
	AuthorizationConfig: &EventApiAuthConfig{
		AuthProviders: []appSyncAuthProvider{
			apiKeyProvider,
		},
		ConnectionAuthModeTypes: []appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
	},
	// Custom Domain Settings
	DomainName: &AppSyncDomainOptions{
		Certificate: *Certificate,
		DomainName: myDomainName,
	},
})

api.AddChannelNamespace(jsii.String("default"))

// You can get custom HTTP/Realtime endpoint
// You can get custom HTTP/Realtime endpoint
awscdk.NewCfnOutput(this, jsii.String("AWS AppSync Events HTTP endpoint"), &CfnOutputProps{
	Value: api.customHttpEndpoint,
})
awscdk.NewCfnOutput(this, jsii.String("AWS AppSync Events Realtime endpoint"), &CfnOutputProps{
	Value: api.customRealtimeEndpoint,
})
```

### Log Group

AppSync automatically create a log group with the name `/aws/appsync/apis/<api_id>` upon deployment with log data set to never expire.
If you want to set a different expiration period, use the `logConfig.retention` property.

Also you can choose the log level by setting the `logConfig.fieldLogLevel` property.

For more information, see [Configuring CloudWatch Logs on Event APIs](https://docs.aws.amazon.com/appsync/latest/eventapi/event-api-monitoring-cw-logs.html).

To obtain the Event API's log group as a `logs.ILogGroup` use the `logGroup` property of the
`Api` construct.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


apiKeyProvider := &AppSyncAuthProvider{
	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
}

api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
	ApiName: jsii.String("Api"),
	OwnerContact: jsii.String("OwnerContact"),
	AuthorizationConfig: &EventApiAuthConfig{
		AuthProviders: []appSyncAuthProvider{
			apiKeyProvider,
		},
		ConnectionAuthModeTypes: []appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
	},
	LogConfig: &AppSyncLogConfig{
		FieldLogLevel: appsync.AppSyncFieldLogLevel_INFO,
		Retention: logs.RetentionDays_ONE_WEEK,
	},
})

api.AddChannelNamespace(jsii.String("default"))
```

### WAF Protection

You can use AWS WAF to protect your AppSync API from common web exploits, such as SQL injection and cross-site scripting (XSS) attacks.
These could affect API availability and performance, compromise security, or consume excessive resources.

For more information, see [Using AWS WAF to protect AWS AppSync Event APIs](https://docs.aws.amazon.com/appsync/latest/eventapi/using-waf-protect-apis.html).

```go
var api eventApi
var webAcl cfnWebACL


// Associate waf with Event API
// Associate waf with Event API
wafv2.NewCfnWebACLAssociation(this, jsii.String("WafAssociation"), &CfnWebACLAssociationProps{
	ResourceArn: api.ApiArn,
	WebAclArn: webAcl.AttrArn,
})
```

### Channel namespaces

Channel namespaces define the channels that are available on your Event API, and the capabilities and behaviors of these channels.
Channel namespaces provide a scalable approach to managing large numbers of channels.

Instead of configuring each channel individually, developers can apply settings across an entire namespace.

Channel namespace can optionally interact with data sources configured on the Event API by defining optional event handler code or using direct integrations with the data source where applicable.

For more information, see [Understanding channel namespaces](https://docs.aws.amazon.com/appsync/latest/eventapi/channel-namespaces.html).

```go
var api eventApi


// create a channel namespace
// create a channel namespace
appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
	Api: Api,
})

// You can also create a namespace through the addChannelNamespace method
api.AddChannelNamespace(jsii.String("AnotherNameSpace"), &ChannelNamespaceOptions{
})
```

The API's publishing and subscribing authorization configuration is automatically applied to all namespaces.
You can override this configuration at the namespace level. **Note**: the authorization type you select for a namespace must be defined as an authorization provider at the API level.

```go
var api eventApi


appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
	Api: Api,
	AuthorizationConfig: &NamespaceAuthConfig{
		// Override publishing authorization to API Key
		PublishAuthModeTypes: []appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_API_KEY,
		},
		// Override subscribing authorization to Lambda
		SubscribeAuthModeTypes: []*appSyncAuthorizationType{
			appsync.*appSyncAuthorizationType_LAMBDA,
		},
	},
})
```

You can define event handlers on channel namespaces. Event handlers are functions that run on AWS AppSync's JavaScript runtime and enable you to run custom business logic.
You can use an event handler to process published events or process and authorize subscribe requests.

For more information, see [Channel namespace handlers and event processing](https://docs.aws.amazon.com/appsync/latest/eventapi/channel-namespace-handlers.html).

```go
var api eventApi


appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
	Api: Api,
	// set a handler from inline code
	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
})

appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
	Api: Api,
	// set a handler from an asset
	Code: appsync.Code_FromAsset(jsii.String("directory/function_code.js")),
})
```

You can define an integration in your event handler for `onPublish` and/or `onSubscribe` operations. When defining integrations on your channel namespace, you write code in the event handler to submit requests to and process responses from your data source. For example, if you configure an integration with Amazon DynamoDB for `onPublish` operations, you can persist those events to DynamoDB using a `batchPut` operation in the `request` method, and then return the events as normal in the `response` method. For an integration with Amazon OpenSearch Service, you may use this for `onPublish` operations to enrich the events.

When using the AWS Lambda data source integration, you can either invoke the Lambda function using the event handler code or you can directly invoke the Lambda function, bypassing the event handler code all together. When using direct invoke, you can choose to invoke the Lambda function synchronously or asynchronously by specifying the `invokeType` as `REQUEST_RESPONSE` or `EVENT` respectively.

Below are examples using Amazon DynamoDB, Amazon EventBridge, and AWS Lambda. You can leverage any supported data source in the same way.

#### Amazon DynamoDB & Amazon EventBridge

```go
var api eventApi
var ddbDataSource appSyncDynamoDbDataSource
var ebDataSource appSyncEventBridgeDataSource


// DynamoDB data source for publish handler
api.AddChannelNamespace(jsii.String("ddb-eb-ns"), &ChannelNamespaceOptions{
	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
	PublishHandlerConfig: &HandlerConfig{
		DataSource: ddbDataSource,
	},
	SubscribeHandlerConfig: &HandlerConfig{
		DataSource: ebDataSource,
	},
})
```

#### AWS Lambda

```go
var api eventApi
var lambdaDataSource appSyncLambdaDataSource


// Lambda data source for publish handler
api.AddChannelNamespace(jsii.String("lambda-ns"), &ChannelNamespaceOptions{
	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
	PublishHandlerConfig: &HandlerConfig{
		DataSource: lambdaDataSource,
	},
})

// Direct Lambda data source for publish handler
api.AddChannelNamespace(jsii.String("lambda-direct-ns"), &ChannelNamespaceOptions{
	PublishHandlerConfig: &HandlerConfig{
		DataSource: lambdaDataSource,
		Direct: jsii.Boolean(true),
	},
})

api.AddChannelNamespace(jsii.String("lambda-direct-async-ns"), &ChannelNamespaceOptions{
	PublishHandlerConfig: &HandlerConfig{
		DataSource: lambdaDataSource,
		Direct: jsii.Boolean(true),
		LambdaInvokeType: appsync.LambdaInvokeType_EVENT,
	},
})
```
