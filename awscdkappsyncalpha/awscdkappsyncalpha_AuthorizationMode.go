// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Interface to specify default or additional authorization(s).
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
//   // Resolver Mapping Template Reference:
//   // https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-dynamodb.html
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
//   //To enable DynamoDB read consistency with the `MappingTemplate`:
//   demoDS.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemosConsistent"),
//   	requestMappingTemplate: appsync.*mappingTemplate.dynamoDbScanTable(jsii.Boolean(true)),
//   	responseMappingTemplate: appsync.*mappingTemplate.dynamoDbResultList(),
//   })
//
// Experimental.
type AuthorizationMode struct {
	// One of possible four values AppSync supports.
	// See: https://docs.aws.amazon.com/appsync/latest/devguide/security.html
	//
	// Experimental.
	AuthorizationType AuthorizationType `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// If authorizationType is `AuthorizationType.API_KEY`, this option can be configured.
	// Experimental.
	ApiKeyConfig *ApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// If authorizationType is `AuthorizationType.LAMBDA`, this option is required.
	// Experimental.
	LambdaAuthorizerConfig *LambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// If authorizationType is `AuthorizationType.OIDC`, this option is required.
	// Experimental.
	OpenIdConnectConfig *OpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// If authorizationType is `AuthorizationType.USER_POOL`, this option is required.
	// Experimental.
	UserPoolConfig *UserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

