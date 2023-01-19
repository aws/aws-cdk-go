package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the attribute value assignments.
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
type AttributeValues interface {
	// Allows assigning a value to the specified attribute.
	// Experimental.
	Attribute(attr *string) AttributeValuesStep
	// Renders the attribute value assingments to a VTL string.
	// Experimental.
	RenderTemplate() *string
	// Renders the variables required for `renderTemplate`.
	// Experimental.
	RenderVariables() *string
}

// The jsii proxy struct for AttributeValues
type jsiiProxy_AttributeValues struct {
	_ byte // padding
}

// Experimental.
func NewAttributeValues(container *string, assignments *[]Assign) AttributeValues {
	_init_.Initialize()

	if err := validateNewAttributeValuesParameters(container); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttributeValues{}

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValues",
		[]interface{}{container, assignments},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeValues_Override(a AttributeValues, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValues",
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

