// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for attribute value assignments.
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
// Experimental.
type Values interface {
}

// The jsii proxy struct for Values
type jsiiProxy_Values struct {
	_ byte // padding
}

// Experimental.
func NewValues() Values {
	_init_.Initialize()

	j := jsiiProxy_Values{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Values",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewValues_Override(v Values) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Values",
		nil, // no parameters
		v,
	)
}

// Allows assigning a value to the specified attribute.
// Experimental.
func Values_Attribute(attr *string) AttributeValuesStep {
	_init_.Initialize()

	if err := validateValues_AttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns AttributeValuesStep

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Values",
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
// Experimental.
func Values_Projecting(arg *string) AttributeValues {
	_init_.Initialize()

	var returns AttributeValues

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Values",
		"projecting",
		[]interface{}{arg},
		&returns,
	)

	return returns
}

