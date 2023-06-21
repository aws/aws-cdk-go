package awsappsync


// enum with all possible values for AppSync authorization type.
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
type AuthorizationType string

const (
	// API Key authorization type.
	AuthorizationType_API_KEY AuthorizationType = "API_KEY"
	// AWS IAM authorization type.
	//
	// Can be used with Cognito Identity Pool federated credentials.
	AuthorizationType_IAM AuthorizationType = "IAM"
	// Cognito User Pool authorization type.
	AuthorizationType_USER_POOL AuthorizationType = "USER_POOL"
	// OpenID Connect authorization type.
	AuthorizationType_OIDC AuthorizationType = "OIDC"
	// Lambda authorization type.
	AuthorizationType_LAMBDA AuthorizationType = "LAMBDA"
)

