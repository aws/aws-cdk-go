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
api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
	Name: jsii.String("demo"),
	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
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
demoDS.CreateResolver(&BaseResolverProps{
	TypeName: jsii.String("Query"),
	FieldName: jsii.String("getDemos"),
	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbScanTable(),
	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultList(),
})

// Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
demoDS.CreateResolver(&BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("addDemo"),
	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbPutItem(appsync.PrimaryKey_Partition(jsii.String("id")).Auto(), appsync.Values_Projecting(jsii.String("input"))),
	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultItem(),
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
rdsDS.CreateResolver(&BaseResolverProps{
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
rdsDS.CreateResolver(&BaseResolverProps{
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
	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
})

httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
	Name: jsii.String("httpDsWithStepF"),
	Description: jsii.String("from appsync to StepFunctions Workflow"),
	AuthorizationConfig: &AwsIamConfig{
		SigningRegion: jsii.String("us-east-1"),
		SigningServiceName: jsii.String("states"),
	},
})

httpDs.CreateResolver(&BaseResolverProps{
	TypeName: jsii.String("Mutation"),
	FieldName: jsii.String("callStepFunction"),
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
	Version: opensearch.EngineVersion_OPENSEARCH_1_2(),
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

ds.CreateResolver(&BaseResolverProps{
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
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
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
	DomainName: myDomainName,
})
```

## Schema

Every GraphQL Api needs a schema to define the Api. CDK offers `appsync.Schema`
for static convenience methods for various types of schema declaration: code-first
or schema-first.

### Code-First

When declaring your GraphQL Api, CDK defaults to a code-first approach if the
`schema` property is not configured.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
})
```

CDK will declare a `Schema` class that will give your Api access functions to
define your schema code-first: `addType`, `addToSchema`, etc.

You can also declare your `Schema` class outside of your CDK stack, to define
your schema externally.

```go
schema := appsync.NewSchema()
schema.AddType(appsync.NewObjectType(jsii.String("demo"), &ObjectTypeOptions{
	Definition: map[string]iField{
		"id": appsync.GraphqlType_id(),
	},
}))
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
	Schema: Schema,
})
```

See the [code-first schema](#Code-First-Schema) section for more details.

### Schema-First

You can define your GraphQL Schema from a file on disk. For convenience, use
the `appsync.Schema.fromAsset` to specify the file representing your schema.

```go
api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("myApi"),
	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("schema.graphl"))),
})
```

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
	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
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
var api graphqlApi
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
var api graphqlApi
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

Learn more about Pipeline Resolvers and AppSync Functions [here](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html).

## Code-First Schema

CDK offers the ability to generate your schema in a code-first approach.
A code-first approach offers a developer workflow with:

* **modularity**: organizing schema type definitions into different files
* **reusability**: simplifying down boilerplate/repetitive code
* **consistency**: resolvers and schema definition will always be synced

The code-first approach allows for **dynamic** schema generation. You can generate your schema based on variables and templates to reduce code duplication.

### Code-First Example

To showcase the code-first approach. Let's try to model the following schema segment.

```gql
interface Node {
  id: String
}

type Query {
  allFilms(after: String, first: Int, before: String, last: Int): FilmConnection
}

type FilmNode implements Node {
  filmName: String
}

type FilmConnection {
  edges: [FilmEdge]
  films: [Film]
  totalCount: Int
}

type FilmEdge {
  node: Film
  cursor: String
}
```

Above we see a schema that allows for generating paginated responses. For example,
we can query `allFilms(first: 100)` since `FilmConnection` acts as an intermediary
for holding `FilmEdges` we can write a resolver to return the first 100 films.

In a separate file, we can declare our object types and related functions.
We will call this file `object-types.ts` and we will have created it in a way that
allows us to generate other `XxxConnection` and `XxxEdges` in the future.

```go
import "github.com/aws/aws-cdk-go/awscdk"
pluralize := require(jsii.String("pluralize"))

args := map[string]graphqlType{
	"after": appsync.graphqlType_string(),
	"first": appsync.graphqlType_int(),
	"before": appsync.graphqlType_string(),
	"last": appsync.graphqlType_int(),
}

node := appsync.NewInterfaceType(jsii.String("Node"), &IntermediateTypeOptions{
	Definition: map[string]iField{
		"id": appsync.*graphqlType_string(),
	},
})
filmNode := appsync.NewObjectType(jsii.String("FilmNode"), &ObjectTypeOptions{
	InterfaceTypes: []interfaceType{
		*node,
	},
	Definition: map[string]*iField{
		"filmName": appsync.*graphqlType_string(),
	},
})

func GenerateEdgeAndConnection(base *objectType) map[string]objectType {
	edge := appsync.NewObjectType(fmt.Sprintf("%vEdge", *base.Name), &ObjectTypeOptions{
		Definition: map[string]*iField{
			"node": base.attribute(),
			"cursor": appsync.*graphqlType_string(),
		},
	})
	connection := appsync.NewObjectType(fmt.Sprintf("%vConnection", *base.Name), &ObjectTypeOptions{
		Definition: map[string]*iField{
			"edges": edge.attribute(&BaseTypeOptions{
				"isList": jsii.Boolean(true),
			}),
			pluralize(base.*Name): base.attribute(&BaseTypeOptions{
				"isList": jsii.Boolean(true),
			}),
			"totalCount": appsync.*graphqlType_int(),
		},
	})
	return map[string]objectType{
		"edge": edge,
		"connection": connection,
	}
}
```

Finally, we will go to our `cdk-stack` and combine everything together
to generate our schema.

```go
var dummyRequest mappingTemplate
var dummyResponse mappingTemplate


api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
	Name: jsii.String("demo"),
})

objectTypes := []interfaceType{
	node,
	filmNode,
}

filmConnections := generateEdgeAndConnection(filmNode)

api.AddQuery(jsii.String("allFilms"), appsync.NewResolvableField(&ResolvableFieldOptions{
	ReturnType: filmConnections.connection.Attribute(),
	Args: args,
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	RequestMappingTemplate: dummyRequest,
	ResponseMappingTemplate: dummyResponse,
}))

api.AddType(node)
api.AddType(filmNode)
api.AddType(filmConnections.edge)
api.AddType(filmConnections.connection)
```

Notice how we can utilize the `generateEdgeAndConnection` function to generate
Object Types. In the future, if we wanted to create more Object Types, we can simply
create the base Object Type (i.e. Film) and from there we can generate its respective
`Connections` and `Edges`.

Check out a more in-depth example [here](https://github.com/BryanPan342/starwars-code-first).

## GraphQL Types

One of the benefits of GraphQL is its strongly typed nature. We define the
types within an object, query, mutation, interface, etc. as **GraphQL Types**.

GraphQL Types are the building blocks of types, whether they are scalar, objects,
interfaces, etc. GraphQL Types can be:

* [**Scalar Types**](https://docs.aws.amazon.com/appsync/latest/devguide/scalars.html): Id, Int, String, AWSDate, etc.
* [**Object Types**](#Object-Types): types that you generate (i.e. `demo` from the example above)
* [**Interface Types**](#Interface-Types): abstract types that define the base implementation of other
  Intermediate Types

More concretely, GraphQL Types are simply the types appended to variables.
Referencing the object type `Demo` in the previous example, the GraphQL Types
is `String!` and is applied to both the names `id` and `version`.

### Directives

`Directives` are attached to a field or type and affect the execution of queries,
mutations, and types. With AppSync, we use `Directives` to configure authorization.
CDK provides static functions to add directives to your Schema.

* `Directive.iam()` sets a type or field's authorization to be validated through `Iam`
* `Directive.apiKey()` sets a type or field's authorization to be validated through a `Api Key`
* `Directive.oidc()` sets a type or field's authorization to be validated through `OpenID Connect`
* `Directive.cognito(...groups: string[])` sets a type or field's authorization to be validated
  through `Cognito User Pools`

  * `groups` the name of the cognito groups to give access

To learn more about authorization and directives, read these docs [here](https://docs.aws.amazon.com/appsync/latest/devguide/security.html).

### Field and Resolvable Fields

While `GraphqlType` is a base implementation for GraphQL fields, we have abstractions
on top of `GraphqlType` that provide finer grain support.

### Field

`Field` extends `GraphqlType` and will allow you to define arguments. [**Interface Types**](#Interface-Types) are not resolvable and this class will allow you to define arguments,
but not its resolvers.

For example, if we want to create the following type:

```gql
type Node {
  test(argument: string): String
}
```

The CDK code required would be:

```go
field := appsync.NewField(&FieldOptions{
	ReturnType: appsync.GraphqlType_String(),
	Args: map[string]graphqlType{
		"argument": appsync.*graphqlType_*String(),
	},
})
type := appsync.NewInterfaceType(jsii.String("Node"), &IntermediateTypeOptions{
	Definition: map[string]iField{
		"test": field,
	},
})
```

### Resolvable Fields

`ResolvableField` extends `Field` and will allow you to define arguments and its resolvers.
[**Object Types**](#Object-Types) can have fields that resolve and perform operations on
your backend.

You can also create resolvable fields for object types.

```gql
type Info {
  node(id: String): String
}
```

The CDK code required would be:

```go
var api graphqlApi
var dummyRequest mappingTemplate
var dummyResponse mappingTemplate

info := appsync.NewObjectType(jsii.String("Info"), &ObjectTypeOptions{
	Definition: map[string]iField{
		"node": appsync.NewResolvableField(&ResolvableFieldOptions{
			"returnType": appsync.GraphqlType_string(),
			"args": map[string]GraphqlType{
				"id": appsync.GraphqlType_string(),
			},
			"dataSource": api.addNoneDataSource(jsii.String("none")),
			"requestMappingTemplate": dummyRequest,
			"responseMappingTemplate": dummyResponse,
		}),
	},
})
```

To nest resolvers, we can also create top level query types that call upon
other types. Building off the previous example, if we want the following graphql
type definition:

```gql
type Query {
  get(argument: string): Info
}
```

The CDK code required would be:

```go
var api graphqlApi
var dummyRequest mappingTemplate
var dummyResponse mappingTemplate

query := appsync.NewObjectType(jsii.String("Query"), &ObjectTypeOptions{
	Definition: map[string]iField{
		"get": appsync.NewResolvableField(&ResolvableFieldOptions{
			"returnType": appsync.GraphqlType_string(),
			"args": map[string]GraphqlType{
				"argument": appsync.GraphqlType_string(),
			},
			"dataSource": api.addNoneDataSource(jsii.String("none")),
			"requestMappingTemplate": dummyRequest,
			"responseMappingTemplate": dummyResponse,
		}),
	},
})
```

Learn more about fields and resolvers [here](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-overview.html).

### Intermediate Types

Intermediate Types are defined by Graphql Types and Fields. They have a set of defined
fields, where each field corresponds to another type in the system. Intermediate
Types will be the meat of your GraphQL Schema as they are the types defined by you.

Intermediate Types include:

* [**Interface Types**](#Interface-Types)
* [**Object Types**](#Object-Types)
* [**Enum Types**](#Enum-Types)
* [**Input Types**](#Input-Types)
* [**Union Types**](#Union-Types)

#### Interface Types

**Interface Types** are abstract types that define the implementation of other
intermediate types. They are useful for eliminating duplication and can be used
to generate Object Types with less work.

You can create Interface Types ***externally***.

```go
node := appsync.NewInterfaceType(jsii.String("Node"), &IntermediateTypeOptions{
	Definition: map[string]iField{
		"id": appsync.GraphqlType_string(&BaseTypeOptions{
			"isRequired": jsii.Boolean(true),
		}),
	},
})
```

To learn more about **Interface Types**, read the docs [here](https://graphql.org/learn/schema/#interfaces).

#### Object Types

**Object Types** are types that you declare. For example, in the [code-first example](#code-first-example)
the `demo` variable is an **Object Type**. **Object Types** are defined by
GraphQL Types and are only usable when linked to a GraphQL Api.

You can create Object Types in two ways:

1. Object Types can be created ***externally***.

   ```go
   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
   	Name: jsii.String("demo"),
   })
   demo := appsync.NewObjectType(jsii.String("Demo"), &ObjectTypeOptions{
   	Definition: map[string]iField{
   		"id": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   		"version": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   	},
   })

   api.AddType(demo)
   ```

   > This method allows for reusability and modularity, ideal for larger projects.
   > For example, imagine moving all Object Type definition outside the stack.

   `object-types.ts` - a file for object type definitions

   ```go
   import "github.com/aws/aws-cdk-go/awscdk"
   demo := appsync.NewObjectType(jsii.String("Demo"), &ObjectTypeOptions{
   	Definition: map[string]iField{
   		"id": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   		"version": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   	},
   })
   ```

   `cdk-stack.ts` - a file containing our cdk stack

   ```go
   var api graphqlApi

   api.AddType(demo)
   ```
2. Object Types can be created ***externally*** from an Interface Type.

   ```go
   node := appsync.NewInterfaceType(jsii.String("Node"), &IntermediateTypeOptions{
   	Definition: map[string]iField{
   		"id": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   	},
   })
   demo := appsync.NewObjectType(jsii.String("Demo"), &ObjectTypeOptions{
   	InterfaceTypes: []interfaceType{
   		node,
   	},
   	Definition: map[string]*iField{
   		"version": appsync.GraphqlType_string(&BaseTypeOptions{
   			"isRequired": jsii.Boolean(true),
   		}),
   	},
   })
   ```

   > This method allows for reusability and modularity, ideal for reducing code duplication.

To learn more about **Object Types**, read the docs [here](https://graphql.org/learn/schema/#object-types-and-fields).

#### Enum Types

**Enum Types** are a special type of Intermediate Type. They restrict a particular
set of allowed values for other Intermediate Types.

```gql
enum Episode {
  NEWHOPE
  EMPIRE
  JEDI
}
```

> This means that wherever we use the type Episode in our schema, we expect it to
> be exactly one of NEWHOPE, EMPIRE, or JEDI.

The above GraphQL Enumeration Type can be expressed in CDK as the following:

```go
var api graphqlApi

episode := appsync.NewEnumType(jsii.String("Episode"), &EnumTypeOptions{
	Definition: []*string{
		jsii.String("NEWHOPE"),
		jsii.String("EMPIRE"),
		jsii.String("JEDI"),
	},
})
api.AddType(episode)
```

To learn more about **Enum Types**, read the docs [here](https://graphql.org/learn/schema/#enumeration-types).

#### Input Types

**Input Types** are special types of Intermediate Types. They give users an
easy way to pass complex objects for top level Mutation and Queries.

```gql
input Review {
  stars: Int!
  commentary: String
}
```

The above GraphQL Input Type can be expressed in CDK as the following:

```go
var api graphqlApi

review := appsync.NewInputType(jsii.String("Review"), &IntermediateTypeOptions{
	Definition: map[string]iField{
		"stars": appsync.GraphqlType_int(&BaseTypeOptions{
			"isRequired": jsii.Boolean(true),
		}),
		"commentary": appsync.GraphqlType_string(),
	},
})
api.AddType(review)
```

To learn more about **Input Types**, read the docs [here](https://graphql.org/learn/schema/#input-types).

#### Union Types

**Union Types** are a special type of Intermediate Type. They are similar to
Interface Types, but they cannot specify any common fields between types.

**Note:** the fields of a union type need to be `Object Types`. In other words, you
can't create a union type out of interfaces, other unions, or inputs.

```gql
union Search = Human | Droid | Starship
```

The above GraphQL Union Type encompasses the Object Types of Human, Droid and Starship. It
can be expressed in CDK as the following:

```go
var api graphqlApi

string := appsync.GraphqlType_String()
human := appsync.NewObjectType(jsii.String("Human"), &ObjectTypeOptions{
	Definition: map[string]iField{
		"name": string,
	},
})
droid := appsync.NewObjectType(jsii.String("Droid"), &ObjectTypeOptions{
	Definition: map[string]*iField{
		"name": string,
	},
})
starship := appsync.NewObjectType(jsii.String("Starship"), &ObjectTypeOptions{
	Definition: map[string]*iField{
		"name": string,
	},
})
search := appsync.NewUnionType(jsii.String("Search"), &UnionTypeOptions{
	Definition: []iIntermediateType{
		human,
		droid,
		starship,
	},
})
api.AddType(search)
```

To learn more about **Union Types**, read the docs [here](https://graphql.org/learn/schema/#union-types).

### Query

Every schema requires a top level Query type. By default, the schema will look
for the `Object Type` named `Query`. The top level `Query` is the **only** exposed
type that users can access to perform `GET` operations on your Api.

To add fields for these queries, we can simply run the `addQuery` function to add
to the schema's `Query` type.

```go
var api graphqlApi
var filmConnection interfaceType
var dummyRequest mappingTemplate
var dummyResponse mappingTemplate


string := appsync.GraphqlType_String()
int := appsync.GraphqlType_Int()
api.AddQuery(jsii.String("allFilms"), appsync.NewResolvableField(&ResolvableFieldOptions{
	ReturnType: filmConnection.Attribute(),
	Args: map[string]graphqlType{
		"after": string,
		"first": int,
		"before": string,
		"last": int,
	},
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	RequestMappingTemplate: dummyRequest,
	ResponseMappingTemplate: dummyResponse,
}))
```

To learn more about top level operations, check out the docs [here](https://docs.aws.amazon.com/appsync/latest/devguide/graphql-overview.html).

### Mutation

Every schema **can** have a top level Mutation type. By default, the schema will look
for the `ObjectType` named `Mutation`. The top level `Mutation` Type is the only exposed
type that users can access to perform `mutable` operations on your Api.

To add fields for these mutations, we can simply run the `addMutation` function to add
to the schema's `Mutation` type.

```go
var api graphqlApi
var filmNode objectType
var dummyRequest mappingTemplate
var dummyResponse mappingTemplate


string := appsync.GraphqlType_String()
int := appsync.GraphqlType_Int()
api.AddMutation(jsii.String("addFilm"), appsync.NewResolvableField(&ResolvableFieldOptions{
	ReturnType: filmNode.Attribute(),
	Args: map[string]graphqlType{
		"name": string,
		"film_number": int,
	},
	DataSource: api.AddNoneDataSource(jsii.String("none")),
	RequestMappingTemplate: dummyRequest,
	ResponseMappingTemplate: dummyResponse,
}))
```

To learn more about top level operations, check out the docs [here](https://docs.aws.amazon.com/appsync/latest/devguide/graphql-overview.html).

### Subscription

Every schema **can** have a top level Subscription type. The top level `Subscription` Type
is the only exposed type that users can access to invoke a response to a mutation. `Subscriptions`
notify users when a mutation specific mutation is called. This means you can make any data source
real time by specify a GraphQL Schema directive on a mutation.

**Note**: The AWS AppSync client SDK automatically handles subscription connection management.

To add fields for these subscriptions, we can simply run the `addSubscription` function to add
to the schema's `Subscription` type.

```go
var api graphqlApi
var film interfaceType


api.AddSubscription(jsii.String("addedFilm"), appsync.NewField(&FieldOptions{
	ReturnType: film.Attribute(),
	Args: map[string]graphqlType{
		"id": appsync.*graphqlType_id(&BaseTypeOptions{
			"isRequired": jsii.Boolean(true),
		}),
	},
	Directives: []directive{
		appsync.*directive_Subscribe(jsii.String("addFilm")),
	},
}))
```

To learn more about top level operations, check out the docs [here](https://docs.aws.amazon.com/appsync/latest/devguide/real-time-data.html).
