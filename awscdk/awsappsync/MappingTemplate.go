package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MappingTemplates for AppSync resolvers.
//
// Example:
//   // Build a data source for AppSync to access the database.
//   var api GraphqlApi
//   // Create username and password secret for DB Cluster
//   secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
//   	Username: jsii.String("clusteradmin"),
//   })
//
//   // The VPC to place the cluster in
//   vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))
//
//   // Create the serverless cluster, provide all values needed to customise the database.
//   cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	ClusterIdentifier: jsii.String("db-endpoint-test"),
//   	DefaultDatabaseName: jsii.String("demos"),
//   })
//   rdsDS := api.AddRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))
//
//   // Set up a resolver for an RDS query.
//   rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemosRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "SELECT * FROM demos"
//   	    ]
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
//   	  `)),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemoRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "INSERT INTO demos VALUES (:id, :version)",
//   	      "SELECT * WHERE id = :id"
//   	    ],
//   	    "variableMap": {
//   	      ":id": $util.toJson($util.autoId()),
//   	      ":version": $util.toJson($ctx.args.version)
//   	    }
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
//   	  `)),
//   })
//
type MappingTemplate interface {
	// this is called to render the mapping template to a VTL string.
	RenderTemplate() *string
}

// The jsii proxy struct for MappingTemplate
type jsiiProxy_MappingTemplate struct {
	_ byte // padding
}

func NewMappingTemplate_Override(m MappingTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		nil, // no parameters
		m,
	)
}

// Mapping template to delete a single item from a DynamoDB table.
func MappingTemplate_DynamoDbDeleteItem(keyName *string, idArg *string) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_DynamoDbDeleteItemParameters(keyName, idArg); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbDeleteItem",
		[]interface{}{keyName, idArg},
		&returns,
	)

	return returns
}

// Mapping template to get a single item from a DynamoDB table.
func MappingTemplate_DynamoDbGetItem(keyName *string, idArg *string, consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_DynamoDbGetItemParameters(keyName, idArg); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbGetItem",
		[]interface{}{keyName, idArg, consistentRead},
		&returns,
	)

	return returns
}

// Mapping template to save a single item to a DynamoDB table.
func MappingTemplate_DynamoDbPutItem(key PrimaryKey, values AttributeValues) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_DynamoDbPutItemParameters(key, values); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbPutItem",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Mapping template to query a set of items from a DynamoDB table.
func MappingTemplate_DynamoDbQuery(cond KeyCondition, indexName *string, consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_DynamoDbQueryParameters(cond); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbQuery",
		[]interface{}{cond, indexName, consistentRead},
		&returns,
	)

	return returns
}

// Mapping template for a single result item from DynamoDB.
func MappingTemplate_DynamoDbResultItem() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbResultItem",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template for a result list from DynamoDB.
func MappingTemplate_DynamoDbResultList() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbResultList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template to scan a DynamoDB table to fetch all entries.
func MappingTemplate_DynamoDbScanTable(consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"dynamoDbScanTable",
		[]interface{}{consistentRead},
		&returns,
	)

	return returns
}

// Create a mapping template from the given file.
func MappingTemplate_FromFile(fileName *string) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_FromFileParameters(fileName); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"fromFile",
		[]interface{}{fileName},
		&returns,
	)

	return returns
}

// Create a mapping template from the given string.
func MappingTemplate_FromString(template *string) MappingTemplate {
	_init_.Initialize()

	if err := validateMappingTemplate_FromStringParameters(template); err != nil {
		panic(err)
	}
	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Mapping template to invoke a Lambda function.
func MappingTemplate_LambdaRequest(payload *string, operation *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"lambdaRequest",
		[]interface{}{payload, operation},
		&returns,
	)

	return returns
}

// Mapping template to return the Lambda result to the caller.
func MappingTemplate_LambdaResult() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		"lambdaResult",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MappingTemplate) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

