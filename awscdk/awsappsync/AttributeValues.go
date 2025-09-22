package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the attribute value assignments.
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
type AttributeValues interface {
	// Allows assigning a value to the specified attribute.
	Attribute(attr *string) AttributeValuesStep
	// Renders the attribute value assingments to a VTL string.
	RenderTemplate() *string
	// Renders the variables required for `renderTemplate`.
	RenderVariables() *string
}

// The jsii proxy struct for AttributeValues
type jsiiProxy_AttributeValues struct {
	_ byte // padding
}

func NewAttributeValues(container *string, assignments *[]Assign) AttributeValues {
	_init_.Initialize()

	if err := validateNewAttributeValuesParameters(container); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttributeValues{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AttributeValues",
		[]interface{}{container, assignments},
		&j,
	)

	return &j
}

func NewAttributeValues_Override(a AttributeValues, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AttributeValues",
		[]interface{}{container, assignments},
		a,
	)
}

func (a *jsiiProxy_AttributeValues) Attribute(attr *string) AttributeValuesStep {
	if err := a.validateAttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns AttributeValuesStep

	_jsii_.Invoke(
		a,
		"attribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeValues) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeValues) RenderVariables() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderVariables",
		nil, // no parameters
		&returns,
	)

	return returns
}

