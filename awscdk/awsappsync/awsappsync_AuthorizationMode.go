package awsappsync


// Interface to specify default or additional authorization(s).
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
//   	Name: jsii.String("demo"),
//   	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
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
//   demoDS.CreateResolver(&BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemos"),
//   	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbScanTable(),
//   	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultList(),
//   })
//
//   // Resolver for the Mutation "addDemo" that puts the item into the DynamoDb table.
//   demoDS.CreateResolver(&BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemo"),
//   	RequestMappingTemplate: appsync.MappingTemplate_DynamoDbPutItem(appsync.PrimaryKey_Partition(jsii.String("id")).Auto(), appsync.Values_Projecting(jsii.String("input"))),
//   	ResponseMappingTemplate: appsync.MappingTemplate_DynamoDbResultItem(),
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

