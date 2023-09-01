package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the assignment to the primary key.
//
// It either
// contains the full primary key or only the partition key.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
//   	Name: jsii.String("demo"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphql"))),
//   	AuthorizationConfig: &AuthorizationConfig{
//   		DefaultAuthorization: &AuthorizationMode{
//   			AuthorizationType: appsync.AuthorizationType_IAM,
//   		},
//   	},
//   	XrayEnabled: jsii.Boolean(true),
//   })
//
//   demoTable := dynamodb.NewTable(this, jsii.String("DemoTable"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
//   demoDS := api.AddDynamoDbDataSource(jsii.String("demoDataSource"), demoTable)
//
//   // Resolver for the Query "getDemos" that scans the DynamoDb table and returns the entire list.
//   // Resolver Mapping Template Reference:
//   // https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-dynamodb.html
//   demoDS.CreateResolver(jsii.String("QueryGetDemosResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemos"),
//   	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbScanTable(),
//   	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultList(),
//   })
//
//   // Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
//   demoDS.CreateResolver(jsii.String("MutationAddDemoResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemo"),
//   	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbPutItem(appsync.PrimaryKey_Partition(jsii.String("id")).Auto(), appsync.Values_Projecting(jsii.String("input"))),
//   	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultItem(),
//   })
//
//   //To enable DynamoDB read consistency with the `MappingTemplate`:
//   demoDS.CreateResolver(jsii.String("QueryGetDemosConsistentResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemosConsistent"),
//   	RequestMappingTemplate: appsync.MappingTemplate_*DynamoDbScanTable(jsii.Boolean(true)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*DynamoDbResultList(),
//   })
//
type PrimaryKey interface {
	Pkey() Assign
	// Renders the key assignment to a VTL string.
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


func NewPrimaryKey(pkey Assign, skey Assign) PrimaryKey {
	_init_.Initialize()

	if err := validateNewPrimaryKeyParameters(pkey); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrimaryKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		&j,
	)

	return &j
}

func NewPrimaryKey_Override(p PrimaryKey, pkey Assign, skey Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		p,
	)
}

// Allows assigning a value to the partition key.
func PrimaryKey_Partition(key *string) PartitionKeyStep {
	_init_.Initialize()

	if err := validatePrimaryKey_PartitionParameters(key); err != nil {
		panic(err)
	}
	var returns PartitionKeyStep

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.PrimaryKey",
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

