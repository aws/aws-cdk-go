package awsappsync


// Interface to specify default or additional authorization(s).
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &GraphqlApiProps{
//   	Name: jsii.String("demo"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
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
type AuthorizationMode struct {
	// One of possible four values AppSync supports.
	// See: https://docs.aws.amazon.com/appsync/latest/devguide/security.html
	//
	// Default: - `AuthorizationType.API_KEY`
	//
	AuthorizationType AuthorizationType `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// If authorizationType is `AuthorizationType.API_KEY`, this option can be configured.
	// Default: - name: 'DefaultAPIKey' | description: 'Default API Key created by CDK'.
	//
	ApiKeyConfig *ApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// If authorizationType is `AuthorizationType.LAMBDA`, this option is required.
	// Default: - none.
	//
	LambdaAuthorizerConfig *LambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// If authorizationType is `AuthorizationType.OIDC`, this option is required.
	// Default: - none.
	//
	OpenIdConnectConfig *OpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// If authorizationType is `AuthorizationType.USER_POOL`, this option is required.
	// Default: - none.
	//
	UserPoolConfig *UserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

