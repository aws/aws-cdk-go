package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the assignment to the primary key.
//
// It either
// contains the full primary key or only the partition key.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &graphqlApiProps{
//   	name: jsii.String("demo"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   	authorizationConfig: &authorizationConfig{
//   		defaultAuthorization: &authorizationMode{
//   			authorizationType: appsync.authorizationType_IAM,
//   		},
//   	},
//   	xrayEnabled: jsii.Boolean(true),
//   })
//
//   demoTable := dynamodb.NewTable(this, jsii.String("DemoTable"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   })
//
//   demoDS := api.addDynamoDbDataSource(jsii.String("demoDataSource"), demoTable)
//
//   // Resolver for the Query "getDemos" that scans the DynamoDb table and returns the entire list.
//   demoDS.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemos"),
//   	requestMappingTemplate: appsync.mappingTemplate.dynamoDbScanTable(),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
//   })
//
//   // Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
//   demoDS.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("addDemo"),
//   	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbPutItem(appsync.primaryKey.partition(jsii.String("id")).auto(), appsync.values.projecting(jsii.String("input"))),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultItem(),
//   })
//
// Experimental.
type PrimaryKey interface {
	// Experimental.
	Pkey() Assign
	// Renders the key assignment to a VTL string.
	// Experimental.
	RenderTemplate() *string
}

// The jsii proxy struct for PrimaryKey
type jsiiProxy_PrimaryKey struct {
	_ byte // padding
}

func (j *jsiiProxy_PrimaryKey) Pkey() Assign {
	var returns Assign
	_jsii_.Get(
		j,
		"pkey",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrimaryKey(pkey Assign, skey Assign) PrimaryKey {
	_init_.Initialize()

	if err := validateNewPrimaryKeyParameters(pkey); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrimaryKey{}

	_jsii_.Create(
		"monocdk.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		&j,
	)

	return &j
}

// Experimental.
func NewPrimaryKey_Override(p PrimaryKey, pkey Assign, skey Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		p,
	)
}

// Allows assigning a value to the partition key.
// Experimental.
func PrimaryKey_Partition(key *string) PartitionKeyStep {
	_init_.Initialize()

	if err := validatePrimaryKey_PartitionParameters(key); err != nil {
		panic(err)
	}
	var returns PartitionKeyStep

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.PrimaryKey",
		"partition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrimaryKey) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

