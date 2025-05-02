package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for attribute value assignments.
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
type Values interface {
}

// The jsii proxy struct for Values
type jsiiProxy_Values struct {
	_ byte // padding
}

func NewValues() Values {
	_init_.Initialize()

	j := jsiiProxy_Values{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Values",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewValues_Override(v Values) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Values",
		nil, // no parameters
		v,
	)
}

// Allows assigning a value to the specified attribute.
func Values_Attribute(attr *string) AttributeValuesStep {
	_init_.Initialize()

	if err := validateValues_AttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns AttributeValuesStep

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Values",
		"attribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Treats the specified object as a map of assignments, where the property names represent attribute names.
//
// It’s opinionated about how it represents
// some of the nested objects: e.g., it will use lists (“L”) rather than sets
// (“SS”, “NS”, “BS”). By default it projects the argument container ("$ctx.args").
func Values_Projecting(arg *string) AttributeValues {
	_init_.Initialize()

	var returns AttributeValues

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Values",
		"projecting",
		[]interface{}{arg},
		&returns,
	)

	return returns
}

