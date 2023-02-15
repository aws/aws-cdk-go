package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the attribute value assignments.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &graphqlApiProps{
//   	name: jsii.String("demo"),
//   	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
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
//   // Resolver Mapping Template Reference:
//   // https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-dynamodb.html
//   demoDS.createResolver(jsii.String("QueryGetDemosResolver"), &baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemos"),
//   	requestMappingTemplate: appsync.mappingTemplate.dynamoDbScanTable(),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
//   })
//
//   // Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
//   demoDS.createResolver(jsii.String("MutationAddDemoResolver"), &baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("addDemo"),
//   	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbPutItem(appsync.primaryKey.partition(jsii.String("id")).auto(), appsync.values.projecting(jsii.String("input"))),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultItem(),
//   })
//
//   //To enable DynamoDB read consistency with the `MappingTemplate`:
//   demoDS.createResolver(jsii.String("QueryGetDemosConsistentResolver"), &baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemosConsistent"),
//   	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbScanTable(jsii.Boolean(true)),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
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

